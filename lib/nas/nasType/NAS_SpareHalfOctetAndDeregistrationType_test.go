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

func TestNasTypeNewNewSpareHalfOctetAndDeregistrationType(t *testing.T) {
	a := nasType.NewSpareHalfOctetAndDeregistrationType()
	assert.NotNil(t, a)
}

type nasTypeDeregistrationTypeAndSpareHalfOctetSwitchOff struct {
	in  uint8
	out uint8
}

var nasTypeDeregistrationTypeAndSpareHalfOctetSwitchOffTable = []nasTypeDeregistrationTypeAndSpareHalfOctetSwitchOff{
	{0x01, 0x01},
}

func TestNasTypeDeregistrationTypeAndSpareHalfOctetGetSetSwitchOff(t *testing.T) {
	a := nasType.NewSpareHalfOctetAndDeregistrationType()
	for _, table := range nasTypeDeregistrationTypeAndSpareHalfOctetSwitchOffTable {
		a.SetSwitchOff(table.in)
		assert.Equal(t, table.out, a.GetSwitchOff())
	}
}

type nasTypeDeregistrationTypeAndSpareHalfOctetReRegistrationRequired struct {
	in  uint8
	out uint8
}

var nasTypeDeregistrationTypeAndSpareHalfOctetReRegistrationRequiredTable = []nasTypeDeregistrationTypeAndSpareHalfOctetReRegistrationRequired{
	{0x01, 0x01},
}

func TestNasTypeDeregistrationTypeAndSpareHalfOctetGetSetReRegistrationRequired(t *testing.T) {
	a := nasType.NewSpareHalfOctetAndDeregistrationType()
	for _, table := range nasTypeDeregistrationTypeAndSpareHalfOctetReRegistrationRequiredTable {
		a.SetReRegistrationRequired(table.in)
		assert.Equal(t, table.out, a.GetReRegistrationRequired())
	}
}

type nasTypeDeregistrationTypeAndSpareHalfOctetAccessType struct {
	in  uint8
	out uint8
}

var nasTypeDeregistrationTypeAndSpareHalfOctetAccessTypeTable = []nasTypeDeregistrationTypeAndSpareHalfOctetAccessType{
	{0x03, 0x3},
}

func TestNasTypeDeregistrationTypeAndSpareHalfOctetGetSetAccessType(t *testing.T) {
	a := nasType.NewSpareHalfOctetAndDeregistrationType()
	for _, table := range nasTypeDeregistrationTypeAndSpareHalfOctetAccessTypeTable {
		a.SetAccessType(table.in)
		assert.Equal(t, table.out, a.GetAccessType())
	}
}

type testDeregistrationTypeAndSpareHalfOctetDataTemplate struct {
	inSwitchOff              uint8
	inReRegistrationRequired uint8
	inAccessType             uint8
	in                       nasType.SpareHalfOctetAndDeregistrationType
	out                      nasType.SpareHalfOctetAndDeregistrationType
}

var deregistrationTypeAndSpareHalfOctetTestData = []nasType.SpareHalfOctetAndDeregistrationType{
	{0x0f},
}

var deregistrationTypeAndSpareHalfOctetExpectedData = []nasType.SpareHalfOctetAndDeregistrationType{
	{0xf},
}

var deregistrationTypeAndSpareHalfOctetTestTable = []testDeregistrationTypeAndSpareHalfOctetDataTemplate{
	{0x01, 0x01, 0x03, deregistrationTypeAndSpareHalfOctetTestData[0], deregistrationTypeAndSpareHalfOctetExpectedData[0]},
}

func TestNasTypeDeregistrationTypeAndSpareHalfOctet(t *testing.T) {

	for _, table := range deregistrationTypeAndSpareHalfOctetTestTable {
		a := nasType.NewSpareHalfOctetAndDeregistrationType()

		a.SetSwitchOff(table.inSwitchOff)
		a.SetReRegistrationRequired(table.inReRegistrationRequired)
		a.SetAccessType(table.inAccessType)

		assert.Equal(t, table.out.Octet, a.Octet)

	}
}
