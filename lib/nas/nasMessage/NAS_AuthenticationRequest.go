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

type AuthenticationRequest struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.AuthenticationRequestMessageIdentity
	nasType.SpareHalfOctetAndNgksi
	nasType.ABBA
	*nasType.AuthenticationParameterRAND
	*nasType.AuthenticationParameterAUTN
	*nasType.EAPMessage
}

func NewAuthenticationRequest(iei uint8) (authenticationRequest *AuthenticationRequest) {
	authenticationRequest = &AuthenticationRequest{}
	return authenticationRequest
}

const (
	AuthenticationRequestAuthenticationParameterRANDType uint8 = 0x21
	AuthenticationRequestAuthenticationParameterAUTNType uint8 = 0x20
	AuthenticationRequestEAPMessageType                  uint8 = 0x78
)

func (a *AuthenticationRequest) EncodeAuthenticationRequest(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Write(buffer, binary.BigEndian, &a.AuthenticationRequestMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.SpareHalfOctetAndNgksi.Octet)
	binary.Write(buffer, binary.BigEndian, a.ABBA.GetLen())
	binary.Write(buffer, binary.BigEndian, &a.ABBA.Buffer)
	if a.AuthenticationParameterRAND != nil {
		binary.Write(buffer, binary.BigEndian, a.AuthenticationParameterRAND.GetIei())
		binary.Write(buffer, binary.BigEndian, &a.AuthenticationParameterRAND.Octet)
	}
	if a.AuthenticationParameterAUTN != nil {
		binary.Write(buffer, binary.BigEndian, a.AuthenticationParameterAUTN.GetIei())
		binary.Write(buffer, binary.BigEndian, a.AuthenticationParameterAUTN.GetLen())
		binary.Write(buffer, binary.BigEndian, a.AuthenticationParameterAUTN.Octet[:a.AuthenticationParameterAUTN.GetLen()])
	}
	if a.EAPMessage != nil {
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei())
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
	}
}

func (a *AuthenticationRequest) DecodeAuthenticationRequest(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet)
	binary.Read(buffer, binary.BigEndian, &a.AuthenticationRequestMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndNgksi.Octet)
	binary.Read(buffer, binary.BigEndian, &a.ABBA.Len)
	a.ABBA.SetLen(a.ABBA.GetLen())
	binary.Read(buffer, binary.BigEndian, &a.ABBA.Buffer)
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
		case AuthenticationRequestAuthenticationParameterRANDType:
			a.AuthenticationParameterRAND = nasType.NewAuthenticationParameterRAND(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AuthenticationParameterRAND.Octet)
		case AuthenticationRequestAuthenticationParameterAUTNType:
			a.AuthenticationParameterAUTN = nasType.NewAuthenticationParameterAUTN(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.AuthenticationParameterAUTN.Len)
			a.AuthenticationParameterAUTN.SetLen(a.AuthenticationParameterAUTN.GetLen())
			binary.Read(buffer, binary.BigEndian, a.AuthenticationParameterAUTN.Octet[:a.AuthenticationParameterAUTN.GetLen()])
		case AuthenticationRequestEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()])
		default:
		}
	}
}
