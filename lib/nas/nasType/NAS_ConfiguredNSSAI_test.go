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

func TestNasTypeNewConfiguredNSSAI(t *testing.T) {
	a := nasType.NewConfiguredNSSAI(nasMessage.ConfigurationUpdateCommandAllowedNSSAIType)
	assert.NotNil(t, a)

}

var nasTypeAuthenticationRequestConfiguredNSSAIIeiTable = []NasTypeIeiData{
	{nasMessage.ConfigurationUpdateCommandAllowedNSSAIType, nasMessage.ConfigurationUpdateCommandAllowedNSSAIType},
}

func TestNasTypeConfiguredNSSAIGetSetIei(t *testing.T) {
	a := nasType.NewConfiguredNSSAI(nasMessage.ConfigurationUpdateCommandAllowedNSSAIType)
	for _, table := range nasTypeAuthenticationRequestConfiguredNSSAIIeiTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

var nasTypeConfiguredNSSAILenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeConfiguredNSSAIGetSetLen(t *testing.T) {
	a := nasType.NewConfiguredNSSAI(nasMessage.ConfigurationUpdateCommandAllowedNSSAIType)
	for _, table := range nasTypeConfiguredNSSAILenTable {
		a.SetLen(table.in)
		assert.Equal(t, table.out, a.GetLen())
	}
}

type nasTypetConfiguredNSSAISNSSAIValue struct {
	inLen uint8
	in    []uint8
	out   []uint8
}

var nasTypeConfiguredNSSAISNSSAIValueTable = []nasTypetConfiguredNSSAISNSSAIValue{
	{2, []uint8{0x00, 0x01}, []uint8{0x00, 0x1}},
}

func TestNasTypeConfiguredNSSAIGetSetSNSSAIValue(t *testing.T) {
	a := nasType.NewConfiguredNSSAI(0)
	for _, table := range nasTypeConfiguredNSSAISNSSAIValueTable {
		a.SetLen(table.inLen)
		a.SetSNSSAIValue(table.in)
		assert.Equalf(t, table.out, a.GetSNSSAIValue(), "in(%v): out %v, actual %x", table.in, table.out, a.GetSNSSAIValue())
	}
}

type testConfiguredNSSAIDataTemplate struct {
	in  nasType.ConfiguredNSSAI
	out nasType.ConfiguredNSSAI
}

var configuredNSSAITestData = []nasType.ConfiguredNSSAI{
	{nasMessage.ConfigurationUpdateCommandAllowedNSSAIType, 2, []byte{0x00, 0x00}},
}

var configuredNSSAIExpectedData = []nasType.ConfiguredNSSAI{
	{nasMessage.ConfigurationUpdateCommandAllowedNSSAIType, 2, []byte{0x00, 0x00}},
}

var configuredNSSAITestTable = []testConfiguredNSSAIDataTemplate{
	{configuredNSSAITestData[0], configuredNSSAIExpectedData[0]},
}

func TestNasTypeConfiguredNSSAI(t *testing.T) {

	for i, table := range configuredNSSAITestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewConfiguredNSSAI(0)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetSNSSAIValue(table.in.Buffer)

		assert.Equalf(t, table.out.Iei, a.Iei, "in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		assert.Equalf(t, table.out.Len, a.Len, "in(%v): out %v, actual %x", table.in.Len, table.out.Len, a.Len)
		assert.Equalf(t, table.out.Buffer, a.Buffer, "in(%v): out %v, actual %x", table.in.Buffer, table.out.Buffer, a.Buffer)

	}
}
