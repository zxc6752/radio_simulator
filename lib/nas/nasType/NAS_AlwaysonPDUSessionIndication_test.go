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

var PDUSessionEstablishmentAcceptAlwaysonPDUSessionIndicationTypeIeiInput uint8 = 0x08

func TestNasTypeNewAlwaysonPDUSessionIndication(t *testing.T) {
	a := nasType.NewAlwaysonPDUSessionIndication(PDUSessionEstablishmentAcceptAlwaysonPDUSessionIndicationTypeIeiInput)
	assert.NotNil(t, a)
}

var nasTypePDUSessionEstablishmentRequestAlwaysonPDUSessionIndicationTable = []NasTypeIeiData{
	{PDUSessionEstablishmentAcceptAlwaysonPDUSessionIndicationTypeIeiInput, 0x08},
}

func TestNasTypeAlwaysonPDUSessionIndicationGetSetIei(t *testing.T) {
	a := nasType.NewAlwaysonPDUSessionIndication(nasMessage.PDUSessionEstablishmentAcceptAlwaysonPDUSessionIndicationType)
	for _, table := range nasTypePDUSessionEstablishmentRequestAlwaysonPDUSessionIndicationTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

type nasTypeAlwaysonPDUSessionIndicationAPSI struct {
	in  uint8
	out uint8
}

var nasTypeAlwaysonPDUSessionIndicationAPSITable = []nasTypeAlwaysonPDUSessionIndicationAPSI{
	{0x01, 0x01},
}

func TestNasTypeAlwaysonPDUSessionIndicationGetSetAPSI(t *testing.T) {
	a := nasType.NewAlwaysonPDUSessionIndication(nasMessage.PDUSessionEstablishmentAcceptAlwaysonPDUSessionIndicationType)
	for _, table := range nasTypeAlwaysonPDUSessionIndicationAPSITable {
		a.SetAPSI(table.in)
		assert.Equal(t, table.out, a.GetAPSI())
	}
}

type testAlwaysonPDUSessionIndicationDataTemplate struct {
	in  nasType.AlwaysonPDUSessionIndication
	out nasType.AlwaysonPDUSessionIndication
}

var alwaysonPDUSessionIndicationTestData = []nasType.AlwaysonPDUSessionIndication{
	{(0x80 + 0x01)},
}

var alwaysonPDUSessionIndicationExpectedTestData = []nasType.AlwaysonPDUSessionIndication{
	{(0x80 + 0x01)},
}

var alwaysonPDUSessionIndicationTestTable = []testAlwaysonPDUSessionIndicationDataTemplate{
	{alwaysonPDUSessionIndicationTestData[0], alwaysonPDUSessionIndicationExpectedTestData[0]},
}

func TestNasTypeAlwaysonPDUSessionIndication(t *testing.T) {

	for _, table := range alwaysonPDUSessionIndicationTestTable {
		a := nasType.NewAlwaysonPDUSessionIndication(PDUSessionEstablishmentAcceptAlwaysonPDUSessionIndicationTypeIeiInput)

		a.SetIei(PDUSessionEstablishmentAcceptAlwaysonPDUSessionIndicationTypeIeiInput)
		a.SetAPSI(table.in.GetAPSI())

		assert.Equal(t, table.out.Octet, a.Octet)

	}
}
