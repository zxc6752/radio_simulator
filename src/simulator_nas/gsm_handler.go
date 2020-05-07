package simulator_nas

import (
	"fmt"
	"net"
	"os/exec"
	"radio_simulator/lib/nas/nasConvert"
	"radio_simulator/lib/nas/nasMessage"
	"radio_simulator/src/simulator_context"
	"radio_simulator/src/simulator_nas/nas_packet"
	"radio_simulator/src/simulator_ngap"
)

func HandlePduSessionEstblishmentAccept(ue *simulator_context.UeContext, request *nasMessage.PDUSessionEstablishmentAccept) error {

	nasLog.Infof("UE[%s] Handle PDU Session Establishment Accept", ue.Supi)

	pduSessionId := int64(request.GetPDUSessionID())
	sess, exist := ue.PduSession[pduSessionId]
	if !exist {
		return fmt.Errorf("pduSessionId[%d] is not exist in UE", pduSessionId)
	}
	if request.DNN != nil {
		sess.Dnn = string(request.GetDNN())
	}
	if request.SNSSAI != nil {
		sess.Snssai = nasConvert.SnssaiToModels(request.SNSSAI)
	}
	if request.PDUAddress != nil {
		ipBytes := request.PDUAddress.GetPDUAddressInformation()
		switch request.PDUAddress.GetPDUSessionTypeValue() {
		case nasMessage.PDUSessionTypeIPv4:
			sess.UeIp = net.IP(ipBytes[:4]).String()
			_, err := exec.Command("ip", "addr", "add", sess.UeIp, "dev", "lo").Output()
			if err != nil {
				return fmt.Errorf("Create ue addr failed[%s]", err.Error())
			}
		case nasMessage.PDUSessionTypeIPv6, nasMessage.PDUSessionTypeIPv4IPv6:
			return fmt.Errorf("Ipv6 is not support yet")
		}
	}
	return nil
}

func HandlePduSessionReleaseCommand(ue *simulator_context.UeContext, request *nasMessage.PDUSessionReleaseCommand) error {

	nasLog.Infof("UE[%s] Handle PDU Session Release Command", ue.Supi)

	pduSessionId := request.GetPDUSessionID()
	sess, exist := ue.PduSession[int64(pduSessionId)]
	if !exist {
		return fmt.Errorf("pduSessionId[%d] is not exist in UE", pduSessionId)
	}
	// Send Pdu Session Release Complete to SMF
	nasPdu, err := nas_packet.GetUlNasTransport_PduSessionCommonData(ue, pduSessionId, nas_packet.PDUSesRelCmp)
	if err != nil {
		return err
	}
	simulator_ngap.SendUplinkNasTransport(ue.Ran, ue, nasPdu)
	// Send Nootify Msg to UE
	sess.SendMsg(fmt.Sprintf("[SESSION] DEL %d SUCCESS\n", pduSessionId))
	sess.Remove()
	if sess.UeIp != "" {
		_, err := exec.Command("ip", "addr", "del", sess.UeIp, "dev", "lo").Output()
		if err != nil {
			return fmt.Errorf("Delete ue addr failed[%s]", err.Error())
		}
	}
	return nil
}
