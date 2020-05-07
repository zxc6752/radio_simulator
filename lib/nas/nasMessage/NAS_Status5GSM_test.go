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

type nasMessageStatus5GSMData struct {
	inExtendedProtocolDiscriminator uint8
	inPDUSessionID                  nasType.PDUSessionID
	inPTI                           nasType.PTI
	inStatus5GSMMessageIdentity     uint8
	inCause5GSM                     nasType.Cause5GSM
}

var nasMessageStatus5GSMTable = []nasMessageStatus5GSMData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSSessionManagementMessage,
		inPDUSessionID: nasType.PDUSessionID{
			Octet: 0x01,
		},
		inPTI: nasType.PTI{
			Octet: 0x01,
		},
		inStatus5GSMMessageIdentity: nas.MsgTypeStatus5GSM,
		inCause5GSM: nasType.Cause5GSM{
			Octet: 0x01,
		},
	},
}

func TestNasTypeNewStatus5GSM(t *testing.T) {
	a := nasMessage.NewStatus5GSM(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewStatus5GSMMessage(t *testing.T) {

	for i, table := range nasMessageStatus5GSMTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewStatus5GSM(0)
		b := nasMessage.NewStatus5GSM(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.PDUSessionID = table.inPDUSessionID
		a.PTI = table.inPTI

		a.STATUSMessageIdentity5GSM.SetMessageType(table.inStatus5GSMMessageIdentity)

		a.Cause5GSM = table.inCause5GSM

		buff := new(bytes.Buffer)
		a.EncodeStatus5GSM(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodeStatus5GSM(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
