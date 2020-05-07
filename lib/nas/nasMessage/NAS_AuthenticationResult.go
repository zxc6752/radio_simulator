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

type AuthenticationResult struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.AuthenticationResultMessageIdentity
	nasType.SpareHalfOctetAndNgksi
	nasType.EAPMessage
	*nasType.ABBA
}

func NewAuthenticationResult(iei uint8) (authenticationResult *AuthenticationResult) {
	authenticationResult = &AuthenticationResult{}
	return authenticationResult
}

const (
	AuthenticationResultABBAType uint8 = 0x38
)

func (a *AuthenticationResult) EncodeAuthenticationResult(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.AuthenticationResultMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndNgksi.Octet)
	binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen())
	binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
	if a.ABBA != nil {
		binary.Write(buffer, binary.BigEndian, a.ABBA.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ABBA.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ABBA.Buffer)
	}
}

func (a *AuthenticationResult) DecodeAuthenticationResult(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.AuthenticationResultMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndNgksi.Octet)
	binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
	a.EAPMessage.SetLen(a.EAPMessage.GetLen())
	binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
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
		case AuthenticationResultABBAType:
			a.ABBA = nasType.NewABBA(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ABBA.Len)
			a.ABBA.SetLen(a.ABBA.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ABBA.Buffer[:a.ABBA.GetLen()])
		default:
		}
	}
}
