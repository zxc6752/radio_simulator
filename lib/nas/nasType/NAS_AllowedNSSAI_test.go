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

func TestNasTypeNewAllowedNSSAI(t *testing.T) {
	a := nasType.NewAllowedNSSAI(nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType)
	assert.NotNil(t, a)
}

var nasTypeConfigurationUpdateCommandConfiguredNSSAITable = []NasTypeIeiData{
	{nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType, nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType},
}

func TestNasTypeAllowedNSSAIGetSetIei(t *testing.T) {
	a := nasType.NewAllowedNSSAI(nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType)
	for _, table := range nasTypeConfigurationUpdateCommandConfiguredNSSAITable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

var nasTypeConfigurationUpdateCommandConfiguredNSSAILenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeAllowedNSSAIGetSetLen(t *testing.T) {
	a := nasType.NewAllowedNSSAI(nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType)
	for _, table := range nasTypeConfigurationUpdateCommandConfiguredNSSAILenTable {
		a.SetLen(table.in)
		assert.Equal(t, table.out, a.GetLen())
	}
}

type SNSSAIValue struct {
	inLen uint8
	in    []uint8
	out   []uint8
}

var nasTypeSNSSAIValueTable = []SNSSAIValue{
	{2, []uint8{0x00, 0x01}, []uint8{0x00, 0x01}},
}

func TestNasTypeAllowedNSSAIGetSetSNSSAIValue(t *testing.T) {
	a := nasType.NewAllowedNSSAI(nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType)
	for _, table := range nasTypeSNSSAIValueTable {
		a.SetLen(table.inLen)
		a.SetSNSSAIValue(table.in)
		assert.Equalf(t, table.out, a.GetSNSSAIValue(), "in(%v): out %v, actual %x", table.in, table.out, a.GetSNSSAIValue())
	}
}

type testAllowedNSSAIDataTemplate struct {
	in  nasType.AllowedNSSAI
	out nasType.AllowedNSSAI
}

var AllowedNSSAITestData = []nasType.AllowedNSSAI{
	{nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType, 2, []uint8{0x00, 0x01}},
}

var AllowedNSSAIExpectedTestData = []nasType.AllowedNSSAI{
	{nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType, 2, []uint8{0x00, 0x01}},
}

var AllowedNSSAITable = []testAllowedNSSAIDataTemplate{
	{AllowedNSSAITestData[0], AllowedNSSAIExpectedTestData[0]},
}

func TestNasTypeAllowedNSSAI(t *testing.T) {
	for i, table := range AllowedNSSAITable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewAllowedNSSAI(nasMessage.ConfigurationUpdateCommandConfiguredNSSAIType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetSNSSAIValue(table.in.Buffer)

		assert.Equalf(t, table.out.Iei, a.Iei, "in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		assert.Equalf(t, table.out.Len, a.Len, "in(%v): out %v, actual %x", table.in.Len, table.out.Len, a.Len)
		assert.Equalf(t, table.out.Buffer, a.Buffer, "in(%v): out %v, actual %x", table.in.Buffer, table.out.Buffer, a.Buffer)
	}

}
