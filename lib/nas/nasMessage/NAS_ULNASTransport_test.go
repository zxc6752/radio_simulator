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

type nasMessageULNASTransportData struct {
	inExtendedProtocolDiscriminator         uint8
	inSecurityHeader                        uint8
	inSpareHalfOctet                        uint8
	inULNASTRANSPORTMessageIdentity         uint8
	inSpareHalfOctetAndPayloadContainerType nasType.SpareHalfOctetAndPayloadContainerType
	inPayloadContainer                      nasType.PayloadContainer
	inPduSessionID2Value                    nasType.PduSessionID2Value
	inOldPDUSessionID                       nasType.OldPDUSessionID
	inRequestType                           nasType.RequestType
	inSNSSAI                                nasType.SNSSAI
	inDNN                                   nasType.DNN
	inAdditionalInformation                 nasType.AdditionalInformation
}

var nasMessageULNASTransportTable = []nasMessageULNASTransportData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		inSecurityHeader:                0x01,
		inSpareHalfOctet:                0x01,
		inULNASTRANSPORTMessageIdentity: nas.MsgTypeULNASTransport,
		inSpareHalfOctetAndPayloadContainerType: nasType.SpareHalfOctetAndPayloadContainerType{
			Octet: 0x01,
		},
		inPayloadContainer: nasType.PayloadContainer{
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inPduSessionID2Value: nasType.PduSessionID2Value{
			Iei:   nasMessage.ULNASTransportPduSessionID2ValueType,
			Octet: 0x01,
		},
		inOldPDUSessionID: nasType.OldPDUSessionID{
			Iei:   nasMessage.ULNASTransportOldPDUSessionIDType,
			Octet: 0x01,
		},
		inRequestType: nasType.RequestType{
			Octet: 0x80,
		},
		inDNN: nasType.DNN{
			Iei:    nasMessage.ULNASTransportDNNType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inSNSSAI: nasType.SNSSAI{
			Iei:   nasMessage.ULNASTransportSNSSAIType,
			Len:   8,
			Octet: [8]uint8{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01},
		},
		inAdditionalInformation: nasType.AdditionalInformation{
			Iei:    nasMessage.ULNASTransportAdditionalInformationType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewULNASTransport(t *testing.T) {
	a := nasMessage.NewULNASTransport(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewULNASTransportMessage(t *testing.T) {

	for i, table := range nasMessageULNASTransportTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewULNASTransport(0)
		b := nasMessage.NewULNASTransport(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.ULNASTRANSPORTMessageIdentity.SetMessageType(table.inULNASTRANSPORTMessageIdentity)

		a.SpareHalfOctetAndPayloadContainerType = table.inSpareHalfOctetAndPayloadContainerType

		a.PayloadContainer = table.inPayloadContainer

		a.PduSessionID2Value = nasType.NewPduSessionID2Value(nasMessage.ULNASTransportPduSessionID2ValueType)
		a.PduSessionID2Value = &table.inPduSessionID2Value

		a.OldPDUSessionID = nasType.NewOldPDUSessionID(nasMessage.ULNASTransportOldPDUSessionIDType)
		a.OldPDUSessionID = &table.inOldPDUSessionID

		a.RequestType = nasType.NewRequestType(nasMessage.ULNASTransportRequestTypeType)
		a.RequestType = &table.inRequestType

		a.SNSSAI = nasType.NewSNSSAI(nasMessage.ULNASTransportSNSSAIType)
		a.SNSSAI = &table.inSNSSAI

		a.DNN = nasType.NewDNN(nasMessage.ULNASTransportDNNType)
		a.DNN = &table.inDNN

		a.AdditionalInformation = nasType.NewAdditionalInformation(nasMessage.ULNASTransportAdditionalInformationType)
		a.AdditionalInformation = &table.inAdditionalInformation

		buff := new(bytes.Buffer)
		a.EncodeULNASTransport(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodeULNASTransport(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
