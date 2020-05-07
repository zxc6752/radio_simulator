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

type nasMessagePDUSessionModificationCommandRejectData struct {
	inExtendedProtocolDiscriminator                      uint8
	inPDUSessionID                                       uint8
	inPTI                                                uint8
	inPDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity uint8
	inCause5GSM                                          nasType.Cause5GSM
	inExtendedProtocolConfigurationOptions               nasType.ExtendedProtocolConfigurationOptions
}

var nasMessagePDUSessionModificationCommandRejectTable = []nasMessagePDUSessionModificationCommandRejectData{
	{
		inExtendedProtocolDiscriminator: nas.MsgTypePDUSessionModificationCommandReject,
		inPDUSessionID:                  0x01,
		inPTI:                           0x01,
		inPDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity: 0x01,
		inCause5GSM: nasType.Cause5GSM{
			Iei:   0,
			Octet: 0x01,
		},
		inExtendedProtocolConfigurationOptions: nasType.ExtendedProtocolConfigurationOptions{
			Iei:    nasMessage.PDUSessionModificationCommandRejectExtendedProtocolConfigurationOptionsType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewPDUSessionModificationCommandReject(t *testing.T) {
	a := nasMessage.NewPDUSessionModificationCommandReject(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewPDUSessionModificationCommandRejectMessage(t *testing.T) {

	for i, table := range nasMessagePDUSessionModificationCommandRejectTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewPDUSessionModificationCommandReject(0)
		b := nasMessage.NewPDUSessionModificationCommandReject(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.PDUSessionID.SetPDUSessionID(table.inPDUSessionID)
		a.PTI.SetPTI(table.inPTI)
		a.PDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity.SetMessageType(table.inPDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity)
		a.Cause5GSM = table.inCause5GSM

		a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(nasMessage.PDUSessionModificationCommandRejectExtendedProtocolConfigurationOptionsType)
		a.ExtendedProtocolConfigurationOptions = &table.inExtendedProtocolConfigurationOptions

		buff := new(bytes.Buffer)
		a.EncodePDUSessionModificationCommandReject(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodePDUSessionModificationCommandReject(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}

	}
}
