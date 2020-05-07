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
	"radio_simulator/lib/nas/nasMessage"
	"radio_simulator/lib/nas/nasType"
	"testing"

	"reflect"

	"github.com/stretchr/testify/assert"
)

type nasMessageAuthenticationRejectData struct {
	inExtendedProtocolDiscriminator       uint8
	inSecurityHeader                      uint8
	inSpareHalfOctet                      uint8
	inAuthenticationRejectMessageIdentity uint8
	inEAPMessage                          nasType.EAPMessage
}

var nasMessageAuthenticationRejectTable = []nasMessageAuthenticationRejectData{
	{
		inExtendedProtocolDiscriminator:       0x01,
		inSecurityHeader:                      0x01,
		inSpareHalfOctet:                      0x01,
		inAuthenticationRejectMessageIdentity: 0x01,
		inEAPMessage:                          nasType.EAPMessage{nasMessage.AuthenticationRejectEAPMessageType, 2, []byte{0x00, 0x00}},
	},
}

func TestNasTypeNewAuthenticationReject(t *testing.T) {
	a := nasMessage.NewAuthenticationReject(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewAuthenticationRejectMessage(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: AuthenticationRejectMessage---")
	for i, table := range nasMessageAuthenticationRejectTable {
		t.Logf("Test Cnt:%d", i)

		a := nasMessage.NewAuthenticationReject(nasMessage.AuthenticationRejectEAPMessageType)
		b := nasMessage.NewAuthenticationReject(nasMessage.AuthenticationRejectEAPMessageType)
		assert.NotNil(t, a)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.AuthenticationRejectMessageIdentity.SetMessageType(table.inAuthenticationRejectMessageIdentity)

		a.EAPMessage = nasType.NewEAPMessage(nasMessage.AuthenticationRejectEAPMessageType)
		a.EAPMessage = &table.inEAPMessage

		buff := new(bytes.Buffer)
		a.EncodeAuthenticationReject(buff)
		//fmt.Println("Encode: ", buff)
		logger.NasMsgLog.Debugln(a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		b.DecodeAuthenticationReject(&data)
		logger.NasMsgLog.Debugln("Decode: ", data)
		logger.NasMsgLog.Infoln(b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}

	}
}
