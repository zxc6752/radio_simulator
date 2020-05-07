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

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"radio_simulator/lib/nas/nasType"
)

type DLNASTransport struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.DLNASTRANSPORTMessageIdentity
	nasType.SpareHalfOctetAndPayloadContainerType
	nasType.PayloadContainer
	*nasType.PduSessionID2Value
	*nasType.AdditionalInformation
	*nasType.Cause5GMM
	*nasType.BackoffTimerValue
}

func NewDLNASTransport(iei uint8) (dLNASTransport *DLNASTransport) {
	dLNASTransport = &DLNASTransport{}
	return dLNASTransport
}

const (
	DLNASTransportPduSessionID2ValueType    uint8 = 0x12
	DLNASTransportAdditionalInformationType uint8 = 0x24
	DLNASTransportCause5GMMType             uint8 = 0x58
	DLNASTransportBackoffTimerValueType     uint8 = 0x37
)

func (a *DLNASTransport) EncodeDLNASTransport(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.DLNASTRANSPORTMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndPayloadContainerType.Octet)
	binary.Write(buffer, binary.BigEndian, a.PayloadContainer.GetLen())
	binary.Write(buffer, binary.BigEndian, &a.PayloadContainer.Buffer)
	if a.PduSessionID2Value != nil {
		binary.Write(buffer, binary.BigEndian, a.PduSessionID2Value.GetIei())
		binary.Write(buffer, binary.BigEndian, &a.PduSessionID2Value.Octet)
	}
	if a.AdditionalInformation != nil {
		binary.Write(buffer, binary.BigEndian, a.AdditionalInformation.GetIei())
		binary.Write(buffer, binary.BigEndian, a.AdditionalInformation.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.AdditionalInformation.Buffer)
	}
	if a.Cause5GMM != nil {
		binary.Write(buffer, binary.BigEndian, a.Cause5GMM.GetIei())
		binary.Write(buffer, binary.BigEndian, &a.Cause5GMM.Octet)
	}
	if a.BackoffTimerValue != nil {
		binary.Write(buffer, binary.BigEndian, a.BackoffTimerValue.GetIei())
		binary.Write(buffer, binary.BigEndian, a.BackoffTimerValue.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.BackoffTimerValue.Octet)
	}
}

func (a *DLNASTransport) DecodeDLNASTransport(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.DLNASTRANSPORTMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndPayloadContainerType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PayloadContainer.Len)
	a.PayloadContainer.SetLen(a.PayloadContainer.GetLen())
	binary.Read(buffer, binary.BigEndian, &a.PayloadContainer.Buffer)
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		binary.Read(buffer, binary.BigEndian, &ieiN)
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case DLNASTransportPduSessionID2ValueType:
			a.PduSessionID2Value = nasType.NewPduSessionID2Value(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.PduSessionID2Value.Octet)
		case DLNASTransportAdditionalInformationType:
			a.AdditionalInformation = nasType.NewAdditionalInformation(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AdditionalInformation.Len)
			a.AdditionalInformation.SetLen(a.AdditionalInformation.GetLen())
			binary.Read(buffer, binary.BigEndian, a.AdditionalInformation.Buffer[:a.AdditionalInformation.GetLen()])
		case DLNASTransportCause5GMMType:
			a.Cause5GMM = nasType.NewCause5GMM(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.Cause5GMM.Octet)
		case DLNASTransportBackoffTimerValueType:
			a.BackoffTimerValue = nasType.NewBackoffTimerValue(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Len)
			a.BackoffTimerValue.SetLen(a.BackoffTimerValue.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Octet)
		default:
		}
	}
}
