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

func TestNasTypeNewEPSNASMessageContainer(t *testing.T) {
	a := nasType.NewEPSNASMessageContainer(nasMessage.RegistrationRequestEPSNASMessageContainerType)
	assert.NotNil(t, a)

}

var nasTypeRegistrationRequestEPSNASMessageContainerIeiTable = []NasTypeIeiData{
	{nasMessage.RegistrationRequestEPSNASMessageContainerType, nasMessage.RegistrationRequestEPSNASMessageContainerType},
}

func TestNasTypeEPSNASMessageContainerGetSetIei(t *testing.T) {
	a := nasType.NewEPSNASMessageContainer(nasMessage.RegistrationRequestEPSNASMessageContainerType)
	for _, table := range nasTypeRegistrationRequestEPSNASMessageContainerIeiTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

var nasTypeEPSNASMessageContainerLenTable = []NasTypeLenUint16Data{
	{2, 2},
}

func TestNasTypeEPSNASMessageContainerGetSetLen(t *testing.T) {
	a := nasType.NewEPSNASMessageContainer(nasMessage.RegistrationRequestEPSNASMessageContainerType)
	for _, table := range nasTypeEPSNASMessageContainerLenTable {
		a.SetLen(table.in)
		assert.Equal(t, table.out, a.GetLen())
	}
}

type nasTypeEPSNASMessageContainerEPANASMessageContainer struct {
	inLen uint16
	in    []uint8
	out   []uint8
}

var nasTypeEPSNASMessageContainerEPANASMessageContainerTable = []nasTypeEPSNASMessageContainerEPANASMessageContainer{
	{2, []uint8{0x01, 0x01}, []uint8{0x01, 0x01}},
}

func TestNasTypeEPSNASMessageContainerGetSetEPANASMessageContainer(t *testing.T) {
	a := nasType.NewEPSNASMessageContainer(nasMessage.RegistrationRequestEPSNASMessageContainerType)
	for _, table := range nasTypeEPSNASMessageContainerEPANASMessageContainerTable {
		a.SetLen(table.inLen)
		a.SetEPANASMessageContainer(table.in)
		assert.Equalf(t, table.out, a.GetEPANASMessageContainer(), "in(%v): out %v, actual %x", table.in, table.out, a.GetEPANASMessageContainer())
	}
}

type testEPSNASMessageContainerDataTemplate struct {
	in  nasType.EPSNASMessageContainer
	out nasType.EPSNASMessageContainer
}

var ePSNASMessageContainerTestData = []nasType.EPSNASMessageContainer{
	{nasMessage.RegistrationRequestEPSNASMessageContainerType, 3, []byte{0x02, 0x1f, 0x22}},
}

var ePSNASMessageContainerExpectedData = []nasType.EPSNASMessageContainer{
	{nasMessage.RegistrationRequestEPSNASMessageContainerType, 3, []byte{0x02, 0x1f, 0x22}},
}

var ePSNASMessageContainerTestTable = []testEPSNASMessageContainerDataTemplate{
	{ePSNASMessageContainerTestData[0], ePSNASMessageContainerExpectedData[0]},
}

func TestNasTypeEPSNASMessageContainer(t *testing.T) {

	for i, table := range ePSNASMessageContainerTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewEPSNASMessageContainer(0)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetEPANASMessageContainer(table.in.Buffer)

		assert.Equalf(t, table.out.Iei, a.Iei, "in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		assert.Equalf(t, table.out.Len, a.Len, "in(%v): out %v, actual %x", table.in.Len, table.out.Len, a.Len)
		assert.Equalf(t, table.out.Buffer, a.Buffer, "in(%v): out %v, actual %x", table.in.Buffer, table.out.Buffer, a.Buffer)

	}
}
