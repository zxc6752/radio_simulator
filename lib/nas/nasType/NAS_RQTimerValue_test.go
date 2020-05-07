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

func TestNasTypeNewRQTimerValue(t *testing.T) {
	a := nasType.NewRQTimerValue(nasMessage.PDUSessionEstablishmentAcceptRQTimerValueType)
	assert.NotNil(t, a)
}

var nasTypePDUSessionReleaseCompleteRQTimerValueTable = []NasTypeIeiData{
	{nasMessage.PDUSessionEstablishmentAcceptRQTimerValueType, nasMessage.PDUSessionEstablishmentAcceptRQTimerValueType},
}

func TestNasTypeRQTimerValueGetSetIei(t *testing.T) {
	a := nasType.NewRQTimerValue(nasMessage.PDUSessionEstablishmentAcceptRQTimerValueType)
	for _, table := range nasTypePDUSessionReleaseCompleteRQTimerValueTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

var nasTypeRQTimerValueUintTable = []NasTypeLenuint8Data{
	{0x01, 0x01},
}

func TestNasTypeRQTimerValueGetSetUint(t *testing.T) {
	a := nasType.NewRQTimerValue(nasMessage.PDUSessionEstablishmentAcceptRQTimerValueType)
	for _, table := range nasTypeRQTimerValueUintTable {
		a.SetUnit(table.in)
		assert.Equal(t, table.out, a.GetUnit())
	}
}

type nasTypeRQTimerValueTimerValueData struct {
	in  uint8
	out uint8
}

var nasTypeRQTimerValueTimerValueTable = []nasTypeRQTimerValueTimerValueData{
	{0x01, 0x01},
}

func TestNasTypeRQTimerValueGetSetTimerValue(t *testing.T) {
	a := nasType.NewRQTimerValue(nasMessage.PDUSessionEstablishmentAcceptRQTimerValueType)
	for _, table := range nasTypeRQTimerValueTimerValueTable {
		a.SetTimerValue(table.in)
		assert.Equal(t, table.out, a.GetTimerValue())
	}
}

type testRQTimerValueDataTemplate struct {
	inUnit       uint8
	inTimerValue uint8
	in           nasType.RQTimerValue
	out          nasType.RQTimerValue
}

var rQTimerValueTestData = []nasType.RQTimerValue{
	{nasMessage.PDUSessionEstablishmentAcceptRQTimerValueType, 0x01},
}

var rQTimerValueExpectedTestData = []nasType.RQTimerValue{
	{nasMessage.PDUSessionEstablishmentAcceptRQTimerValueType, 0x21},
}

var rQTimerValueTestTable = []testRQTimerValueDataTemplate{
	{0x01, 0x01, rQTimerValueTestData[0], rQTimerValueExpectedTestData[0]},
}

func TestNasTypeRQTimerValue(t *testing.T) {

	for _, table := range rQTimerValueTestTable {
		a := nasType.NewRQTimerValue(nasMessage.PDUSessionEstablishmentAcceptRQTimerValueType)

		a.SetIei(table.in.GetIei())
		a.SetUnit(table.inUnit)
		a.SetTimerValue(table.inTimerValue)

		assert.Equalf(t, table.out.Iei, a.Iei, "in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		assert.Equalf(t, table.out.Octet, a.Octet, "in(%v): out %v, actual %x", table.in.Octet, table.out.Octet, a.Octet)

	}
}
