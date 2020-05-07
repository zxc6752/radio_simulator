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
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type nasMessageSecurityProtected5GSNASMessageData struct {
	inExtendedProtocolDiscriminator uint8
	inSecurityHeader                uint8
	inSpareHalfOctet                uint8
	inMessageAuthenticationCode     nasType.MessageAuthenticationCode
	inSequenceNumber                nasType.SequenceNumber
	inPlain5GSNASMessage            nasType.Plain5GSNASMessage
}

var nasMessageSecurityProtected5GSNASMessageTable = []nasMessageSecurityProtected5GSNASMessageData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		inSecurityHeader:                0x01,
		inSpareHalfOctet:                0x01,
		inMessageAuthenticationCode: nasType.MessageAuthenticationCode{
			Octet: [4]uint8{0x01, 0x01, 0x01, 0x01},
		},
		inSequenceNumber: nasType.SequenceNumber{
			Octet: 0x01,
		},
		inPlain5GSNASMessage: nasType.Plain5GSNASMessage{},
	},
}

func TestNasTypeNewSecurityProtected5GSNASMessage(t *testing.T) {
	a := nasMessage.NewSecurityProtected5GSNASMessage(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewSecurityProtected5GSNASMessageMessage(t *testing.T) {

	for i, table := range nasMessageSecurityProtected5GSNASMessageTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewSecurityProtected5GSNASMessage(0)
		b := nasMessage.NewSecurityProtected5GSNASMessage(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet)

		a.MessageAuthenticationCode = table.inMessageAuthenticationCode
		a.SequenceNumber = table.inSequenceNumber
		a.Plain5GSNASMessage = table.inPlain5GSNASMessage

		buff := new(bytes.Buffer)
		a.EncodeSecurityProtected5GSNASMessage(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodeSecurityProtected5GSNASMessage(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
