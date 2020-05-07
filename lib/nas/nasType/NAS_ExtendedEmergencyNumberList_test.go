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

func TestNasTypeNewExtendedEmergencyNumberList(t *testing.T) {
	a := nasType.NewExtendedEmergencyNumberList(nasMessage.RegistrationAcceptExtendedEmergencyNumberListType)
	assert.NotNil(t, a)

}

var nasTypeRegistrationAcceptExtendedEmergencyNumberListIeiTable = []NasTypeIeiData{
	{nasMessage.RegistrationAcceptExtendedEmergencyNumberListType, nasMessage.RegistrationAcceptExtendedEmergencyNumberListType},
}

func TestNasTypeExtendedEmergencyNumberListGetSetIei(t *testing.T) {
	a := nasType.NewExtendedEmergencyNumberList(nasMessage.RegistrationAcceptExtendedEmergencyNumberListType)
	for _, table := range nasTypeRegistrationAcceptExtendedEmergencyNumberListIeiTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

var nasTypeExtendedEmergencyNumberListLenTable = []NasTypeLenUint16Data{
	{4, 4},
}

func TestNasTypeExtendedEmergencyNumberListGetSetLen(t *testing.T) {
	a := nasType.NewExtendedEmergencyNumberList(nasMessage.RegistrationAcceptExtendedEmergencyNumberListType)
	for _, table := range nasTypeExtendedEmergencyNumberListLenTable {
		a.SetLen(table.in)
		assert.Equal(t, table.out, a.GetLen())
	}
}

type nasTypetExtendedEmergencyNumberListEENL struct {
	inLen uint16
	in    uint8
	out   uint8
}

var nasTypeExtendedEmergencyNumberListEENLTable = []nasTypetExtendedEmergencyNumberListEENL{
	{2, 0x01, 0x01},
}

func TestNasTypeExtendedEmergencyNumberListGetSetEENL(t *testing.T) {
	a := nasType.NewExtendedEmergencyNumberList(nasMessage.RegistrationAcceptExtendedEmergencyNumberListType)
	for _, table := range nasTypeExtendedEmergencyNumberListEENLTable {
		a.SetLen(table.inLen)
		a.SetEENL(table.in)
		assert.Equalf(t, table.out, a.GetEENL(), "in(%v): out %v, actual %x", table.in, table.out, a.GetEENL())
	}
}

type nasTypetExtendedEmergencyNumberListEmergencyInformation struct {
	inLen uint16
	in    []uint8
	out   []uint8
}

var nasTypeExtendedEmergencyNumberListEmergencyInformationTable = []nasTypetExtendedEmergencyNumberListEmergencyInformation{
	{3, []uint8{0x00, 0x00, 0x01}, []uint8{0x00, 0x00, 0x01}},
}

func TestNasTypeExtendedEmergencyNumberListGetSetExtendedEmergencyNumberList(t *testing.T) {
	a := nasType.NewExtendedEmergencyNumberList(0)
	for _, table := range nasTypeExtendedEmergencyNumberListEmergencyInformationTable {
		a.SetLen(table.inLen)
		a.SetEmergencyInformation(table.in)
		assert.Equalf(t, table.out, a.GetEmergencyInformation(), "in(%v): out %v, actual %x", table.in, table.out, a.GetEmergencyInformation())
		assert.Equalf(t, table.out, a.Buffer, "in(%v): out %v, actual %x", table.in, table.out, a.Buffer)
	}
}
