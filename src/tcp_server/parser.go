package tcp_server

import (
	"fmt"
	"radio_simulator/lib/nas/nasMessage"
	"radio_simulator/lib/openapi/models"
	"radio_simulator/src/logger"
	"radio_simulator/src/simulator_context"
	"radio_simulator/src/simulator_nas/nas_packet"
	"radio_simulator/src/simulator_ngap"
	"regexp"
	"strconv"
	"time"
)

var stringFormat = regexp.MustCompile(`\S+`)
var sessFormat = regexp.MustCompile(`dnn=([^\,]+),sst=([^\,]+),sd=(\S+)`)

// Parse cli Command from Tcp Server
// Usage: [ options ]
//
// options:
//   show [all|${PduSessionId}]
// 		show ue current state
//
//   reg {ip}
// 		register to CoreNetwork
// 		ip is specific Ran Ip to Connect(default={firstRanIp})
//
// 	 dereg
// 		deregister to CoreNetwork
//
//   sess i [add|del] {dnn=%s,sst=%d,sd=%s}
// 		pduSessionId i add or delete
//
// Output Format
//   show  [all|${PduSessionId}]
//		first line in all case:
//			"[SHOW] REGISTERED\n" or "[SHOW] REGISTERING\n" or "[SHOW] DEREGISTERED\n"
//		sessInfo:
//			"[SHOW] ID=%d,DNN=%s,SST=%d,SD=%s,UEIP=%s,ULAddr=%s,ULTEID=%d,DLAddr=%s,DLTEID=%d\n"
// 		all means show all Pdu Session Id
//
//   reg {ip}
// 		"[REG] SUCCESS\n" or
// 		"[REG] FAIL\n" or
//
// 	 dereg
// 		"[DEREG] SUCCESS\n" or
// 		"[DEREG] FAIL\n" or
//
//   sess i [add|del]
//		"[SESSION] ID=%d,DNN=%s,SST=%d,SD=%s,UEIP=%s,ULAddr=%s,ULTEID=%d,DLAddr=%s,DLTEID=%d\n" for add case, "[SESSION] DEL %d\n" for del case or
// 		"[SESSION] ADD/DEL %d FAIL\n"
//
func parseCmd(ue *simulator_context.UeContext, raddr string, cmd string) string {
	params := stringFormat.FindAllString(cmd, -1)
	cnt := len(params)
	if cnt == 0 {
		return ""
	}
	var msg string
	switch params[0] {
	case "show":
		if cnt == 1 {
			msg = "show missing action[all/{id}]"
		} else {
			switch params[1] {
			case "all":
				msg = "[SHOW] " + ue.RegisterState + "\n"
				for _, sess := range ue.PduSession {
					sessInfo := sess.GetTunnelMsg()
					if sessInfo == "" {
						continue
					}
					msg = msg + "[SHOW] " + sessInfo
				}
			default:
				id, err := strconv.Atoi(params[1])
				if err != nil {
					msg = "sess id is not digit"
					break
				}
				sess := ue.PduSession[int64(id)]
				if sess == nil {
					msg = "sess " + params[1] + " has not established yet"
					break
				}
				sessInfo := sess.GetTunnelMsg()
				if sessInfo == "" {
					msg = "sess " + params[1] + " is still establishing\n"
					break
				}
				msg = "[SHOW] " + sessInfo
			}
		}
	case "reg":
		if ue.RegisterState == simulator_context.RegisterStateRegistered {
			msg = "[REG] SUCCESS\n"
			break
		}
		ran := self.RanPool[self.DefaultRanSctpUri]
		if cnt > 1 {
			// Use Specific RanSctpUri
			ran = self.RanPool[params[1]]
			if ran == nil {
				msg = "ranIp " + params[1] + " does not exist"
				break
			}
		}
		if ue.RegisterState == simulator_context.RegisterStateDeregitered {
			// Send Registration Request
			ue.AttachRan(ran)
			ue.RegisterState = simulator_context.RegisterStateRegistering
			simulator_ngap.SendInitailUeMessage_RegistraionRequest(ran, ue)
		}
		msg = ReadChannelMsg(ue, raddr)
	case "dereg":
		if ue.RegisterState == simulator_context.RegisterStateDeregitered {
			msg = "[DEREG] SUCCESS\n"
		} else {
			// Send Deregistration Request
			nasPdu, err := nas_packet.GetDeregistrationRequest(ue, 0) //normoal release
			if err != nil {
				logger.TcpServerLog.Error(err.Error())
				msg = "[DEREG] FAIL\n"
				break
			}
			simulator_ngap.SendUplinkNasTransport(ue.Ran, ue, nasPdu)
			msg = ReadChannelMsg(ue, raddr)
		}
	case "sess":
		if ue.RegisterState != simulator_context.RegisterStateRegistered {
			msg = "need to registrate first"
			break
		}
		if cnt <= 2 {
			msg = "sess need id and action[add/del]"
			break
		}
		id, err := strconv.Atoi(params[1])
		if err != nil {
			msg = "sess id is not digit"
			break
		}
		sess := ue.PduSession[int64(id)]
		switch params[2] {
		case "add":
			if sess == nil {
				pduSessionId := uint8(id)
				dnn := "internet" // default
				snssai := ue.Nssai.DefaultSingleNssais[0]
				snssaiInfo, exist := ue.SmfSelData.SubscribedSnssaiInfos[fmt.Sprintf("%2d%s", snssai.Sst, snssai.Sd)]
				if exist {
					dnn = snssaiInfo.DnnInfos[0].Dnn
				}
				if cnt > 3 {
					match := sessFormat.FindStringSubmatch(params[3])
					if match == nil {
						msg = "sess parameter " + params[3] + " not match format \"dnn=%s,sst=%d,sd=%s\""
						break
					}
					dnn = match[1]
					sst, _ := strconv.ParseInt(match[2], 10, 32)
					snssai = models.Snssai{
						Sst: int32(sst),
						Sd:  match[3],
					}
				}
				// Send Pdu Session Estblishment
				gsmPdu, err := nas_packet.GetPduSessionEstablishmentRequest(pduSessionId, nasMessage.PDUSessionTypeIPv4)
				if err != nil {
					logger.TcpServerLog.Error(err.Error())
					msg = fmt.Sprintf("[SESSION] ADD %d FAIL\n", pduSessionId)
					break
				}
				nasPdu, err := nas_packet.GetUlNasTransport_PduSessionEstablishmentRequest(ue, pduSessionId, nasMessage.ULNASTransportRequestTypeInitialRequest, dnn, &snssai, gsmPdu)
				if err != nil {
					logger.TcpServerLog.Error(err.Error())
					msg = fmt.Sprintf("[SESSION] ADD %d FAIL\n", pduSessionId)
					break
				}
				simulator_ngap.SendUplinkNasTransport(ue.Ran, ue, nasPdu)
				sess := ue.AddPduSession(pduSessionId, dnn, snssai)
				msg = ReadSessChannelMsg(sess)
				break
			}
			sessInfo := sess.GetTunnelMsg()
			if sessInfo == "" {
				msg = "sess " + params[1] + " is still establishing\n"
				break
			}
			msg = "[SESSION] " + sessInfo
		case "del":
			if sess == nil {
				msg = "[SESSION] DEL " + params[1] + "\n"
				break
			} else {
				// TODO: Send Pdu Session Release
				pduSessionId := uint8(id)
				nasPdu, err := nas_packet.GetUlNasTransport_PduSessionCommonData(ue, pduSessionId, nas_packet.PDUSesRelReq)
				if err != nil {
					logger.TcpServerLog.Error(err.Error())
					msg = fmt.Sprintf("[SESSION] DEL %d FAIL\n", pduSessionId)
					break
				}
				simulator_ngap.SendUplinkNasTransport(ue.Ran, ue, nasPdu)
				msg = ReadSessChannelMsg(sess)
			}
		default:
			msg = "sess action is not [add/del]"
		}
	}
	return msg

}

func ReadChannelMsg(ue *simulator_context.UeContext, raddr string) (msg string) {
	mtx.Lock()
	ue.TcpChannelMsg[raddr] = make(chan string)
	mtx.Unlock()
	select {
	case msg = <-ue.TcpChannelMsg[raddr]:
	case <-time.After(5 * time.Second):
		msg = fmt.Sprintf("[TIMEOUT]\n")
	}
	mtx.Lock()
	delete(ue.TcpChannelMsg, raddr)
	mtx.Unlock()
	return
}

func ReadSessChannelMsg(sess *simulator_context.SessionContext) string {
	select {
	case msg := <-sess.SessTcpChannelMsg:
		return msg
	case <-time.After(5 * time.Second):
		return fmt.Sprintf("[TIMEOUT]\n")
	}
}
