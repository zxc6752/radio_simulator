package simulator_ngap

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os/exec"
	"radio_simulator/lib/nas/nasMessage"
	"radio_simulator/lib/ngap"
	"radio_simulator/lib/ngap/ngapType"
	"radio_simulator/src/logger"
	"radio_simulator/src/simulator_context"
)

var ngapLog *logrus.Entry

func init() {
	ngapLog = logger.NgapLog
}

func check(err error) {
	if err != nil {
		logger.InitLog.Error(err.Error())
	}
}

func SendNGSetupRequest(ran *simulator_context.RanContext) {
	var n int
	var recvMsg = make([]byte, 2048)

	// send NGSetupRequest Msg
	pkt, err := BuildNGSetupRequest(ran)
	if err != nil {
		ngapLog.Errorf("Build NGSetUp failed : %s", err.Error())
		return
	}
	SendToAmf(ran, pkt)

	// receive NGSetupResponse Msg
	n, err = ran.SctpConn.Read(recvMsg)
	check(err)
	pdu, err := ngap.Decoder(recvMsg[:n])
	check(err)
	if pdu.Present != ngapType.NGAPPDUPresentSuccessfulOutcome {
		logger.NgapLog.Error("NG SetUp Fail!!!!")
	}
	return
}

func SendInitailUeMessage_RegistraionRequest(ran *simulator_context.RanContext, ue *simulator_context.UeContext) {
	ngapLog.Info("[RAN] Initail Ue Message (Initail Registration Request)")
	pkt, err := BuildInitialUEMessage(ue, nasMessage.RegistrationType5GSInitialRegistration, "")
	if err != nil {
		ngapLog.Errorf("Build InitialUEMessage failed : %s", err.Error())
		return
	}
	SendToAmf(ran, pkt)
}

func SendUplinkNasTransport(ran *simulator_context.RanContext, ue *simulator_context.UeContext, nasPdu []byte) {

	ngapLog.Info("[RAN] Send Uplink Nas Transport")

	pkt, err := BuildUplinkNasTransport(ue, nasPdu)
	if err != nil {
		ngapLog.Errorf("Build Uplink Nas Transport failed : %s", err.Error())
		return
	}
	SendToAmf(ran, pkt)
}

func SendIntialContextSetupResponse(ran *simulator_context.RanContext, ue *simulator_context.UeContext, pduSessionIds []string) {

	ngapLog.Info("[RAN] Send Intial Context Setup Response")

	pkt, err := BuildInitialContextSetupResponse(ue, pduSessionIds, nil)
	if err != nil {
		ngapLog.Errorf("Build Uplink Nas Transport failed : %s", err.Error())
		return
	}
	SendToAmf(ran, pkt)

}

func SendUeContextReleaseComplete(ran *simulator_context.RanContext, ue *simulator_context.UeContext) {

	ngapLog.Info("[RAN] Send Ue Context Release Complete")

	pkt, err := BuildUEContextReleaseComplete(ue)
	if err != nil {
		ngapLog.Errorf("Build Ue Context Release Complete failed : %s", err.Error())
		return
	}

	// Reset Ue Context
	ue.AmfUeNgapId = simulator_context.AmfNgapIdUnspecified
	for _, sess := range ue.PduSession {
		sess.Remove()
		if sess.UeIp != "" {
			exec.Command("ip", "addr", "del", sess.UeIp, "dev", "lo").Output()
		}
	}

	SendToAmf(ran, pkt)
	if ue.RegisterState == simulator_context.RegisterStateDeregitered {
		// Complete Deregistration
		ue.SendMsg("[DEREG] SUCCESS\n")
	}
}

func SendPDUSessionResourceSetupResponse(
	ran *simulator_context.RanContext,
	ue *simulator_context.UeContext,
	responseList *ngapType.PDUSessionResourceSetupListSURes,
	failedListSURes *ngapType.PDUSessionResourceFailedToSetupListSURes) {

	ngapLog.Infoln("[RAN] Send PDU Session Resource Setup Response")

	pkt, err := BuildPDUSessionResourceSetupResponse(ue, responseList, failedListSURes)
	if err != nil {
		ngapLog.Errorf("Build PDU Session Resource Setup Response failed : %+v", err)
		return
	}

	SendToAmf(ran, pkt)
	// Send Callback To Tcp Client
	if responseList != nil {
		for _, item := range responseList.List {
			sess := ue.PduSession[item.PDUSessionID.Value]
			msg := "[SESSION] " + sess.GetTunnelMsg()
			sess.SendMsg(msg)
		}
	}
	if failedListSURes != nil {
		for _, item := range failedListSURes.List {
			sess := ue.PduSession[item.PDUSessionID.Value]
			if sess != nil {
				msg := fmt.Sprintf("[SESSION] ADD %d FAIL\n", item.PDUSessionID.Value)
				sess.SendMsg(msg)
				sess.Remove()
			}
		}
	}
}

func SendPDUSessionResourceReleaseResponse(
	ran *simulator_context.RanContext,
	ue *simulator_context.UeContext,
	relList ngapType.PDUSessionResourceReleasedListRelRes,
	diagnostics *ngapType.CriticalityDiagnostics) {

	ngapLog.Infoln("[EAN] Send PDU Session Resource Release Response")

	if len(relList.List) < 1 {
		ngapLog.Errorln("PDUSessionResourceReleasedListRelRes is nil. This message shall contain at least one Item")
		return
	}

	pkt, err := BuildPDUSessionResourceReleaseResponse(ue, relList, diagnostics)
	if err != nil {
		ngapLog.Errorf("Build PDU Session Resource Release Response failed : %+v", err)
		return
	}

	SendToAmf(ran, pkt)
}

func SendErrorIndication(ran *simulator_context.RanContext, amfUeNgapId, ranUeNgapId *int64, cause *ngapType.Cause, criticalityDiagnostics *ngapType.CriticalityDiagnostics) {

	ngapLog.Info("[AMF] Send Error Indication")

	pkt, err := BuildErrorIndication(amfUeNgapId, ranUeNgapId, cause, criticalityDiagnostics)
	if err != nil {
		ngapLog.Errorf("Build ErrorIndication failed : %s", err.Error())
		return
	}
	SendToAmf(ran, pkt)
}

func SendToAmf(ran *simulator_context.RanContext, message []byte) {
	_, err := ran.SctpConn.Write(message)
	check(err)
}
