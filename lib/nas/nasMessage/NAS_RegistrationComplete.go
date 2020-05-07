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

type RegistrationComplete struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.RegistrationCompleteMessageIdentity
	*nasType.SORTransparentContainer
}

func NewRegistrationComplete(iei uint8) (registrationComplete *RegistrationComplete) {
	registrationComplete = &RegistrationComplete{}
	return registrationComplete
}

const (
	RegistrationCompleteSORTransparentContainerType uint8 = 0x73
)

func (a *RegistrationComplete) EncodeRegistrationComplete(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.RegistrationCompleteMessageIdentity.Octet)
	if a.SORTransparentContainer != nil {
		binary.Write(buffer, binary.BigEndian, a.SORTransparentContainer.GetIei())
		binary.Write(buffer, binary.BigEndian, a.SORTransparentContainer.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.SORTransparentContainer.Buffer)
	}
}

func (a *RegistrationComplete) DecodeRegistrationComplete(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.RegistrationCompleteMessageIdentity.Octet)
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
		case RegistrationCompleteSORTransparentContainerType:
			a.SORTransparentContainer = nasType.NewSORTransparentContainer(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.SORTransparentContainer.Len)
			a.SORTransparentContainer.SetLen(a.SORTransparentContainer.GetLen())
			binary.Read(buffer, binary.BigEndian, a.SORTransparentContainer.Buffer[:a.SORTransparentContainer.GetLen()])
		default:
		}
	}
}
