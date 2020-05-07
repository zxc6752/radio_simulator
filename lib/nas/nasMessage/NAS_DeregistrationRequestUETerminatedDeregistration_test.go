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

type nasMessageDeregistrationRequestUETerminatedDeregistrationData struct {
	inExtendedProtocolDiscriminator        uint8
	inSecurityHeaderType                   uint8
	inSpareHalfOctet1                      uint8
	inDeregistrationRequestMessageIdentity uint8
	inSpareHalfOctetAndDeregistrationType  nasType.SpareHalfOctetAndDeregistrationType
	inCause5GMM                            nasType.Cause5GMM
	inT3346Value                           nasType.T3346Value
}

var nasMessageDeregistrationRequestUETerminatedDeregistrationTable = []nasMessageDeregistrationRequestUETerminatedDeregistrationData{
	{
		inExtendedProtocolDiscriminator:        nasMessage.Epd5GSSessionManagementMessage,
		inSecurityHeaderType:                   0x01,
		inSpareHalfOctet1:                      0x01,
		inDeregistrationRequestMessageIdentity: 0x01,
		inSpareHalfOctetAndDeregistrationType: nasType.SpareHalfOctetAndDeregistrationType{
			Octet: 0x0F,
		},
		inCause5GMM: nasType.Cause5GMM{
			Iei:   nasMessage.DeregistrationRequestUETerminatedDeregistrationCause5GMMType,
			Octet: 0x01,
		},
		inT3346Value: nasType.T3346Value{
			Iei:   nasMessage.DeregistrationRequestUETerminatedDeregistrationT3346ValueType,
			Len:   2,
			Octet: 0x01,
		},
	},
}

func TestNasTypeNewDeregistrationRequestUETerminatedDeregistration(t *testing.T) {
	a := nasMessage.NewDeregistrationRequestUETerminatedDeregistration(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewDeregistrationRequestUETerminatedDeregistrationMessage(t *testing.T) {

	for i, table := range nasMessageDeregistrationRequestUETerminatedDeregistrationTable {
		logger.NasMsgLog.Infoln("Test Cnt:", i)
		a := nasMessage.NewDeregistrationRequestUETerminatedDeregistration(0)
		b := nasMessage.NewDeregistrationRequestUETerminatedDeregistration(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeaderType)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet1)
		a.DeregistrationRequestMessageIdentity.SetMessageType(table.inDeregistrationRequestMessageIdentity)

		a.SpareHalfOctetAndDeregistrationType = table.inSpareHalfOctetAndDeregistrationType

		a.Cause5GMM = nasType.NewCause5GMM(nasMessage.DeregistrationRequestUETerminatedDeregistrationCause5GMMType)
		a.Cause5GMM = &table.inCause5GMM

		a.T3346Value = nasType.NewT3346Value(nasMessage.DeregistrationRequestUETerminatedDeregistrationT3346ValueType)
		a.T3346Value = &table.inT3346Value

		buff := new(bytes.Buffer)
		a.EncodeDeregistrationRequestUETerminatedDeregistration(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodeDeregistrationRequestUETerminatedDeregistration(&data)
		logger.NasMsgLog.Debugln("Dncode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}

	}
}
