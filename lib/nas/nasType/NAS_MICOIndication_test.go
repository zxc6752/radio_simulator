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
	"radio_simulator/lib/nas/nasType"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ConfigurationUpdateCommandMICOIndicationTypeIeiInput uint8 = 0x0B

func TestNasTypeNewMICOIndication(t *testing.T) {
	a := nasType.NewMICOIndication(ConfigurationUpdateCommandMICOIndicationTypeIeiInput)
	assert.NotNil(t, a)
}

var nasTypeConfigurationUpdateCommandMICOIndicationTable = []NasTypeIeiData{
	{ConfigurationUpdateCommandMICOIndicationTypeIeiInput, 0x0B},
}

func TestNasTypeMICOIndicationGetSetIei(t *testing.T) {
	a := nasType.NewMICOIndication(ConfigurationUpdateCommandMICOIndicationTypeIeiInput)
	for _, table := range nasTypeConfigurationUpdateCommandMICOIndicationTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

type nasTypeMICOIndicationRAAI struct {
	in  uint8
	out uint8
}

var nasTypeMICOIndicationRAAITable = []nasTypeMICOIndicationRAAI{
	{0x01, 0x01},
}

func TestNasTypeMICOIndicationGetSetRAAI(t *testing.T) {
	a := nasType.NewMICOIndication(ConfigurationUpdateCommandMICOIndicationTypeIeiInput)
	for _, table := range nasTypeMICOIndicationRAAITable {
		a.SetRAAI(table.in)
		assert.Equal(t, table.out, a.GetRAAI())
	}
}

type testMICOIndicationDataTemplate struct {
	inRAAI uint8
	in     nasType.MICOIndication
	out    nasType.MICOIndication
}

var mICOIndicationTestData = []nasType.MICOIndication{
	{(0xB0 + 0x01)},
}

var mICOIndicationExpectedData = []nasType.MICOIndication{
	{(0xB0 + 0x01)},
}

var mICOIndicationTestTable = []testMICOIndicationDataTemplate{
	{0x01, mICOIndicationTestData[0], mICOIndicationExpectedData[0]},
}

func TestNasTypeMICOIndication(t *testing.T) {

	for _, table := range mICOIndicationTestTable {
		a := nasType.NewMICOIndication(ConfigurationUpdateCommandMICOIndicationTypeIeiInput)

		a.SetIei(ConfigurationUpdateCommandMICOIndicationTypeIeiInput)
		a.SetRAAI(table.inRAAI)

		assert.Equal(t, table.out.Octet, a.Octet)

	}
}
