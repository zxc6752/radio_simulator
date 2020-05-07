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

func TestNasTypeNewUEStatus(t *testing.T) {
	a := nasType.NewUEStatus(nasMessage.RegistrationRequestUEStatusType)
	assert.NotNil(t, a)
}

var nasTypeUEStatusIeiTable = []NasTypeIeiData{
	{nasMessage.RegistrationRequestUEStatusType, nasMessage.RegistrationRequestUEStatusType},
}

func TestNasTypeUEStatusGetSetIei(t *testing.T) {
	a := nasType.NewUEStatus(nasMessage.RegistrationRequestUEStatusType)
	for _, table := range nasTypeRegistrationAcceptNetworkFeatureSupport5GSTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

var nasTypeUEStatusLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeUEStatusGetSetLen(t *testing.T) {
	a := nasType.NewUEStatus(nasMessage.RegistrationRequestUEStatusType)
	for _, table := range nasTypeRegistrationAcceptNetworkFeatureSupport5GSLenTable {
		a.SetLen(table.in)
		assert.Equal(t, table.out, a.GetLen())
	}
}

type NasTypeUEStatusN1ModeRegData struct {
	in  uint8
	out uint8
}

var nasTypeUEStatusN1ModeRegTable = []NasTypeUEStatusN1ModeRegData{
	{0x01, 0x01},
}

func TestNasTypeUEStatusGetSetN1ModeReg(t *testing.T) {
	a := nasType.NewUEStatus(nasMessage.RegistrationRequestUEStatusType)
	for _, table := range nasTypeUEStatusN1ModeRegTable {
		a.SetN1ModeReg(table.in)
		assert.Equal(t, table.out, a.GetN1ModeReg())
	}
}

type NasTypeUEStatusS1ModeRegData struct {
	in  uint8
	out uint8
}

var nasTypeUEStatusS1ModeRegTable = []NasTypeUEStatusS1ModeRegData{
	{0x01, 0x01},
}

func TestNasTypeUEStatusGetSetS1ModeReg(t *testing.T) {
	a := nasType.NewUEStatus(nasMessage.RegistrationRequestUEStatusType)
	for _, table := range nasTypeUEStatusS1ModeRegTable {
		a.SetS1ModeReg(table.in)
		assert.Equal(t, table.out, a.GetS1ModeReg())
	}
}

type testUEStatusDataTemplate struct {
	in  nasType.UEStatus
	out nasType.UEStatus
}

var UEStatusTestData = []nasType.UEStatus{
	{nasMessage.RegistrationRequestUEStatusType, 1, 0x05},
}
var UEStatusExpectedData = []nasType.UEStatus{
	{nasMessage.RegistrationRequestUEStatusType, 1, 0x03},
}

var UEStatusDataTestTable = []testUEStatusDataTemplate{
	{UEStatusTestData[0], UEStatusExpectedData[0]},
}

func TestNasTypeUEStatus(t *testing.T) {
	for _, table := range UEStatusDataTestTable {
		a := nasType.NewUEStatus(nasMessage.RegistrationRequestUEStatusType)
		a.SetIei(table.in.Iei)
		a.SetLen(table.in.Len)
		a.SetN1ModeReg(0x01)
		a.SetS1ModeReg(0x01)
		assert.Equal(t, table.out.Iei, a.Iei)
		assert.Equal(t, table.out.Len, a.Len)
		assert.Equal(t, table.out.Octet, a.Octet)
	}
}
