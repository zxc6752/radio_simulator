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

type nasMessagePDUSessionReleaseCompleteData struct {
	inExtendedProtocolDiscriminator            uint8
	inPDUSessionID                             uint8
	inPTI                                      uint8
	inPDUSESSIONRELEASECOMPLETEMessageIdentity uint8
	inCause5GSM                                nasType.Cause5GSM
	inExtendedProtocolConfigurationOptions     nasType.ExtendedProtocolConfigurationOptions
}

var nasMessagePDUSessionReleaseCompleteTable = []nasMessagePDUSessionReleaseCompleteData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSSessionManagementMessage,
		inPDUSessionID:                  0x01,
		inPTI:                           0x01,
		inPDUSESSIONRELEASECOMPLETEMessageIdentity: 0x01,
		inCause5GSM: nasType.Cause5GSM{
			Iei:   nasMessage.PDUSessionReleaseCompleteCause5GSMType,
			Octet: 0x01,
		},
		inExtendedProtocolConfigurationOptions: nasType.ExtendedProtocolConfigurationOptions{
			Iei:    nasMessage.PDUSessionReleaseCompleteExtendedProtocolConfigurationOptionsType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewPDUSessionReleaseComplete(t *testing.T) {
	a := nasMessage.NewPDUSessionReleaseComplete(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewPDUSessionReleaseCompleteMessage(t *testing.T) {

	for i, table := range nasMessagePDUSessionReleaseCompleteTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewPDUSessionReleaseComplete(0)
		b := nasMessage.NewPDUSessionReleaseComplete(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.PDUSessionID.SetPDUSessionID(table.inPDUSessionID)
		a.PTI.SetPTI(table.inPTI)
		a.PDUSESSIONRELEASECOMPLETEMessageIdentity.SetMessageType(table.inPDUSESSIONRELEASECOMPLETEMessageIdentity)

		a.Cause5GSM = nasType.NewCause5GSM(nasMessage.PDUSessionReleaseCompleteCause5GSMType)
		a.Cause5GSM = &table.inCause5GSM

		a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(nasMessage.PDUSessionReleaseCompleteExtendedProtocolConfigurationOptionsType)
		a.ExtendedProtocolConfigurationOptions = &table.inExtendedProtocolConfigurationOptions

		buff := new(bytes.Buffer)
		a.EncodePDUSessionReleaseComplete(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodePDUSessionReleaseComplete(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)
		//fmt.Println(a.Cause5GSM)
		//fmt.Println(b.Cause5GSM)
		//fmt.Println(a.ExtendedProtocolConfigurationOptions)
		//fmt.Println(b.ExtendedProtocolConfigurationOptions)
		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}

	}
}
