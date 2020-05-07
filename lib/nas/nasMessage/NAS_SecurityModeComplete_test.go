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

type nasMessageSecurityModeCompleteData struct {
	inExtendedProtocolDiscriminator       uint8
	inSecurityHeader                      uint8
	inSpareHalfOctet                      uint8
	inSecurityModeCompleteMessageIdentity uint8
	inIMEISV                              nasType.IMEISV
	inNASMessageContainer                 nasType.NASMessageContainer
}

var nasMessageSecurityModeCompleteTable = []nasMessageSecurityModeCompleteData{
	{
		inExtendedProtocolDiscriminator:       nasMessage.Epd5GSMobilityManagementMessage,
		inSecurityHeader:                      0x01,
		inSpareHalfOctet:                      0x01,
		inSecurityModeCompleteMessageIdentity: nas.MsgTypeSecurityModeComplete,
		inIMEISV: nasType.IMEISV{
			Iei:   nasMessage.SecurityModeCompleteIMEISVType,
			Len:   9,
			Octet: [9]uint8{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01},
		},
		inNASMessageContainer: nasType.NASMessageContainer{
			Iei:    nasMessage.SecurityModeCompleteNASMessageContainerType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewSecurityModeComplete(t *testing.T) {
	a := nasMessage.NewSecurityModeComplete(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewSecurityModeCompleteMessage(t *testing.T) {

	for i, table := range nasMessageSecurityModeCompleteTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewSecurityModeComplete(0)
		b := nasMessage.NewSecurityModeComplete(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.SecurityModeCompleteMessageIdentity.SetMessageType(table.inSecurityModeCompleteMessageIdentity)

		a.IMEISV = nasType.NewIMEISV(nasMessage.SecurityModeCompleteIMEISVType)
		a.IMEISV = &table.inIMEISV

		a.NASMessageContainer = nasType.NewNASMessageContainer(nasMessage.SecurityModeCompleteNASMessageContainerType)
		a.NASMessageContainer = &table.inNASMessageContainer

		buff := new(bytes.Buffer)
		a.EncodeSecurityModeComplete(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodeSecurityModeComplete(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
