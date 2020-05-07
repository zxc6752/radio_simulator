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

type nasMessagePDUSessionModificationCommandData struct {
	inExtendedProtocolDiscriminator                uint8
	inPDUSessionID                                 uint8
	inPTI                                          uint8
	inPDUSESSIONMODIFICATIONCOMMANDMessageIdentity uint8
	inCause5GSM                                    nasType.Cause5GSM
	inSessionAMBR                                  nasType.SessionAMBR
	inRQTimerValue                                 nasType.RQTimerValue
	inAlwaysonPDUSessionIndication                 nasType.AlwaysonPDUSessionIndication
	inAuthorizedQosRules                           nasType.AuthorizedQosRules
	inMappedEPSBearerContexts                      nasType.MappedEPSBearerContexts
	inAuthorizedQosFlowDescriptions                nasType.AuthorizedQosFlowDescriptions
	inExtendedProtocolConfigurationOptions         nasType.ExtendedProtocolConfigurationOptions
}

var nasMessagePDUSessionModificationCommandTable = []nasMessagePDUSessionModificationCommandData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSSessionManagementMessage,
		inPDUSessionID:                  0x01,
		inPTI:                           0x01,
		inPDUSESSIONMODIFICATIONCOMMANDMessageIdentity: 0x01,
		inCause5GSM: nasType.Cause5GSM{
			Iei:   nasMessage.PDUSessionModificationCommandCause5GSMType,
			Octet: 0x01,
		},
		inSessionAMBR: nasType.SessionAMBR{
			Iei:   nasMessage.PDUSessionModificationCommandSessionAMBRType,
			Len:   6,
			Octet: [6]uint8{0x01, 0x02, 0x03, 0x04, 0x05, 0x06},
		},
		inRQTimerValue: nasType.RQTimerValue{
			Iei:   nasMessage.PDUSessionModificationCommandRQTimerValueType,
			Octet: 0x01,
		},
		inAlwaysonPDUSessionIndication: nasType.AlwaysonPDUSessionIndication{
			Octet: 0x80,
		},
		inAuthorizedQosRules: nasType.AuthorizedQosRules{
			Iei:    nasMessage.PDUSessionModificationCommandAuthorizedQosRulesType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inMappedEPSBearerContexts: nasType.MappedEPSBearerContexts{
			Iei:    nasMessage.PDUSessionModificationCommandMappedEPSBearerContextsType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inAuthorizedQosFlowDescriptions: nasType.AuthorizedQosFlowDescriptions{
			Iei:    nasMessage.PDUSessionModificationCommandAuthorizedQosFlowDescriptionsType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inExtendedProtocolConfigurationOptions: nasType.ExtendedProtocolConfigurationOptions{
			Iei:    nasMessage.PDUSessionModificationCommandExtendedProtocolConfigurationOptionsType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewPDUSessionModificationCommand(t *testing.T) {
	a := nasMessage.NewPDUSessionModificationCommand(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewPDUSessionModificationCommandMessage(t *testing.T) {

	for i, table := range nasMessagePDUSessionModificationCommandTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewPDUSessionModificationCommand(0)
		b := nasMessage.NewPDUSessionModificationCommand(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.PDUSessionID.SetPDUSessionID(table.inPDUSessionID)
		a.PTI.SetPTI(table.inPTI)
		a.PDUSESSIONMODIFICATIONCOMMANDMessageIdentity.SetMessageType(table.inPDUSESSIONMODIFICATIONCOMMANDMessageIdentity)

		a.Cause5GSM = nasType.NewCause5GSM(nasMessage.PDUSessionModificationCommandCause5GSMType)
		a.Cause5GSM = &table.inCause5GSM

		a.SessionAMBR = nasType.NewSessionAMBR(nasMessage.PDUSessionModificationCommandSessionAMBRType)
		a.SessionAMBR = &table.inSessionAMBR

		a.RQTimerValue = nasType.NewRQTimerValue(nasMessage.PDUSessionModificationCommandRQTimerValueType)
		a.RQTimerValue = &table.inRQTimerValue

		a.AlwaysonPDUSessionIndication = nasType.NewAlwaysonPDUSessionIndication(nasMessage.PDUSessionModificationCommandAlwaysonPDUSessionIndicationType)
		a.AlwaysonPDUSessionIndication = &table.inAlwaysonPDUSessionIndication

		a.AuthorizedQosRules = nasType.NewAuthorizedQosRules(nasMessage.PDUSessionModificationCommandAuthorizedQosRulesType)
		a.AuthorizedQosRules = &table.inAuthorizedQosRules

		a.MappedEPSBearerContexts = nasType.NewMappedEPSBearerContexts(nasMessage.PDUSessionModificationCommandMappedEPSBearerContextsType)
		a.MappedEPSBearerContexts = &table.inMappedEPSBearerContexts

		a.AuthorizedQosFlowDescriptions = nasType.NewAuthorizedQosFlowDescriptions(nasMessage.PDUSessionModificationCommandAuthorizedQosFlowDescriptionsType)
		a.AuthorizedQosFlowDescriptions = &table.inAuthorizedQosFlowDescriptions

		a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(nasMessage.PDUSessionModificationCommandExtendedProtocolConfigurationOptionsType)
		a.ExtendedProtocolConfigurationOptions = &table.inExtendedProtocolConfigurationOptions

		buff := new(bytes.Buffer)
		a.EncodePDUSessionModificationCommand(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodePDUSessionModificationCommand(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}

	}
}
