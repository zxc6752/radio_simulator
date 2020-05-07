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

func TestNasTypeNewTAIList(t *testing.T) {
	a := nasType.NewTAIList(nasMessage.RegistrationAcceptTAIListType)
	assert.NotNil(t, a)

}

var nasTypeTAIListTable = []NasTypeIeiData{
	{nasMessage.RegistrationAcceptTAIListType, nasMessage.RegistrationAcceptTAIListType},
}

func TestNasTypeTAIListGetSetIei(t *testing.T) {
	a := nasType.NewTAIList(nasMessage.RegistrationAcceptTAIListType)
	for _, table := range nasTypeTAIListTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

var nasTypeTAIListLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeTAIListGetSetLen(t *testing.T) {
	a := nasType.NewTAIList(nasMessage.RegistrationAcceptTAIListType)
	for _, table := range nasTypeTAIListLenTable {
		a.SetLen(table.in)
		assert.Equal(t, table.out, a.GetLen())
	}
}

type nasTypeTAIListPartialTrackingAreaIdentityListData struct {
	inLen uint8
	in    []uint8
	out   []uint8
}

var nasTypeTAIListPartialTrackingAreaIdentityListTable = []nasTypeTAIListPartialTrackingAreaIdentityListData{
	{2, []uint8{0x01, 0x01}, []uint8{0x01, 0x01}},
}

func TestNasTypeTAIListGetSetPartialTrackingAreaIdentityList(t *testing.T) {
	a := nasType.NewTAIList(nasMessage.RegistrationAcceptTAIListType)
	for _, table := range nasTypeTAIListPartialTrackingAreaIdentityListTable {
		a.SetLen(table.inLen) // fix it, set input length
		a.SetPartialTrackingAreaIdentityList(table.in)
		assert.Equalf(t, table.out, a.GetPartialTrackingAreaIdentityList(), "in(%v): out %v, actual %x", table.in, table.out, a.GetPartialTrackingAreaIdentityList())
	}
}

type testTAIListDataTemplate struct {
	in  nasType.TAIList
	out nasType.TAIList
}

var TAIListTestData = []nasType.TAIList{
	{nasMessage.RegistrationAcceptTAIListType, 2, []uint8{}},
}

var TAIListExpectedTestData = []nasType.TAIList{
	{nasMessage.RegistrationAcceptTAIListType, 2, []uint8{0x01, 0x01}},
}

var TAIListTestTable = []testTAIListDataTemplate{
	{TAIListTestData[0], TAIListExpectedTestData[0]},
}

func TestNasTypeTAIList(t *testing.T) {

	for i, table := range TAIListTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewTAIList(nasMessage.RegistrationAcceptTAIListType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetPartialTrackingAreaIdentityList([]uint8{0x01, 0x01})

		assert.Equalf(t, table.out.Iei, a.Iei, "in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		assert.Equalf(t, table.out.Len, a.Len, "in(%v): out %v, actual %x", table.in.Len, table.out.Len, a.Len)
		assert.Equalf(t, table.out.Buffer, a.Buffer, "in(%v): out %v, actual %x", table.in.Buffer, table.out.Buffer, a.Buffer)

	}
}
