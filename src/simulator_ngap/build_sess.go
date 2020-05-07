package simulator_ngap

import (
	"encoding/binary"
	"radio_simulator/lib/aper"
	"radio_simulator/lib/ngap/ngapConvert"
	"radio_simulator/lib/ngap/ngapType"
	"radio_simulator/src/simulator_context"
)

func BuildPDUSessionResourceSetupUnsuccessfulTransfer(cause ngapType.Cause, criticalityDiagnostics *ngapType.CriticalityDiagnostics) ([]byte, error) {

	transfer := ngapType.PDUSessionResourceSetupUnsuccessfulTransfer{}

	// Cause
	transfer.Cause = cause

	// Criticality Diagnostics (optional)
	if criticalityDiagnostics != nil {
		transfer.CriticalityDiagnostics = criticalityDiagnostics
	}

	return aper.MarshalWithParams(transfer, "valueExt")
}

func BuildPDUSessionResourceSetupResponseTransfer(sess *simulator_context.SessionContext) ([]byte, error) {

	transfer := ngapType.PDUSessionResourceSetupResponseTransfer{}

	// QOS Flow Per TNL Information
	qosFlowPerTNLInformation := &transfer.QosFlowPerTNLInformation
	qosFlowPerTNLInformation.UPTransportLayerInformation.Present = ngapType.UPTransportLayerInformationPresentGTPTunnel
	qosFlowPerTNLInformation.UPTransportLayerInformation.GTPTunnel = new(ngapType.GTPTunnel)
	teid := make([]byte, 4)
	binary.BigEndian.PutUint32(teid, sess.DLTEID)
	qosFlowPerTNLInformation.UPTransportLayerInformation.GTPTunnel.GTPTEID.Value = teid
	qosFlowPerTNLInformation.UPTransportLayerInformation.GTPTunnel.TransportLayerAddress = ngapConvert.IPAddressToNgap(sess.DLAddr, "") // Only Support Ipv4

	for qfi, _ := range sess.QosFlows {
		item := ngapType.AssociatedQosFlowItem{
			QosFlowIdentifier: ngapType.QosFlowIdentifier{
				Value: qfi,
			},
		}
		qosFlowPerTNLInformation.AssociatedQosFlowList.List = append(qosFlowPerTNLInformation.AssociatedQosFlowList.List, item)
	}

	return aper.MarshalWithParams(transfer, "valueExt")
}

func BuildPDUSessionResourceReleaseResponseTransfer() ([]byte, error) {

	transfer := ngapType.PDUSessionResourceReleaseResponseTransfer{}

	return aper.MarshalWithParams(transfer, "valueExt")
}

func AppendPDUSessionResourceSetupListSURes(list *ngapType.PDUSessionResourceSetupListSURes, pduSessionID int64, transfer []byte) {
	item := ngapType.PDUSessionResourceSetupItemSURes{}
	item.PDUSessionID.Value = pduSessionID
	item.PDUSessionResourceSetupResponseTransfer = transfer
	list.List = append(list.List, item)
}

func AppendPDUSessionResourceFailedToSetupListSURes(list *ngapType.PDUSessionResourceFailedToSetupListSURes, pduSessionID int64, transfer []byte) {
	item := ngapType.PDUSessionResourceFailedToSetupItemSURes{}
	item.PDUSessionID.Value = pduSessionID
	item.PDUSessionResourceSetupUnsuccessfulTransfer = transfer
	list.List = append(list.List, item)
}

func AppendPDUSessionResourceReleasedListRelRes(list *ngapType.PDUSessionResourceReleasedListRelRes, pduSessionID int64, transfer []byte) {
	item := ngapType.PDUSessionResourceReleasedItemRelRes{}
	item.PDUSessionID.Value = pduSessionID
	item.PDUSessionResourceReleaseResponseTransfer = transfer
	list.List = append(list.List, item)
}
