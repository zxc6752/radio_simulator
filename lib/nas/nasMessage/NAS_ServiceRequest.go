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

type ServiceRequest struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.ServiceRequestMessageIdentity
	nasType.ServiceTypeAndNgksi
	nasType.TMSI5GS
	*nasType.UplinkDataStatus
	*nasType.PDUSessionStatus
	*nasType.AllowedPDUSessionStatus
	*nasType.NASMessageContainer
}

func NewServiceRequest(iei uint8) (serviceRequest *ServiceRequest) {
	serviceRequest = &ServiceRequest{}
	return serviceRequest
}

const (
	ServiceRequestUplinkDataStatusType        uint8 = 0x40
	ServiceRequestPDUSessionStatusType        uint8 = 0x50
	ServiceRequestAllowedPDUSessionStatusType uint8 = 0x25
	ServiceRequestNASMessageContainerType     uint8 = 0x71
)

func (a *ServiceRequest) EncodeServiceRequest(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.ServiceRequestMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.ServiceTypeAndNgksi.Octet)
	binary.Write(buffer, binary.BigEndian, a.TMSI5GS.GetLen())
	binary.Write(buffer, binary.BigEndian, &a.TMSI5GS.Octet)
	if a.UplinkDataStatus != nil {
		binary.Write(buffer, binary.BigEndian, a.UplinkDataStatus.GetIei())
		binary.Write(buffer, binary.BigEndian, a.UplinkDataStatus.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.UplinkDataStatus.Buffer)
	}
	if a.PDUSessionStatus != nil {
		binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetIei())
		binary.Write(buffer, binary.BigEndian, a.PDUSessionStatus.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.PDUSessionStatus.Buffer)
	}
	if a.AllowedPDUSessionStatus != nil {
		binary.Write(buffer, binary.BigEndian, a.AllowedPDUSessionStatus.GetIei())
		binary.Write(buffer, binary.BigEndian, a.AllowedPDUSessionStatus.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.AllowedPDUSessionStatus.Buffer)
	}
	if a.NASMessageContainer != nil {
		binary.Write(buffer, binary.BigEndian, a.NASMessageContainer.GetIei())
		binary.Write(buffer, binary.BigEndian, a.NASMessageContainer.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.NASMessageContainer.Buffer)
	}
}

func (a *ServiceRequest) DecodeServiceRequest(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.ServiceRequestMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.ServiceTypeAndNgksi.Octet)
	binary.Read(buffer, binary.BigEndian, &a.TMSI5GS.Len)
	a.TMSI5GS.SetLen(a.TMSI5GS.GetLen())
	binary.Read(buffer, binary.BigEndian, &a.TMSI5GS.Octet)
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
		case ServiceRequestUplinkDataStatusType:
			a.UplinkDataStatus = nasType.NewUplinkDataStatus(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.UplinkDataStatus.Len)
			a.UplinkDataStatus.SetLen(a.UplinkDataStatus.GetLen())
			binary.Read(buffer, binary.BigEndian, a.UplinkDataStatus.Buffer[:a.UplinkDataStatus.GetLen()])
		case ServiceRequestPDUSessionStatusType:
			a.PDUSessionStatus = nasType.NewPDUSessionStatus(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.PDUSessionStatus.Len)
			a.PDUSessionStatus.SetLen(a.PDUSessionStatus.GetLen())
			binary.Read(buffer, binary.BigEndian, a.PDUSessionStatus.Buffer[:a.PDUSessionStatus.GetLen()])
		case ServiceRequestAllowedPDUSessionStatusType:
			a.AllowedPDUSessionStatus = nasType.NewAllowedPDUSessionStatus(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AllowedPDUSessionStatus.Len)
			a.AllowedPDUSessionStatus.SetLen(a.AllowedPDUSessionStatus.GetLen())
			binary.Read(buffer, binary.BigEndian, a.AllowedPDUSessionStatus.Buffer[:a.AllowedPDUSessionStatus.GetLen()])
		case ServiceRequestNASMessageContainerType:
			a.NASMessageContainer = nasType.NewNASMessageContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.NASMessageContainer.Len)
			a.NASMessageContainer.SetLen(a.NASMessageContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.NASMessageContainer.Buffer[:a.NASMessageContainer.GetLen()])
		default:
		}
	}
}
