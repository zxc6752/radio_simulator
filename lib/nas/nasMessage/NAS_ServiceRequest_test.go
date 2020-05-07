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

type nasMessageServiceRequestData struct {
	inExtendedProtocolDiscriminator uint8
	inSecurityHeader                uint8
	inSpareHalfOctet                uint8
	inServiceRequestMessageIdentity uint8
	inTMSI5GS                       nasType.TMSI5GS
	inUplinkDataStatus              nasType.UplinkDataStatus
	inPDUSessionStatus              nasType.PDUSessionStatus
	inAllowedPDUSessionStatus       nasType.AllowedPDUSessionStatus
	inNASMessageContainer           nasType.NASMessageContainer
}

var nasMessageServiceRequestTable = []nasMessageServiceRequestData{
	{
		inExtendedProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		inSecurityHeader:                0x01,
		inSpareHalfOctet:                0x01,
		inServiceRequestMessageIdentity: nas.MsgTypeServiceRequest,
		inTMSI5GS: nasType.TMSI5GS{
			Len:   7,
			Octet: [7]uint8{0x01, 0x01},
		},
		inUplinkDataStatus: nasType.UplinkDataStatus{
			Iei:    nasMessage.ServiceRequestUplinkDataStatusType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inPDUSessionStatus: nasType.PDUSessionStatus{
			Iei:    nasMessage.ServiceRequestPDUSessionStatusType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inAllowedPDUSessionStatus: nasType.AllowedPDUSessionStatus{
			Iei:    nasMessage.ServiceRequestAllowedPDUSessionStatusType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
		inNASMessageContainer: nasType.NASMessageContainer{
			Iei:    nasMessage.ServiceRequestNASMessageContainerType,
			Len:    2,
			Buffer: []uint8{0x01, 0x01},
		},
	},
}

func TestNasTypeNewServiceRequest(t *testing.T) {
	a := nasMessage.NewServiceRequest(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewServiceRequestMessage(t *testing.T) {

	for i, table := range nasMessageServiceRequestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewServiceRequest(0)
		b := nasMessage.NewServiceRequest(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.ServiceRequestMessageIdentity.SetMessageType(table.inServiceRequestMessageIdentity)

		a.TMSI5GS = table.inTMSI5GS

		a.UplinkDataStatus = nasType.NewUplinkDataStatus(nasMessage.ServiceRequestUplinkDataStatusType)
		a.UplinkDataStatus = &table.inUplinkDataStatus

		a.PDUSessionStatus = nasType.NewPDUSessionStatus(nasMessage.ServiceRequestPDUSessionStatusType)
		a.PDUSessionStatus = &table.inPDUSessionStatus

		a.AllowedPDUSessionStatus = nasType.NewAllowedPDUSessionStatus(nasMessage.ServiceRequestAllowedPDUSessionStatusType)
		a.AllowedPDUSessionStatus = &table.inAllowedPDUSessionStatus

		a.NASMessageContainer = nasType.NewNASMessageContainer(nasMessage.ServiceRequestNASMessageContainerType)
		a.NASMessageContainer = &table.inNASMessageContainer

		buff := new(bytes.Buffer)
		a.EncodeServiceRequest(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodeServiceRequest(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
