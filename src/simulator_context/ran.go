package simulator_context

import (
	"encoding/hex"
	"git.cs.nctu.edu.tw/calee/sctp"
	"net"
	"radio_simulator/lib/aper"
	"radio_simulator/lib/ngap/ngapType"
)

const (
	MaxValueOfTeid uint32 = 0xffffffff
)

type RanContext struct {
	TEIDGenerator    uint32
	RanUeIDGenerator int64
	AMFUri           string
	RanSctpUri       string
	RanGtpUri        AddrInfo
	UpfInfoList      map[string]*UpfInfo
	Name             string
	GnbId            aper.BitString
	UePool           map[int64]*UeContext // ranUeNgapId
	SessPool         map[uint32]*SessionContext
	DefaultTAC       string
	SupportTAList    map[string][]PlmnSupportItem // TAC(hex string) -> PlmnSupportItem
	SctpConn         *sctp.SCTPConn
}

type AddrInfo struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

type UpfInfo struct {
	Addr    AddrInfo
	GtpConn *net.UDPConn
}
type PlmnSupportItem struct {
	PlmnId     ngapType.PLMNIdentity
	SNssaiList []ngapType.SNSSAI
}

func (ran *RanContext) AttachSession(sess *SessionContext) {
	sess.DLAddr = ran.RanGtpUri.IP
	sess.DLTEID = ran.TEIDAlloc()
	ran.SessPool[sess.DLTEID] = sess
}

func (ran *RanContext) DetachSession(sess *SessionContext) {
	delete(ran.SessPool, sess.DLTEID)
}

func (ran *RanContext) TEIDAlloc() uint32 {
	ran.TEIDGenerator %= MaxValueOfTeid
	ran.TEIDGenerator++
	for {
		if _, double := ran.SessPool[ran.TEIDGenerator]; double {
			ran.TEIDGenerator++
		} else {
			break
		}
	}
	return ran.TEIDGenerator
}

func (ran *RanContext) FindUeByRanUeNgapID(ranUeNgapID int64) *UeContext {
	if ue, ok := ran.UePool[ranUeNgapID]; ok {
		return ue
	} else {
		return nil
	}
}

func (ran *RanContext) FindUeByAmfUeNgapID(amfUeNgapID int64) *UeContext {
	for _, ranUe := range ran.UePool {
		if ranUe.AmfUeNgapId == amfUeNgapID {
			return ranUe
		}
	}
	return nil
}

func (ran *RanContext) GetUserLocation() ngapType.UserLocationInformation {
	userLocationInformation := ngapType.UserLocationInformation{}
	userLocationInformation.Present = ngapType.UserLocationInformationPresentUserLocationInformationNR
	userLocationInformation.UserLocationInformationNR = new(ngapType.UserLocationInformationNR)

	userLocationInformationNR := userLocationInformation.UserLocationInformationNR
	userLocationInformationNR.NRCGI.PLMNIdentity = ran.SupportTAList[ran.DefaultTAC][0].PlmnId
	userLocationInformationNR.NRCGI.NRCellIdentity.Value = aper.BitString{
		Bytes:     []byte{0x00, 0x00, 0x00, 0x00, 0x10},
		BitLength: 36,
	}

	userLocationInformationNR.TAI.PLMNIdentity = ran.SupportTAList[ran.DefaultTAC][0].PlmnId
	userLocationInformationNR.TAI.TAC.Value, _ = hex.DecodeString(ran.DefaultTAC)
	return userLocationInformation
}

func NewRanContext() *RanContext {
	return &RanContext{
		RanUeIDGenerator: 1,
		TEIDGenerator:    0,
		UePool:           make(map[int64]*UeContext),
		SessPool:         make(map[uint32]*SessionContext),
		SupportTAList:    make(map[string][]PlmnSupportItem),
		UpfInfoList:      make(map[string]*UpfInfo),
	}
}
