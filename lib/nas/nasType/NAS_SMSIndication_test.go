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

package nasType_test

import (
	"radio_simulator/lib/nas/nasMessage"
	"radio_simulator/lib/nas/nasType"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNasTypeNewSMSIndication(t *testing.T) {
	a := nasType.NewSMSIndication(nasMessage.ConfigurationUpdateCommandSMSIndicationType)
	assert.NotNil(t, a)
}

var nasTypeSMSIndicationIeiTable = []NasTypeIeiData{
	{0x01, 0x01},
}

func TestNasTypeSMSIndicationGetSetIei(t *testing.T) {
	a := nasType.NewSMSIndication(nasMessage.ConfigurationUpdateCommandSMSIndicationType)
	assert.NotNil(t, a)
	for _, table := range nasTypeSMSIndicationIeiTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

type nasTypeSMSIndicationSAIType struct {
	in  uint8
	out uint8
}

var nasTypeSMSIndicationSAITable = []nasTypeSMSIndicationSAIType{
	{0x01, 0x01},
}

func TestNasTypeSMSIndicationGetSetSAI(t *testing.T) {
	a := nasType.NewSMSIndication(nasMessage.ConfigurationUpdateCommandSMSIndicationType)
	for _, table := range nasTypeSMSIndicationSAITable {
		a.SetSAI(table.in)
		assert.Equal(t, table.out, a.GetSAI())
	}
}

type SMSIndicationTestDataTemplate struct {
	in  nasType.SMSIndication
	out nasType.SMSIndication
}

var SMSIndicationTestData = []nasType.SMSIndication{
	{},
}

var SMSIndicationExpectedTestData = []nasType.SMSIndication{
	{0x11},
}

var SMSIndicationTable = []SMSIndicationTestDataTemplate{
	{SMSIndicationTestData[0], SMSIndicationExpectedTestData[0]},
}

func TestNasTypeSMSIndication(t *testing.T) {

	for _, table := range SMSIndicationTable {

		a := nasType.NewSMSIndication(nasMessage.ConfigurationUpdateCommandSMSIndicationType)
		a.SetIei(0x01)
		a.SetSAI(0x01)

		assert.Equal(t, table.out.Octet, a.Octet)

	}
}
