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

type RegistrationReject struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.RegistrationRejectMessageIdentity
	nasType.Cause5GMM
	*nasType.T3346Value
	*nasType.T3502Value
	*nasType.EAPMessage
}

func NewRegistrationReject(iei uint8) (registrationReject *RegistrationReject) {
	registrationReject = &RegistrationReject{}
	return registrationReject
}

const (
	RegistrationRejectT3346ValueType uint8 = 0x5F
	RegistrationRejectT3502ValueType uint8 = 0x16
	RegistrationRejectEAPMessageType uint8 = 0x78
)

func (a *RegistrationReject) EncodeRegistrationReject(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.RegistrationRejectMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.Cause5GMM.Octet)
	if a.T3346Value != nil {
		binary.Write(buffer, binary.BigEndian, a.T3346Value.GetIei())
		binary.Write(buffer, binary.BigEndian, a.T3346Value.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.T3346Value.Octet)
	}
	if a.T3502Value != nil {
		binary.Write(buffer, binary.BigEndian, a.T3502Value.GetIei())
		binary.Write(buffer, binary.BigEndian, a.T3502Value.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.T3502Value.Octet)
	}
	if a.EAPMessage != nil {
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei())
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
	}
}

func (a *RegistrationReject) DecodeRegistrationReject(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.RegistrationRejectMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.Cause5GMM.Octet)
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
		case RegistrationRejectT3346ValueType:
			a.T3346Value = nasType.NewT3346Value(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.T3346Value.Len)
			a.T3346Value.SetLen(a.T3346Value.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.T3346Value.Octet)
		case RegistrationRejectT3502ValueType:
			a.T3502Value = nasType.NewT3502Value(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.T3502Value.Len)
			a.T3502Value.SetLen(a.T3502Value.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.T3502Value.Octet)
		case RegistrationRejectEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()])
		default:
		}
	}
}
