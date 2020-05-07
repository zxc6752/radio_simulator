package simulator_nas

import (
	"github.com/sirupsen/logrus"
	"radio_simulator/lib/nas"
	"radio_simulator/lib/nas/nasMessage"
	"radio_simulator/src/logger"
	"radio_simulator/src/simulator_context"
	"radio_simulator/src/simulator_nas/nas_security"
)

var nasLog *logrus.Entry

func init() {
	nasLog = logger.NasLog
}

func checkMsgError(err error, msg string) {
	if err != nil {
		nasLog.Errorf("Handle %s Error: %s", msg, err.Error())
	}
}

func HandleNAS(ue *simulator_context.UeContext, nasPdu []byte) {

	if ue == nil {
		nasLog.Error("Ue is nil")
		return
	}

	if nasPdu == nil {
		nasLog.Error("nasPdu is nil")
		return
	}

	if nas.GetEPD(nasPdu) == nasMessage.Epd5GSSessionManagementMessage {
		// GSM Message
		msg := new(nas.Message)
		err := msg.PlainNasDecode(&nasPdu)
		if err != nil {
			nasLog.Error(err.Error())
			return
		}
		switch msg.GsmMessage.GetMessageType() {
		case nas.MsgTypePDUSessionEstablishmentAccept:
			checkMsgError(HandlePduSessionEstblishmentAccept(ue, msg.GsmMessage.PDUSessionEstablishmentAccept), "PduSessionEstblishmentAccept")
		case nas.MsgTypePDUSessionReleaseCommand:
			checkMsgError(HandlePduSessionReleaseCommand(ue, msg.GsmMessage.PDUSessionReleaseCommand), "PduSessionReleaseCommand")
		default:
			nasLog.Errorf("Unknown GsmMessage[%d]\n", msg.GsmMessage.GetMessageType())
		}
		return
	}

	// GMM Message
	msg, err := nas_security.NASDecode(ue, nas.GetSecurityHeaderType(nasPdu)&0x0f, nasPdu)
	if err != nil {
		nasLog.Error(err.Error())
		return
	}

	switch msg.GmmMessage.GetMessageType() {
	case nas.MsgTypeAuthenticationRequest:
		checkMsgError(HandleAuthenticationRequest(ue, msg.GmmMessage.AuthenticationRequest), "AuthenticationRequest")
	case nas.MsgTypeSecurityModeCommand:
		checkMsgError(HandleSecurityModeCommand(ue, msg.GmmMessage.SecurityModeCommand), "SecurityModeCommand")
	case nas.MsgTypeRegistrationAccept:
		checkMsgError(HandleRegistrationAccept(ue, msg.GmmMessage.RegistrationAccept), "RegistrationAccept")
	case nas.MsgTypeDeregistrationAcceptUEOriginatingDeregistration:
		checkMsgError(HandleDeregistrationAccept(ue, msg.GmmMessage.DeregistrationAcceptUEOriginatingDeregistration), "DeregistraionAccept")
	case nas.MsgTypeDLNASTransport:
		checkMsgError(HandleDLNASTransport(ue, msg.GmmMessage.DLNASTransport), "DLNASTransport")
	default:
		nasLog.Errorf("Unknown GmmMessage[%d]\n", msg.GmmMessage.GetMessageType())
	}

	return

}
