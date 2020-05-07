/* 
 * Copyright 2019 Communication Service/Software Laboratory, National Chiao Tung University
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Modifications Copyright 2020 Weiting Hu <zxc6752.cs03@g2.nctu.edu.tw>
 */

package ngapConvert

import (
	"encoding/hex"
	"radio_simulator/lib/aper"
	"radio_simulator/lib/ngap/logger"
	"radio_simulator/lib/ngap/ngapType"
	"radio_simulator/lib/openapi/models"
	"strings"
)

func TraceDataToModels(traceActivation ngapType.TraceActivation) (traceData models.TraceData) {
	// TODO: finish this function when need
	return
}

func TraceDataToNgap(traceData models.TraceData, trsr string) (traceActivation ngapType.TraceActivation) {

	if len(trsr) != 4 {
		logger.NgapLog.Warningln("Trace Recording Session Reference should be 2 octets")
		return
	}

	//NG-RAN Trace ID (left most 6 octet Trace Reference + last 2 octet Trace Recoding Session Reference)
	subStringSlice := strings.Split(traceData.TraceRef, "-")

	if len(subStringSlice) != 2 {
		logger.NgapLog.Warningln("TraceRef format is not correct")
		return
	}

	plmnID := models.PlmnId{}
	plmnID.Mcc = subStringSlice[0][:3]
	plmnID.Mnc = subStringSlice[0][3:]
	traceID, _ := hex.DecodeString(subStringSlice[1])

	tmp := PlmnIdToNgap(plmnID)
	traceReference := append(tmp.Value, traceID...)
	trsrNgap, _ := hex.DecodeString(trsr)

	nGRANTraceID := append(traceReference, trsrNgap...)

	traceActivation.NGRANTraceID.Value = nGRANTraceID

	// Interfaces To Trace
	interfacesToTrace, _ := hex.DecodeString(traceData.InterfaceList)
	traceActivation.InterfacesToTrace.Value = aper.BitString{
		Bytes:     interfacesToTrace,
		BitLength: 8,
	}

	// Trace Collection Entity IP Address
	ngapIP := IPAddressToNgap(traceData.CollectionEntityIpv4Addr, traceData.CollectionEntityIpv6Addr)
	traceActivation.TraceCollectionEntityIPAddress = ngapIP

	// Trace Depth
	switch traceData.TraceDepth {
	case models.TraceDepth_MINIMUM:
		traceActivation.TraceDepth.Value = ngapType.TraceDepthPresentMinimum
	case models.TraceDepth_MEDIUM:
		traceActivation.TraceDepth.Value = ngapType.TraceDepthPresentMedium
	case models.TraceDepth_MAXIMUM:
		traceActivation.TraceDepth.Value = ngapType.TraceDepthPresentMaximum
	case models.TraceDepth_MINIMUM_WO_VENDOR_EXTENSION:
		traceActivation.TraceDepth.Value = ngapType.TraceDepthPresentMinimumWithoutVendorSpecificExtension
	case models.TraceDepth_MEDIUM_WO_VENDOR_EXTENSION:
		traceActivation.TraceDepth.Value = ngapType.TraceDepthPresentMediumWithoutVendorSpecificExtension
	case models.TraceDepth_MAXIMUM_WO_VENDOR_EXTENSION:
		traceActivation.TraceDepth.Value = ngapType.TraceDepthPresentMaximumWithoutVendorSpecificExtension
	}

	return
}
