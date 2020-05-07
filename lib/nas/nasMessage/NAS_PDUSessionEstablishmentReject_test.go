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

type nasMessagePDUSessionEstablishmentRejectData struct {
	inExtendedProtocolDiscriminator                uint8
	inPDUSessionID                                 uint8
	inPTI                                          uint8
	inPDUSESSIONESTABLISHMENTREJECTMessageIdentity uint8
	inCause5GSM                                    nasType.Cause5GSM
	inBackoffTimerValue                            nasType.BackoffTimerValue
	inAllowedSSCMode                               nasType.AllowedSSCMode
	inEAPMessage                                   nasType.EAPMessage
	inExtendedProtocolConfigurationOptions         nasType.ExtendedProtocolConfigurationOptions
}

var nasMessagePDUSessionEstablishmentRejectTable = []nasMessagePDUSessionEstablishmentRejectData{
	{
		inExtendedProtocolDiscriminator: nas.MsgTypePDUSessionEstablishmentReject,
		inPDUSessionID:                  0x01,
		inPTI:                           0x01,
		inPDUSESSIONESTABLISHMENTREJECTMessageIdentity: 0x01,
		inCause5GSM: nasType.Cause5GSM{
			Iei:   0,
			Octet: 0x01,
		},
		inBackoffTimerValue: nasType.BackoffTimerValue{
			Iei:   nasMessage.PDUSessionEstablishmentRejectBackoffTimerValueType,
			Len:   2,
			Octet: 0x01,
		},
		inAllowedSSCMode: nasType.AllowedSSCMode{
			Octet: 0xF0,
		},
		inEAPMessage: nasType.EAPMessage{
			Iei:    nasMessage.PDUSessionEstablishmentRejectEAPMessageType,
			Len:    1,
			Buffer: []uint8{0x01},
		},
		inExtendedProtocolConfigurationOptions: nasType.ExtendedProtocolConfigurationOptions{
			Iei:    nasMessage.PDUSessionEstablishmentRejectExtendedProtocolConfigurationOptionsType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewPDUSessionEstablishmentReject(t *testing.T) {
	a := nasMessage.NewPDUSessionEstablishmentReject(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewPDUSessionEstablishmentRejectMessage(t *testing.T) {

	for i, table := range nasMessagePDUSessionEstablishmentRejectTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewPDUSessionEstablishmentReject(0)
		b := nasMessage.NewPDUSessionEstablishmentReject(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.PDUSessionID.SetPDUSessionID(table.inPDUSessionID)
		a.PTI.SetPTI(table.inPTI)
		a.PDUSESSIONESTABLISHMENTREJECTMessageIdentity.SetMessageType(0)
		a.Cause5GSM = table.inCause5GSM

		a.BackoffTimerValue = nasType.NewBackoffTimerValue(nasMessage.PDUSessionEstablishmentRejectBackoffTimerValueType)
		a.BackoffTimerValue = &table.inBackoffTimerValue

		a.AllowedSSCMode = nasType.NewAllowedSSCMode(nasMessage.PDUSessionEstablishmentRejectAllowedSSCModeType)
		a.AllowedSSCMode = &table.inAllowedSSCMode

		a.EAPMessage = nasType.NewEAPMessage(nasMessage.PDUSessionEstablishmentRejectEAPMessageType)
		a.EAPMessage = &table.inEAPMessage

		a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(nasMessage.PDUSessionEstablishmentRejectExtendedProtocolConfigurationOptionsType)
		a.ExtendedProtocolConfigurationOptions = &table.inExtendedProtocolConfigurationOptions

		buff := new(bytes.Buffer)
		a.EncodePDUSessionEstablishmentReject(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodePDUSessionEstablishmentReject(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}

	}
}
