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

type nasMessageNotificationResponseData struct {
	inExtendedProtocolDiscriminator       uint8
	inSecurityHeader                      uint8
	inSpareHalfOctet                      uint8
	inNotificationResponseMessageIdentity uint8
	inPDUSessionStatus                    nasType.PDUSessionStatus
}

var nasMessageNotificationResponseTable = []nasMessageNotificationResponseData{
	{
		inExtendedProtocolDiscriminator:       0x01,
		inSecurityHeader:                      0x08,
		inSpareHalfOctet:                      0x01,
		inNotificationResponseMessageIdentity: 0x01,
		inPDUSessionStatus: nasType.PDUSessionStatus{
			Iei:    nasMessage.NotificationResponsePDUSessionStatusType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewNotificationResponse(t *testing.T) {
	a := nasMessage.NewNotificationResponse(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewNotificationResponseMessage(t *testing.T) {

	for i, table := range nasMessageNotificationResponseTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewNotificationResponse(0)
		b := nasMessage.NewNotificationResponse(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.NotificationResponseMessageIdentity.SetMessageType(table.inNotificationResponseMessageIdentity)

		a.PDUSessionStatus = nasType.NewPDUSessionStatus(nasMessage.NotificationResponsePDUSessionStatusType)
		a.PDUSessionStatus = &table.inPDUSessionStatus

		buff := new(bytes.Buffer)
		a.EncodeNotificationResponse(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		b.DecodeNotificationResponse(&data)
		logger.NasMsgLog.Debugln(data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}

	}
}
