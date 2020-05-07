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

type SecurityModeComplete struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.SecurityModeCompleteMessageIdentity
	*nasType.IMEISV
	*nasType.NASMessageContainer
}

func NewSecurityModeComplete(iei uint8) (securityModeComplete *SecurityModeComplete) {
	securityModeComplete = &SecurityModeComplete{}
	return securityModeComplete
}

const (
	SecurityModeCompleteIMEISVType              uint8 = 0x77
	SecurityModeCompleteNASMessageContainerType uint8 = 0x71
)

func (a *SecurityModeComplete) EncodeSecurityModeComplete(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SecurityModeCompleteMessageIdentity.Octet)
	if a.IMEISV != nil {
		binary.Write(buffer, binary.BigEndian, a.IMEISV.GetIei())
		binary.Write(buffer, binary.BigEndian, a.IMEISV.GetLen())
		binary.Write(buffer, binary.BigEndian, a.IMEISV.Octet[:a.IMEISV.GetLen()])
	}
	if a.NASMessageContainer != nil {
		binary.Write(buffer, binary.BigEndian, a.NASMessageContainer.GetIei())
		binary.Write(buffer, binary.BigEndian, a.NASMessageContainer.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.NASMessageContainer.Buffer)
	}
}

func (a *SecurityModeComplete) DecodeSecurityModeComplete(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SecurityModeCompleteMessageIdentity.Octet)
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
		case SecurityModeCompleteIMEISVType:
			a.IMEISV = nasType.NewIMEISV(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.IMEISV.Len)
			a.IMEISV.SetLen(a.IMEISV.GetLen())
			binary.Read(buffer, binary.BigEndian, a.IMEISV.Octet[:a.IMEISV.GetLen()])
		case SecurityModeCompleteNASMessageContainerType:
			a.NASMessageContainer = nasType.NewNASMessageContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.NASMessageContainer.Len)
			a.NASMessageContainer.SetLen(a.NASMessageContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.NASMessageContainer.Buffer[:a.NASMessageContainer.GetLen()])
		default:
		}
	}
}
