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

package nasMessage_test

import (
	"bytes"
	"radio_simulator/lib/nas/logger"

	//"fmt"
	"radio_simulator/lib/nas/nasMessage"
	"radio_simulator/lib/nas/nasType"
	"testing"

	"reflect"

	"github.com/stretchr/testify/assert"
)

type nasMessageAuthenticationResponseData struct {
	inExtendedProtocolDiscriminator         uint8
	inSecurityHeader                        uint8
	inSpareHalfOctet                        uint8
	inAuthenticationResponseMessageIdentity uint8
	inAuthenticationResponseParameter       nasType.AuthenticationResponseParameter
	inEAPMessage                            nasType.EAPMessage
}

var nasMessageAuthenticationResponseTable = []nasMessageAuthenticationResponseData{
	{
		inExtendedProtocolDiscriminator:         0x01,
		inSecurityHeader:                        0x08,
		inSpareHalfOctet:                        0x01,
		inAuthenticationResponseMessageIdentity: 0x01,
		inAuthenticationResponseParameter:       nasType.AuthenticationResponseParameter{nasMessage.AuthenticationResponseAuthenticationResponseParameterType, 16, [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
		inEAPMessage:                            nasType.EAPMessage{nasMessage.AuthenticationResponseEAPMessageType, 2, []uint8{0x01, 0x01}},
	},
}

func TestNasTypeNewAuthenticationResponse(t *testing.T) {
	a := nasMessage.NewAuthenticationResponse(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewAuthenticationResponseMessage(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: AuthenticationResponseMessage---")
	for i, table := range nasMessageAuthenticationResponseTable {
		logger.NasMsgLog.Infoln("Test Cnt:", i)
		a := nasMessage.NewAuthenticationResponse(0)
		b := nasMessage.NewAuthenticationResponse(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.AuthenticationResponseMessageIdentity.SetMessageType(table.inAuthenticationResponseMessageIdentity)

		a.AuthenticationResponseParameter = nasType.NewAuthenticationResponseParameter(nasMessage.AuthenticationResponseAuthenticationResponseParameterType)
		a.AuthenticationResponseParameter = &table.inAuthenticationResponseParameter

		a.EAPMessage = nasType.NewEAPMessage(nasMessage.AuthenticationResponseEAPMessageType)
		a.EAPMessage = &table.inEAPMessage

		buff := new(bytes.Buffer)
		a.EncodeAuthenticationResponse(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		b.DecodeAuthenticationResponse(&data)
		logger.NasMsgLog.Debugln(data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}

	}
}
