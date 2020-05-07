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

func TestNasTypeNewSORTransparentContainer(t *testing.T) {
	a := nasType.NewSORTransparentContainer(nasMessage.RegistrationAcceptSORTransparentContainerType)
	assert.NotNil(t, a)

}

var nasTypeSORTransparentContainerTable = []NasTypeIeiData{
	{nasMessage.RegistrationAcceptSORTransparentContainerType, nasMessage.RegistrationAcceptSORTransparentContainerType},
}

func TestNasTypeSORTransparentContainerGetSetIei(t *testing.T) {
	a := nasType.NewSORTransparentContainer(nasMessage.RegistrationAcceptSORTransparentContainerType)
	for _, table := range nasTypeSORTransparentContainerTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

type nasTypeSORTransparentContainerLenData struct {
	in  uint16
	out uint16
}

var nasTypeSORTransparentContainerLenTable = []nasTypeSORTransparentContainerLenData{
	{2, 2},
}

func TestNasTypeSORTransparentContainerGetSetLen(t *testing.T) {
	a := nasType.NewSORTransparentContainer(nasMessage.RegistrationAcceptSORTransparentContainerType)
	for _, table := range nasTypeSORTransparentContainerLenTable {
		a.SetLen(table.in)
		assert.Equal(t, table.out, a.GetLen())
	}
}

type nasTypeSORTransparentContainerSORContentData struct {
	inLen uint16
	in    []uint8
	out   []uint8
}

var nasTypeSORTransparentContainerSORContentTable = []nasTypeSORTransparentContainerSORContentData{
	{2, []uint8{0x01, 0x01}, []uint8{0x01, 0x01}},
}

func TestNasTypeSORTransparentContainerGetSetSORContent(t *testing.T) {
	a := nasType.NewSORTransparentContainer(nasMessage.RegistrationAcceptSORTransparentContainerType)
	for _, table := range nasTypeSORTransparentContainerSORContentTable {
		a.SetLen(table.inLen)
		a.SetSORContent(table.in)
		assert.Equalf(t, table.out, a.GetSORContent(), "in(%v): out %v, actual %x", table.in, table.out, a.GetSORContent())
	}
}

type testSORTransparentContainerDataTemplate struct {
	in  nasType.SORTransparentContainer
	out nasType.SORTransparentContainer
}

var SORTransparentContainerTestData = []nasType.SORTransparentContainer{
	{nasMessage.RegistrationAcceptSORTransparentContainerType, 2, []uint8{}},
}

var SORTransparentContainerExpectedTestData = []nasType.SORTransparentContainer{
	{nasMessage.RegistrationAcceptSORTransparentContainerType, 2, []uint8{0x01, 0x01}},
}

var SORTransparentContainerTestTable = []testSORTransparentContainerDataTemplate{
	{SORTransparentContainerTestData[0], SORTransparentContainerExpectedTestData[0]},
}

func TestNasTypeSORTransparentContainer(t *testing.T) {

	for i, table := range SORTransparentContainerTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewSORTransparentContainer(nasMessage.RegistrationAcceptSORTransparentContainerType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetSORContent([]uint8{0x01, 0x01})

		assert.Equalf(t, table.out.Iei, a.Iei, "in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		assert.Equalf(t, table.out.Len, a.Len, "in(%v): out %v, actual %x", table.in.Len, table.out.Len, a.Len)
		assert.Equalf(t, table.out.Buffer, a.Buffer, "in(%v): out %v, actual %x", table.in.Buffer, table.out.Buffer, a.Buffer)

	}
}
