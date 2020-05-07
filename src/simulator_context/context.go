package simulator_context

import (
	"encoding/hex"
	"fmt"
	"net"
	"os/exec"
	"radio_simulator/lib/ngap/ngapType"
	"radio_simulator/lib/openapi/models"
	"radio_simulator/src/logger"
)

var simContext = Simulator{}

const MaxValueOfPdrFar uint64 = 0xffffffffffffffff

func init() {
	Simulator_Self().RanPool = make(map[string]*RanContext)
	Simulator_Self().PdrPool = make(map[string]*SessionContext)
	Simulator_Self().FarPool = make(map[string]*SessionContext)
	Simulator_Self().UeContextPool = make(map[string]*UeContext)
	Simulator_Self().FARGenerator = 1
	Simulator_Self().PDRGenerator = 0
	// Simulator_Self().GtpConnPool = make(map[string]*net.UDPConn)
}

type Simulator struct {
	DefaultRanSctpUri string
	RanPool           map[string]*RanContext     // RanSctpUri -> RAN_CONTEXT
	PdrPool           map[string]*SessionContext // PDRID -> RAN_CONTEXT
	FarPool           map[string]*SessionContext // FARID -> RAN_CONTEXT
	UeContextPool     map[string]*UeContext      // Supi -> UeTestInfo
	// GtpConnPool       map[string]*net.UDPConn    // "ranGtpuri,upfUri" -> conn
	TcpServer       net.Listener
	FARGenerator    uint64
	PDRGenerator    uint64
	TunDev          string
	Gtp5gTunnelExec string
	// TunMtx            sync.Mutex
	// TunFd             int
	// TunSockAddr       syscall.Sockaddr
	// ListenRawConn     *ipv4.RawConn
}

type UeDBInfo struct {
	AmDate     models.AccessAndMobilitySubscriptionData
	SmfSelData models.SmfSelectionSubscriptionData
	AmPolicy   models.AmPolicyData
	AuthsSubs  models.AuthenticationSubscription
	PlmnId     string
}

// func (s *Simulator) SendToTunDev(msg []byte) {
// 	s.TunMtx.Lock()
// 	syscall.Sendto(s.TunFd, msg, 0, s.TunSockAddr)
// 	s.TunMtx.Unlock()
// }

func (s *Simulator) AddRanContext(AmfUri, ranSctpUri, ranName string, ranGtpUri AddrInfo, plmnId ngapType.PLMNIdentity, GnbId string, gnbIdLength int) *RanContext {
	ran := NewRanContext()
	ran.AMFUri = AmfUri
	ran.RanSctpUri = ranSctpUri
	ran.RanGtpUri = ranGtpUri
	ran.Name = ranName
	ran.GnbId.BitLength = uint64(gnbIdLength)
	ran.GnbId.Bytes, _ = hex.DecodeString(GnbId)
	s.RanPool[ranSctpUri] = ran
	return ran
}

func (s *Simulator) AttachSession(sess *SessionContext) {
	ulAct, dlAct := "mod", "mod"
	if sess.ULFarID == "" {
		// not establish yet
		ulAct = "add"
		sess.ULFarID = s.FARAlloc()
		sess.ULPdrID = s.PDRAlloc()
	}
	if sess.DLPdrID == "" {
		// not establish yet
		dlAct = "add"
		sess.DLPdrID = s.PDRAlloc()
	}
	{
		// "<add|mod> far <tunDev> <id> --action 2 --hdr-creation 1 <ul-teid> <ul-addr> 2152"
		_, err := exec.Command(s.Gtp5gTunnelExec, ulAct, "far", s.TunDev, sess.ULFarID, "--action", "2", "--hdr-creation", "1", fmt.Sprintf("%d", sess.ULTEID), sess.ULAddr, "2152").Output()
		if err != nil {
			logger.ContextLog.Errorf("Create UL FAR Failed[%s]", err.Error())
			return
		}
	}
	{
		// "<add|mod> pdr <tunDev> <id> --pcd 2 --ue-ipv4 <ueIP> --far-id <farId>"
		_, err := exec.Command(s.Gtp5gTunnelExec, ulAct, "pdr", s.TunDev, sess.ULPdrID, "--pcd", "2", "--ue-ipv4", sess.UeIp, "--far-id", sess.ULFarID).Output()
		if err != nil {
			logger.ContextLog.Errorf("Create UL PDR Failed[%s]", err.Error())
			return
		}
	}
	{
		// "<add|mod> pdr <tunDev> <id> --pcd 1 --hdr-rm 0 --ue-ipv4 <ueIP> --f-teid <dl-teid> <dl-addr> --far-id 1"
		_, err := exec.Command(s.Gtp5gTunnelExec, dlAct, "pdr", s.TunDev, sess.DLPdrID, "--pcd", "1", "--hdr-rm", "0", "--ue-ipv4", sess.UeIp, "--f-teid", fmt.Sprintf("%d", sess.DLTEID), sess.DLAddr, "--far-id", "1").Output()
		if err != nil {
			logger.ContextLog.Errorf("Create DL PDR Failed[%s]", err.Error())
			return
		}
	}
	s.PdrPool[sess.DLPdrID] = sess
	s.PdrPool[sess.ULPdrID] = sess
	s.FarPool[sess.ULFarID] = sess
}

func (s *Simulator) DetachSession(sess *SessionContext) {
	{
		_, err := exec.Command(s.Gtp5gTunnelExec, "del", "far", s.TunDev, sess.ULFarID).Output()
		if err != nil {
			logger.ContextLog.Errorf("Delete UL FAR Failed[%s]", err.Error())
		}
	}
	{
		_, err := exec.Command(s.Gtp5gTunnelExec, "del", "pdr", s.TunDev, sess.ULPdrID).Output()
		if err != nil {
			logger.ContextLog.Errorf("Delete UL PDR Failed[%s]", err.Error())
		}
	}
	{
		_, err := exec.Command(s.Gtp5gTunnelExec, "del", "pdr", s.TunDev, sess.DLPdrID).Output()
		if err != nil {
			logger.ContextLog.Errorf("Delete DL PDR Failed[%s]", err.Error())
		}
	}
	delete(s.PdrPool, sess.DLPdrID)
	delete(s.PdrPool, sess.ULPdrID)
	delete(s.FarPool, sess.ULFarID)
}

func (s *Simulator) PDRAlloc() string {
	s.PDRGenerator %= MaxValueOfPdrFar
	s.PDRGenerator++
	for {
		key := fmt.Sprintf("%d", s.PDRGenerator)
		if _, double := s.PdrPool[key]; double {
			s.PDRGenerator++
		} else {
			return key
		}
	}
	return ""
}

func (s *Simulator) FARAlloc() string {
	s.FARGenerator %= MaxValueOfPdrFar
	s.FARGenerator++
	for {
		key := fmt.Sprintf("%d", s.FARGenerator)
		if _, double := s.FarPool[key]; double {
			s.FARGenerator++
		} else {
			return key
		}
	}
	return ""
}

// Create new AMF context
func Simulator_Self() *Simulator {
	return &simContext
}
