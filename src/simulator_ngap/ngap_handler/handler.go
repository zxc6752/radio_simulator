package ngap_handler

import (
	"radio_simulator/lib/aper"
	"radio_simulator/lib/ngap/ngapType"
	"radio_simulator/src/simulator_context"
	"radio_simulator/src/simulator_nas"
	"radio_simulator/src/simulator_ngap"
)

func HandleDownlinkNASTransport(ran *simulator_context.RanContext, message *ngapType.NGAPPDU) {
	var aMFUENGAPID *ngapType.AMFUENGAPID
	var rANUENGAPID *ngapType.RANUENGAPID
	// var oldAMF *ngapType.AMFName
	// var rANPagingPriority *ngapType.RANPagingPriority
	var nASPDU *ngapType.NASPDU
	// var mobilityRestrictionList *ngapType.MobilityRestrictionList
	// var indexToRFSP *ngapType.IndexToRFSP
	// var uEAggregateMaximumBitRate *ngapType.UEAggregateMaximumBitRate
	// var allowedNSSAI *ngapType.AllowedNSSAI

	var iesCriticalityDiagnostics ngapType.CriticalityDiagnosticsIEList

	if ran == nil {
		ngapLog.Error("RAN Context is nil")
		return
	}

	if message == nil {
		ngapLog.Error("NGAP Message is nil")
		return
	}

	initiatingMessage := message.InitiatingMessage
	if initiatingMessage == nil {
		ngapLog.Error("InitiatingMessage is nil")
		return
	}

	downlinkNASTransport := initiatingMessage.Value.DownlinkNASTransport
	if downlinkNASTransport == nil {
		ngapLog.Error("downlinkNASTransport is nil")
		return
	}

	for _, ie := range downlinkNASTransport.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			ngapLog.Traceln("[NGAP] Decode IE AMFUENGAPID")
			aMFUENGAPID = ie.Value.AMFUENGAPID
			if aMFUENGAPID == nil {
				ngapLog.Error("AMFUENGAPID is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ie.Id.Value, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			ngapLog.Traceln("[NGAP] Decode IE RANUENGAPID")
			rANUENGAPID = ie.Value.RANUENGAPID
			if rANUENGAPID == nil {
				ngapLog.Error("RANUENGAPID is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ie.Id.Value, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		case ngapType.ProtocolIEIDOldAMF:
			ngapLog.Traceln("[NGAP] Decode IE OldAMF")
			// oldAMF = ie.Value.OldAMF
		case ngapType.ProtocolIEIDRANPagingPriority:
			ngapLog.Traceln("[NGAP] Decode IE RANPagingPriority")
			// rANPagingPriority = ie.Value.RANPagingPriority
		case ngapType.ProtocolIEIDNASPDU:
			ngapLog.Traceln("[NGAP] Decode IE NASPDU")
			nASPDU = ie.Value.NASPDU
			if nASPDU == nil {
				ngapLog.Error("NASPDU is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ie.Id.Value, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		case ngapType.ProtocolIEIDMobilityRestrictionList:
			ngapLog.Traceln("[NGAP] Decode IE MobilityRestrictionList")
			// mobilityRestrictionList = ie.Value.MobilityRestrictionList
		case ngapType.ProtocolIEIDIndexToRFSP:
			ngapLog.Traceln("[NGAP] Decode IE IndexToRFSP")
			// indexToRFSP = ie.Value.IndexToRFSP
		case ngapType.ProtocolIEIDUEAggregateMaximumBitRate:
			ngapLog.Traceln("[NGAP] Decode IE UEAggregateMaximumBitRate")
			// uEAggregateMaximumBitRate = ie.Value.UEAggregateMaximumBitRate
		case ngapType.ProtocolIEIDAllowedNSSAI:
			ngapLog.Traceln("[NGAP] Decode IE AllowedNSSAI")
			// allowedNSSAI = ie.Value.AllowedNSSAI
		}
	}

	if len(iesCriticalityDiagnostics.List) > 0 {
		procudureCode := ngapType.ProcedureCodeDownlinkNASTransport
		trigger := ngapType.TriggeringMessagePresentInitiatingMessage
		criticality := ngapType.CriticalityPresentIgnore
		criticalityDiagnostics := buildCriticalityDiagnostics(&procudureCode, &trigger, &criticality, &iesCriticalityDiagnostics)
		simulator_ngap.SendErrorIndication(ran, nil, nil, nil, &criticalityDiagnostics)
		return
	}

	var ue *simulator_context.UeContext
	if rANUENGAPID != nil {
		ue = ran.FindUeByRanUeNgapID(rANUENGAPID.Value)
		if ue == nil {
			ngapLog.Warnf("No UE Context[RanUeNgapID:%d]\n", rANUENGAPID.Value)
			return
		}
	}

	if aMFUENGAPID != nil {
		if ue.AmfUeNgapId == simulator_context.AmfNgapIdUnspecified {
			ngapLog.Tracef("Create new logical UE-associated NG-connection")
			ue.AmfUeNgapId = aMFUENGAPID.Value
			// n3iwfUe.SCTPAddr = amf.SCTPAddr
		} else {
			if ue.AmfUeNgapId != aMFUENGAPID.Value {
				ngapLog.Warn("AMFUENGAPID unmatched")
				return
			}
		}
	}

	if nASPDU != nil {
		simulator_nas.HandleNAS(ue, nASPDU.Value)
	}
}

func HandleInitialContextSetupRequest(ran *simulator_context.RanContext, message *ngapType.NGAPPDU) {
	var aMFUENGAPID *ngapType.AMFUENGAPID
	var rANUENGAPID *ngapType.RANUENGAPID
	// var oldAMF *ngapType.AMFName
	var uEAggregateMaximumBitRate *ngapType.UEAggregateMaximumBitRate
	// var coreNetworkAssistanceInformation *ngapType.CoreNetworkAssistanceInformation
	var gUAMI *ngapType.GUAMI
	// var pDUSessionResourceSetupListCxtReq *ngapType.PDUSessionResourceSetupListCxtReq
	var allowedNSSAI *ngapType.AllowedNSSAI
	var uESecurityCapabilities *ngapType.UESecurityCapabilities
	var securityKey *ngapType.SecurityKey
	// var traceActivation *ngapType.TraceActivation
	// var mobilityRestrictionList *ngapType.MobilityRestrictionList
	// var uERadioCapability *ngapType.UERadioCapability
	// var indexToRFSP *ngapType.IndexToRFSP
	// var maskedIMEISV *ngapType.MaskedIMEISV
	var nASPDU *ngapType.NASPDU
	// var emergencyFallbackIndicator *ngapType.EmergencyFallbackIndicator
	// var rRCInactiveTransitionReportRequest *ngapType.RRCInactiveTransitionReportRequest
	// var uERadioCapabilityForPaging *ngapType.UERadioCapabilityForPaging

	var iesCriticalityDiagnostics ngapType.CriticalityDiagnosticsIEList

	if ran == nil {
		ngapLog.Error("RAN Context is nil")
		return
	}

	if message == nil {
		ngapLog.Error("NGAP Message is nil")
		return
	}

	initiatingMessage := message.InitiatingMessage
	if initiatingMessage == nil {
		ngapLog.Error("InitiatingMessage is nil")
		return
	}

	initialContextSetupRequest := initiatingMessage.Value.InitialContextSetupRequest
	if initialContextSetupRequest == nil {
		ngapLog.Error("initialContextSetupRequest is nil")
		return
	}

	for _, ie := range initialContextSetupRequest.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			ngapLog.Traceln("[NGAP] Decode IE AMFUENGAPID")
			aMFUENGAPID = ie.Value.AMFUENGAPID
			if aMFUENGAPID == nil {
				ngapLog.Error("AMFUENGAPID is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ie.Id.Value, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			ngapLog.Traceln("[NGAP] Decode IE RANUENGAPID")
			rANUENGAPID = ie.Value.RANUENGAPID
			if rANUENGAPID == nil {
				ngapLog.Error("RANUENGAPID is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ie.Id.Value, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		case ngapType.ProtocolIEIDOldAMF:
			ngapLog.Traceln("[NGAP] Decode IE OldAMF")
			// oldAMF = ie.Value.OldAMF
		case ngapType.ProtocolIEIDUEAggregateMaximumBitRate:
			ngapLog.Traceln("[NGAP] Decode IE UEAggregateMaximumBitRate")
			uEAggregateMaximumBitRate = ie.Value.UEAggregateMaximumBitRate
			if uEAggregateMaximumBitRate == nil {
				ngapLog.Error("UEAggregateMaximumBitRate is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ie.Id.Value, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		case ngapType.ProtocolIEIDCoreNetworkAssistanceInformation:
			ngapLog.Traceln("[NGAP] Decode IE CoreNetworkAssistanceInformation")
			// coreNetworkAssistanceInformation = ie.Value.CoreNetworkAssistanceInformation
		case ngapType.ProtocolIEIDGUAMI:
			ngapLog.Traceln("[NGAP] Decode IE GUAMI")
			gUAMI = ie.Value.GUAMI
			if gUAMI == nil {
				ngapLog.Error("GUAMI is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ie.Id.Value, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		case ngapType.ProtocolIEIDPDUSessionResourceSetupListCxtReq:
			ngapLog.Traceln("[NGAP] Decode IE PDUSessionResourceSetupListCxtReq")
			// pDUSessionResourceSetupListCxtReq = ie.Value.PDUSessionResourceSetupListCxtReq
		case ngapType.ProtocolIEIDAllowedNSSAI:
			ngapLog.Traceln("[NGAP] Decode IE AllowedNSSAI")
			allowedNSSAI = ie.Value.AllowedNSSAI
			if allowedNSSAI == nil {
				ngapLog.Error("AllowedNSSAI is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ie.Id.Value, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		case ngapType.ProtocolIEIDUESecurityCapabilities:
			ngapLog.Traceln("[NGAP] Decode IE UESecurityCapabilities")
			uESecurityCapabilities = ie.Value.UESecurityCapabilities
			if uESecurityCapabilities == nil {
				ngapLog.Error("UESecurityCapabilities is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ie.Id.Value, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		case ngapType.ProtocolIEIDSecurityKey:
			ngapLog.Traceln("[NGAP] Decode IE SecurityKey")
			securityKey = ie.Value.SecurityKey
			if securityKey == nil {
				ngapLog.Error("SecurityKey is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ie.Id.Value, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		case ngapType.ProtocolIEIDTraceActivation:
			ngapLog.Traceln("[NGAP] Decode IE TraceActivation")
			// traceActivation = ie.Value.TraceActivation
		case ngapType.ProtocolIEIDMobilityRestrictionList:
			ngapLog.Traceln("[NGAP] Decode IE MobilityRestrictionList")
			// mobilityRestrictionList = ie.Value.MobilityRestrictionList
		case ngapType.ProtocolIEIDUERadioCapability:
			ngapLog.Traceln("[NGAP] Decode IE UERadioCapability")
			// uERadioCapability = ie.Value.UERadioCapability
		case ngapType.ProtocolIEIDIndexToRFSP:
			ngapLog.Traceln("[NGAP] Decode IE IndexToRFSP")
			// indexToRFSP = ie.Value.IndexToRFSP
		case ngapType.ProtocolIEIDMaskedIMEISV:
			ngapLog.Traceln("[NGAP] Decode IE MaskedIMEISV")
			// maskedIMEISV = ie.Value.MaskedIMEISV
		case ngapType.ProtocolIEIDNASPDU:
			ngapLog.Traceln("[NGAP] Decode IE NASPDU")
			nASPDU = ie.Value.NASPDU
		case ngapType.ProtocolIEIDEmergencyFallbackIndicator:
			ngapLog.Traceln("[NGAP] Decode IE EmergencyFallbackIndicator")
			// emergencyFallbackIndicator = ie.Value.EmergencyFallbackIndicator
		case ngapType.ProtocolIEIDRRCInactiveTransitionReportRequest:
			ngapLog.Traceln("[NGAP] Decode IE RRCInactiveTransitionReportRequest")
			// rRCInactiveTransitionReportRequest = ie.Value.RRCInactiveTransitionReportRequest
		case ngapType.ProtocolIEIDUERadioCapabilityForPaging:
			ngapLog.Traceln("[NGAP] Decode IE UERadioCapabilityForPaging")
			// uERadioCapabilityForPaging = ie.Value.UERadioCapabilityForPaging
		}
	}

	if len(iesCriticalityDiagnostics.List) > 0 {
		procudureCode := ngapType.ProcedureCodeInitialContextSetup
		trigger := ngapType.TriggeringMessagePresentInitiatingMessage
		criticality := ngapType.CriticalityPresentReject
		criticalityDiagnostics := buildCriticalityDiagnostics(&procudureCode, &trigger, &criticality, &iesCriticalityDiagnostics)
		simulator_ngap.SendErrorIndication(ran, nil, nil, nil, &criticalityDiagnostics)
		return
	}

	ue := ran.FindUeByRanUeNgapID(rANUENGAPID.Value)
	if ue == nil {
		ngapLog.Warnf("No UE Context[RanUeNgapID:%d]\n", rANUENGAPID.Value)
		return
	}

	// TODO: Service Request Case
	simulator_ngap.SendIntialContextSetupResponse(ran, ue, nil)

	if nASPDU != nil {
		simulator_nas.HandleNAS(ue, nASPDU.Value)
	}

}

func HandleUeContextReleaseCommand(ran *simulator_context.RanContext, message *ngapType.NGAPPDU) {
	var uENGAPIDs *ngapType.UENGAPIDs
	var cause *ngapType.Cause

	var iesCriticalityDiagnostics ngapType.CriticalityDiagnosticsIEList

	if ran == nil {
		ngapLog.Error("RAN Context is nil")
		return
	}

	if message == nil {
		ngapLog.Error("NGAP Message is nil")
		return
	}

	initiatingMessage := message.InitiatingMessage
	if initiatingMessage == nil {
		ngapLog.Error("InitiatingMessage is nil")
		return
	}

	uEContextReleaseCommand := initiatingMessage.Value.UEContextReleaseCommand
	if uEContextReleaseCommand == nil {
		ngapLog.Error("uEContextReleaseCommand is nil")
		return
	}

	for _, ie := range uEContextReleaseCommand.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDUENGAPIDs:
			ngapLog.Traceln("[NGAP] Decode IE UENGAPIDs")
			uENGAPIDs = ie.Value.UENGAPIDs
			if uENGAPIDs == nil {
				ngapLog.Error("UENGAPIDs is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ie.Id.Value, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		case ngapType.ProtocolIEIDCause:
			ngapLog.Traceln("[NGAP] Decode IE Cause")
			cause = ie.Value.Cause
		}
	}

	if len(iesCriticalityDiagnostics.List) > 0 {
		procudureCode := ngapType.ProcedureCodeUEContextRelease
		trigger := ngapType.TriggeringMessagePresentInitiatingMessage
		criticality := ngapType.CriticalityPresentReject
		criticalityDiagnostics := buildCriticalityDiagnostics(&procudureCode, &trigger, &criticality, &iesCriticalityDiagnostics)
		simulator_ngap.SendErrorIndication(ran, nil, nil, nil, &criticalityDiagnostics)
		return
	}

	var ue *simulator_context.UeContext

	switch uENGAPIDs.Present {
	case ngapType.UENGAPIDsPresentAMFUENGAPID:
		ue = ran.FindUeByAmfUeNgapID(uENGAPIDs.AMFUENGAPID.Value)
		if ue == nil {
			ngapLog.Warnf("No UE Context[AmfUeNgapID:%d]\n", uENGAPIDs.AMFUENGAPID.Value)
			return
		}
	case ngapType.UENGAPIDsPresentUENGAPIDPair:
		pair := uENGAPIDs.UENGAPIDPair
		ue = ran.FindUeByRanUeNgapID(pair.RANUENGAPID.Value)
		if ue == nil {
			ngapLog.Warnf("No UE Context[RanUeNgapID:%d]\n", pair.RANUENGAPID.Value)
			return
		}
	}

	printAndGetCause(cause)
	simulator_ngap.SendUeContextReleaseComplete(ran, ue)
}

func HandlePduSessionResourceSetupRequest(ran *simulator_context.RanContext, message *ngapType.NGAPPDU) {
	var aMFUENGAPID *ngapType.AMFUENGAPID
	var rANUENGAPID *ngapType.RANUENGAPID
	// var rANPagingPriority *ngapType.RANPagingPriority
	var nASPDU *ngapType.NASPDU
	var pDUSessionResourceSetupListSUReq *ngapType.PDUSessionResourceSetupListSUReq

	var iesCriticalityDiagnostics ngapType.CriticalityDiagnosticsIEList

	if ran == nil {
		ngapLog.Error("RAN Context is nil")
		return
	}

	if message == nil {
		ngapLog.Error("NGAP Message is nil")
		return
	}

	initiatingMessage := message.InitiatingMessage
	if initiatingMessage == nil {
		ngapLog.Error("InitiatingMessage is nil")
		return
	}

	pDUSessionResourceSetupRequest := initiatingMessage.Value.PDUSessionResourceSetupRequest
	if pDUSessionResourceSetupRequest == nil {
		ngapLog.Error("pDUSessionResourceSetupRequest is nil")
		return
	}

	for _, ie := range pDUSessionResourceSetupRequest.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			ngapLog.Traceln("[NGAP] Decode IE AMFUENGAPID")
			aMFUENGAPID = ie.Value.AMFUENGAPID
			if aMFUENGAPID == nil {
				ngapLog.Error("AMFUENGAPID is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ie.Id.Value, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			ngapLog.Traceln("[NGAP] Decode IE RANUENGAPID")
			rANUENGAPID = ie.Value.RANUENGAPID
			if rANUENGAPID == nil {
				ngapLog.Error("RANUENGAPID is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ie.Id.Value, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		case ngapType.ProtocolIEIDRANPagingPriority:
			ngapLog.Traceln("[NGAP] Decode IE RANPagingPriority")
			// rANPagingPriority = ie.Value.RANPagingPriority
		case ngapType.ProtocolIEIDNASPDU:
			ngapLog.Traceln("[NGAP] Decode IE NASPDU")
			nASPDU = ie.Value.NASPDU
		case ngapType.ProtocolIEIDPDUSessionResourceSetupListSUReq:
			ngapLog.Traceln("[NGAP] Decode IE PDUSessionResourceSetupListSUReq")
			pDUSessionResourceSetupListSUReq = ie.Value.PDUSessionResourceSetupListSUReq
			if pDUSessionResourceSetupListSUReq == nil {
				ngapLog.Error("PDUSessionResourceSetupListSUReq is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ie.Id.Value, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		}
	}

	if len(iesCriticalityDiagnostics.List) > 0 {
		procudureCode := ngapType.ProcedureCodePDUSessionResourceSetup
		trigger := ngapType.TriggeringMessagePresentInitiatingMessage
		criticality := ngapType.CriticalityPresentReject
		criticalityDiagnostics := buildCriticalityDiagnostics(&procudureCode, &trigger, &criticality, &iesCriticalityDiagnostics)
		simulator_ngap.SendErrorIndication(ran, nil, nil, nil, &criticalityDiagnostics)
		return
	}

	ue := ran.FindUeByRanUeNgapID(rANUENGAPID.Value)
	if ue == nil {
		ngapLog.Warnf("No UE Context[RanUeNgapID:%d]\n", rANUENGAPID.Value)
		return
	}

	responseList := new(ngapType.PDUSessionResourceSetupListSURes)
	failedListSURes := new(ngapType.PDUSessionResourceFailedToSetupListSURes)

	for _, pduSession := range pDUSessionResourceSetupListSUReq.List {
		pduSessionId := pduSession.PDUSessionID.Value
		sess, exist := ue.PduSession[pduSessionId]
		if !exist {
			ngapLog.Warnf("No PduSession Context[PduSessionId:%d]\n", pduSessionId)
			continue
		}
		if pduSession.PDUSessionNASPDU != nil {
			// Handle Nas Msg
			simulator_nas.HandleNAS(ue, pduSession.PDUSessionNASPDU.Value)
		}
		sess.Mtx.Lock()
		ran.AttachSession(sess)
		sess.Mtx.Unlock()
		resTransfer, err := handlePDUSessionResourceSetupRequestTransfer(sess, pduSession.PDUSessionResourceSetupRequestTransfer)
		if err == nil {
			simulator_ngap.AppendPDUSessionResourceSetupListSURes(responseList, pduSessionId, resTransfer)
			// build ULPDR, ULFAR, DLPDR
			simulator_context.Simulator_Self().AttachSession(sess)
		} else {
			ngapLog.Warnf("Pdu Session Resource Setup Fail: %s", err.Error())
			simulator_ngap.AppendPDUSessionResourceFailedToSetupListSURes(failedListSURes, pduSessionId, resTransfer)
		}
	}
	simulator_ngap.SendPDUSessionResourceSetupResponse(ran, ue, responseList, failedListSURes)
	if nASPDU != nil {
		simulator_nas.HandleNAS(ue, nASPDU.Value)
	}
}

func HandlePduSessionResourceReleaseCommand(ran *simulator_context.RanContext, message *ngapType.NGAPPDU) {
	var aMFUENGAPID *ngapType.AMFUENGAPID
	var rANUENGAPID *ngapType.RANUENGAPID
	// var rANPagingPriority *ngapType.RANPagingPriority
	var nASPDU *ngapType.NASPDU
	var pDUSessionResourceToReleaseListRelCmd *ngapType.PDUSessionResourceToReleaseListRelCmd

	var iesCriticalityDiagnostics ngapType.CriticalityDiagnosticsIEList

	if ran == nil {
		ngapLog.Error("RAN Context is nil")
		return
	}

	if message == nil {
		ngapLog.Error("NGAP Message is nil")
		return
	}

	initiatingMessage := message.InitiatingMessage
	if initiatingMessage == nil {
		ngapLog.Error("InitiatingMessage is nil")
		return
	}

	pDUSessionResourceReleaseCommand := initiatingMessage.Value.PDUSessionResourceReleaseCommand
	if pDUSessionResourceReleaseCommand == nil {
		ngapLog.Error("pDUSessionResourceReleaseCommand is nil")
		return
	}

	for _, ie := range pDUSessionResourceReleaseCommand.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDAMFUENGAPID:
			ngapLog.Traceln("[NGAP] Decode IE AMFUENGAPID")
			aMFUENGAPID = ie.Value.AMFUENGAPID
			if aMFUENGAPID == nil {
				ngapLog.Error("AMFUENGAPID is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ie.Id.Value, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		case ngapType.ProtocolIEIDRANUENGAPID:
			ngapLog.Traceln("[NGAP] Decode IE RANUENGAPID")
			rANUENGAPID = ie.Value.RANUENGAPID
			if rANUENGAPID == nil {
				ngapLog.Error("RANUENGAPID is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ie.Id.Value, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		case ngapType.ProtocolIEIDRANPagingPriority:
			ngapLog.Traceln("[NGAP] Decode IE RANPagingPriority")
			// rANPagingPriority = ie.Value.RANPagingPriority
		case ngapType.ProtocolIEIDNASPDU:
			ngapLog.Traceln("[NGAP] Decode IE NASPDU")
			nASPDU = ie.Value.NASPDU
		case ngapType.ProtocolIEIDPDUSessionResourceToReleaseListRelCmd:
			ngapLog.Traceln("[NGAP] Decode IE PDUSessionResourceToReleaseListRelCmd")
			pDUSessionResourceToReleaseListRelCmd = ie.Value.PDUSessionResourceToReleaseListRelCmd
			if pDUSessionResourceToReleaseListRelCmd == nil {
				ngapLog.Error("PDUSessionResourceToReleaseListRelCmd is nil")
				item := buildCriticalityDiagnosticsIEItem(ngapType.CriticalityPresentReject, ie.Id.Value, ngapType.TypeOfErrorPresentMissing)
				iesCriticalityDiagnostics.List = append(iesCriticalityDiagnostics.List, item)
			}
		}
	}

	if len(iesCriticalityDiagnostics.List) > 0 {
		procudureCode := ngapType.ProcedureCodePDUSessionResourceRelease
		trigger := ngapType.TriggeringMessagePresentInitiatingMessage
		criticality := ngapType.CriticalityPresentReject
		criticalityDiagnostics := buildCriticalityDiagnostics(&procudureCode, &trigger, &criticality, &iesCriticalityDiagnostics)
		simulator_ngap.SendErrorIndication(ran, nil, nil, nil, &criticalityDiagnostics)
		return
	}

	ue := ran.FindUeByRanUeNgapID(rANUENGAPID.Value)
	if ue == nil {
		ngapLog.Warnf("No UE Context[RanUeNgapID:%d]\n", rANUENGAPID.Value)
		return
	}

	responseList := ngapType.PDUSessionResourceReleasedListRelRes{}
	for _, pduSession := range pDUSessionResourceToReleaseListRelCmd.List {
		pduSessionId := pduSession.PDUSessionID.Value
		sess, exist := ue.PduSession[pduSessionId]
		if !exist {
			ngapLog.Warnf("No PduSession Context[PduSessionId:%d]\n", pduSessionId)
			continue
		}
		resTransfer, err := handlePDUSessionResourceReleaseCommandTransfer(sess, pduSession.PDUSessionResourceReleaseCommandTransfer)
		if err != nil {
			ngapLog.Warn(err.Error())
		}
		simulator_ngap.AppendPDUSessionResourceReleasedListRelRes(&responseList, pduSessionId, resTransfer)
		ran.DetachSession(sess)
	}

	simulator_ngap.SendPDUSessionResourceReleaseResponse(ran, ue, responseList, nil)

	if nASPDU != nil {
		simulator_nas.HandleNAS(ue, nASPDU.Value)
	}
}

func buildCriticalityDiagnostics(
	procedureCode *int64,
	triggeringMessage *aper.Enumerated,
	procedureCriticality *aper.Enumerated,
	iesCriticalityDiagnostics *ngapType.CriticalityDiagnosticsIEList) (criticalityDiagnostics ngapType.CriticalityDiagnostics) {

	if procedureCode != nil {
		criticalityDiagnostics.ProcedureCode = new(ngapType.ProcedureCode)
		criticalityDiagnostics.ProcedureCode.Value = *procedureCode
	}

	if triggeringMessage != nil {
		criticalityDiagnostics.TriggeringMessage = new(ngapType.TriggeringMessage)
		criticalityDiagnostics.TriggeringMessage.Value = *triggeringMessage
	}

	if procedureCriticality != nil {
		criticalityDiagnostics.ProcedureCriticality = new(ngapType.Criticality)
		criticalityDiagnostics.ProcedureCriticality.Value = *procedureCriticality
	}

	if iesCriticalityDiagnostics != nil {
		criticalityDiagnostics.IEsCriticalityDiagnostics = iesCriticalityDiagnostics
	}

	return criticalityDiagnostics
}

func buildCriticalityDiagnosticsIEItem(ieCriticality aper.Enumerated, ieID int64, typeOfErr aper.Enumerated) (item ngapType.CriticalityDiagnosticsIEItem) {

	item = ngapType.CriticalityDiagnosticsIEItem{
		IECriticality: ngapType.Criticality{
			Value: ieCriticality,
		},
		IEID: ngapType.ProtocolIEID{
			Value: ieID,
		},
		TypeOfError: ngapType.TypeOfError{
			Value: typeOfErr,
		},
	}

	return item
}

func buildCause(present int, value aper.Enumerated) (cause *ngapType.Cause) {
	cause = new(ngapType.Cause)
	cause.Present = present

	switch present {
	case ngapType.CausePresentRadioNetwork:
		cause.RadioNetwork = new(ngapType.CauseRadioNetwork)
		cause.RadioNetwork.Value = value
	case ngapType.CausePresentTransport:
		cause.Transport = new(ngapType.CauseTransport)
		cause.Transport.Value = value
	case ngapType.CausePresentNas:
		cause.Nas = new(ngapType.CauseNas)
		cause.Nas.Value = value
	case ngapType.CausePresentProtocol:
		cause.Protocol = new(ngapType.CauseProtocol)
		cause.Protocol.Value = value
	case ngapType.CausePresentMisc:
		cause.Misc = new(ngapType.CauseMisc)
		cause.Misc.Value = value
	case ngapType.CausePresentNothing:
	}

	return
}

func printAndGetCause(cause *ngapType.Cause) (present int, value aper.Enumerated) {

	present = cause.Present
	switch cause.Present {
	case ngapType.CausePresentRadioNetwork:
		ngapLog.Warnf("Cause RadioNetwork[%d]", cause.RadioNetwork.Value)
		value = cause.RadioNetwork.Value
	case ngapType.CausePresentTransport:
		ngapLog.Warnf("Cause Transport[%d]", cause.Transport.Value)
		value = cause.Transport.Value
	case ngapType.CausePresentProtocol:
		ngapLog.Warnf("Cause Protocol[%d]", cause.Protocol.Value)
		value = cause.Protocol.Value
	case ngapType.CausePresentNas:
		ngapLog.Warnf("Cause Nas[%d]", cause.Nas.Value)
		value = cause.Nas.Value
	case ngapType.CausePresentMisc:
		ngapLog.Warnf("Cause Misc[%d]", cause.Misc.Value)
		value = cause.Misc.Value
	default:
		ngapLog.Errorf("Invalid Cause group[%d]", cause.Present)
	}
	return
}
