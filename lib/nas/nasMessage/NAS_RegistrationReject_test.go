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
	"radio_simulator/lib/nas"
	"radio_simulator/lib/nas/logger"
	"radio_simulator/lib/nas/nasMessage"
	"radio_simulator/lib/nas/nasType"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type nasMessageRegistrationRejectData struct {
	inExtendedProtocolDiscriminator     uint8
	inSecurityHeader                    uint8
	inSpareHalfOctet                    uint8
	inRegistrationRejectMessageIdentity uint8
	inCause5GMM                         nasType.Cause5GMM
	inT3346Value                        nasType.T3346Value
	inT3502Value                        nasType.T3502Value
	inEAPMessage                        nasType.EAPMessage
}

var nasMessageRegistrationRejectTable = []nasMessageRegistrationRejectData{
	{
		inExtendedProtocolDiscriminator:     nasMessage.Epd5GSMobilityManagementMessage,
		inSecurityHeader:                    0x01,
		inSpareHalfOctet:                    0x01,
		inRegistrationRejectMessageIdentity: nas.MsgTypeRegistrationReject,
		inCause5GMM: nasType.Cause5GMM{
			Octet: 0x01,
		},
		inT3346Value: nasType.T3346Value{
			Iei:   nasMessage.RegistrationRejectT3346ValueType,
			Len:   2,
			Octet: 0x01,
		},
		inT3502Value: nasType.T3502Value{
			Iei:   nasMessage.RegistrationRejectT3502ValueType,
			Len:   2,
			Octet: 0x01,
		},
		inEAPMessage: nasType.EAPMessage{
			Iei:    nasMessage.RegistrationRejectEAPMessageType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewRegistrationReject(t *testing.T) {
	a := nasMessage.NewRegistrationReject(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewRegistrationRejectMessage(t *testing.T) {

	for i, table := range nasMessageRegistrationRejectTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewRegistrationReject(0)
		b := nasMessage.NewRegistrationReject(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.RegistrationRejectMessageIdentity.SetMessageType(table.inRegistrationRejectMessageIdentity)

		a.Cause5GMM = table.inCause5GMM

		a.T3346Value = nasType.NewT3346Value(nasMessage.RegistrationRejectT3346ValueType)
		a.T3346Value = &table.inT3346Value

		a.T3502Value = nasType.NewT3502Value(nasMessage.RegistrationRejectT3502ValueType)
		a.T3502Value = &table.inT3502Value

		a.EAPMessage = nasType.NewEAPMessage(nasMessage.RegistrationRejectEAPMessageType)
		a.EAPMessage = &table.inEAPMessage

		buff := new(bytes.Buffer)
		a.EncodeRegistrationReject(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodeRegistrationReject(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
