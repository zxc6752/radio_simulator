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

var ConfigurationUpdateCommandConfigurationUpdateIndicationTypeIeiInput uint8 = 0x0D

func TestNasTypeNewConfigurationUpdateIndication(t *testing.T) {
	a := nasType.NewConfigurationUpdateIndication(ConfigurationUpdateCommandConfigurationUpdateIndicationTypeIeiInput)
	assert.NotNil(t, a)
}

var nasTypePDUSessionEstablishmentRequestConfigurationUpdateIndicationTable = []NasTypeIeiData{
	{ConfigurationUpdateCommandConfigurationUpdateIndicationTypeIeiInput, 0x0D},
}

func TestNasTypeConfigurationUpdateIndicationGetSetIei(t *testing.T) {
	a := nasType.NewConfigurationUpdateIndication(ConfigurationUpdateCommandConfigurationUpdateIndicationTypeIeiInput)
	for _, table := range nasTypePDUSessionEstablishmentRequestConfigurationUpdateIndicationTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

type nasTypeConfigurationUpdateIndicationRED struct {
	in  uint8
	out uint8
}

var nasTypeConfigurationUpdateIndicationREDTable = []nasTypeConfigurationUpdateIndicationRED{
	{0x01, 0x01},
}

func TestNasTypeConfigurationUpdateIndicationGetSetRED(t *testing.T) {
	a := nasType.NewConfigurationUpdateIndication(ConfigurationUpdateCommandConfigurationUpdateIndicationTypeIeiInput)
	for _, table := range nasTypeConfigurationUpdateIndicationREDTable {
		a.SetRED(table.in)
		assert.Equal(t, table.out, a.GetRED())
	}
}

type nasTypeConfigurationUpdateIndicationACK struct {
	in  uint8
	out uint8
}

var nasTypeConfigurationUpdateIndicationACKTable = []nasTypeConfigurationUpdateIndicationACK{
	{0x01, 0x01},
}

func TestNasTypeConfigurationUpdateIndicationGetSetACK(t *testing.T) {
	a := nasType.NewConfigurationUpdateIndication(ConfigurationUpdateCommandConfigurationUpdateIndicationTypeIeiInput)
	for _, table := range nasTypeConfigurationUpdateIndicationACKTable {
		a.SetACK(table.in)
		assert.Equal(t, table.out, a.GetACK())
	}
}

type testConfigurationUpdateIndicationDataTemplate struct {
	inRED uint8
	inACK uint8
	in    nasType.ConfigurationUpdateIndication
	out   nasType.ConfigurationUpdateIndication
}

var configurationUpdateIndicationTestData = []nasType.ConfigurationUpdateIndication{
	{(0xD0 + 0x03)},
}

var configurationUpdateIndicationExpectedData = []nasType.ConfigurationUpdateIndication{
	{(0xD0 + 0x03)},
}

var configurationUpdateIndicationTestTable = []testConfigurationUpdateIndicationDataTemplate{
	{0x01, 0x01, configurationUpdateIndicationTestData[0], configurationUpdateIndicationExpectedData[0]},
}

func TestNasTypeConfigurationUpdateIndication(t *testing.T) {

	for _, table := range configurationUpdateIndicationTestTable {
		a := nasType.NewConfigurationUpdateIndication(ConfigurationUpdateCommandConfigurationUpdateIndicationTypeIeiInput)

		a.SetIei(ConfigurationUpdateCommandConfigurationUpdateIndicationTypeIeiInput)
		a.SetRED(table.inRED)
		a.SetACK(table.inACK)

		assert.Equal(t, table.out.Octet, a.Octet)

	}
}
