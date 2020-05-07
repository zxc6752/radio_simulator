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
	"radio_simulator/lib/nas"
	"radio_simulator/lib/nas/nasMessage"
	"testing"

	"reflect"

	"github.com/stretchr/testify/assert"
)

type nasMessageNotificationData struct {
	inExtendedProtocolDiscriminator uint8
	inSecurityHeader                uint8
	inSpareHalfOctet1               uint8
	inNotificationMessageIdentity   uint8
	inAccessType                    uint8
	inSpareHalfOctet2               uint8
}

var nasMessageNotificationTable = []nasMessageNotificationData{
	{
		inExtendedProtocolDiscriminator: 0x01,
		inSecurityHeader:                0x08,
		inSpareHalfOctet1:               0x01,
		inNotificationMessageIdentity:   nas.MsgTypeNotification,
		inAccessType:                    0x01,
		inSpareHalfOctet2:               0x01,
	},
}

func TestNasTypeNewNotification(t *testing.T) {
	a := nasMessage.NewNotification(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewNotificationMessage(t *testing.T) {

	for i, table := range nasMessageNotificationTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewNotification(0)
		b := nasMessage.NewNotification(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet1)
		a.NotificationMessageIdentity.SetMessageType(table.inNotificationMessageIdentity)
		a.SpareHalfOctetAndAccessType.SetAccessType(table.inAccessType)

		buff := new(bytes.Buffer)
		a.EncodeNotification(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		b.DecodeNotification(&data)
		logger.NasMsgLog.Debugln(data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}

	}
}
