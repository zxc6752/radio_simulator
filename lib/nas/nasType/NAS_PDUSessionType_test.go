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

var PDUSessionEstablishmentRequestPDUSessionTypeTypeIeiInput uint8 = 0x09

func TestNasTypeNewPDUSessionType(t *testing.T) {
	a := nasType.NewPDUSessionType(PDUSessionEstablishmentRequestPDUSessionTypeTypeIeiInput)
	assert.NotNil(t, a)
}

var nasTypePDUSessionEstablishmentRequestPDUSessionTypeTable = []NasTypeIeiData{
	{PDUSessionEstablishmentRequestPDUSessionTypeTypeIeiInput, 0x09},
}

func TestNasTypePDUSessionTypeGetSetIei(t *testing.T) {
	a := nasType.NewPDUSessionType(PDUSessionEstablishmentRequestPDUSessionTypeTypeIeiInput)
	for _, table := range nasTypePDUSessionEstablishmentRequestPDUSessionTypeTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

var nasTypePDUSessionEstablishmentRequestPDUSessionTypeSpareTable = []NasTypeLenuint8Data{
	{0x1, 0x1},
}

func TestNasTypePDUSessionTypeGetSetSpare(t *testing.T) {
	a := nasType.NewPDUSessionType(PDUSessionEstablishmentRequestPDUSessionTypeTypeIeiInput)
	for _, table := range nasTypePDUSessionEstablishmentRequestPDUSessionTypeSpareTable {
		a.SetSpare(table.in)
		assert.Equal(t, table.out, a.GetSpare())
	}
}

var nasTypePDUSessionTypeValue = []NasTypeLenuint8Data{
	{0x0, 0x0},
	{0x1, 0x1},
	{0x2, 0x2},
	{0x3, 0x3},
	{0x4, 0x4},
	{0x5, 0x5},
}

func TestNasTypePDUSessionTypeGetSetPDUSessionTypeValue(t *testing.T) {
	a := nasType.NewPDUSessionType(PDUSessionEstablishmentRequestPDUSessionTypeTypeIeiInput)
	for _, table := range nasTypePDUSessionTypeValue {
		a.SetPDUSessionTypeValue(table.in)
		assert.Equal(t, table.out, a.GetPDUSessionTypeValue())
	}
}

type testPDUSessionTypeDataTemplate struct {
	inPDUSessionTypeValue uint8
	in                    nasType.PDUSessionType
	out                   nasType.PDUSessionType
}

var pDUSessionTypeTestData = []nasType.PDUSessionType{
	{(nasMessage.PDUSessionEstablishmentRequestPDUSessionTypeType)},
}

var pDUSessionTypeExpectedData = []nasType.PDUSessionType{
	{(0x90 + 0x01)},
}

var pDUSessionTypeTestTable = []testPDUSessionTypeDataTemplate{
	{0x01, pDUSessionTypeTestData[0], pDUSessionTypeExpectedData[0]},
}

func TestNasTypePDUSessionType(t *testing.T) {

	for _, table := range pDUSessionTypeTestTable {
		a := nasType.NewPDUSessionType(PDUSessionEstablishmentRequestPDUSessionTypeTypeIeiInput)

		a.SetIei(PDUSessionEstablishmentRequestPDUSessionTypeTypeIeiInput)
		a.SetPDUSessionTypeValue(table.inPDUSessionTypeValue)

		assert.Equal(t, table.out.Octet, a.Octet)

	}
}
