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

var AlwaysonPDUSessionRequestedIeiInput uint8 = 0x0B

func TestNasTypeNewAlwaysonPDUSessionRequested(t *testing.T) {
	a := nasType.NewAlwaysonPDUSessionRequested(AlwaysonPDUSessionRequestedIeiInput)
	assert.NotNil(t, a)
}

var nasTypePDUSessionEstablishmentRequestAlwaysonPDUSessionRequestedTable = []NasTypeIeiData{
	{AlwaysonPDUSessionRequestedIeiInput, 0x0B},
}

func TestNasTypeAlwaysonPDUSessionRequestedGetSetIei(t *testing.T) {
	a := nasType.NewAlwaysonPDUSessionRequested(AlwaysonPDUSessionRequestedIeiInput)
	for _, table := range nasTypePDUSessionEstablishmentRequestAlwaysonPDUSessionRequestedTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

type nasTypeAlwaysonPDUSessionRequestedAPSI struct {
	in  uint8
	out uint8
}

var nasTypeAlwaysonPDUSessionRequestedAPSRTable = []nasTypeAlwaysonPDUSessionRequestedAPSI{
	{0x01, 0x01},
}

func TestNasTypeAlwaysonPDUSessionRequestedGetSetAPSR(t *testing.T) {
	a := nasType.NewAlwaysonPDUSessionRequested(nasMessage.PDUSessionEstablishmentRequestAlwaysonPDUSessionRequestedType)
	for _, table := range nasTypeAlwaysonPDUSessionRequestedAPSRTable {
		a.SetAPSR(table.in)
		assert.Equal(t, table.out, a.GetAPSR())
	}
}

type testAlwaysonPDUSessionRequestedDataTemplate struct {
	in  nasType.AlwaysonPDUSessionRequested
	out nasType.AlwaysonPDUSessionRequested
}

var alwaysonPDUSessionRequestedTestData = []nasType.AlwaysonPDUSessionRequested{
	{(0xB0 + 0x01)},
}

var alwaysonPDUSessionRequestedExpectedTestData = []nasType.AlwaysonPDUSessionRequested{
	{(0xB0 + 0x01)},
}

var alwaysonPDUSessionRequestedTestTable = []testAlwaysonPDUSessionRequestedDataTemplate{
	{alwaysonPDUSessionRequestedTestData[0], alwaysonPDUSessionRequestedExpectedTestData[0]},
}

func TestNasTypeAlwaysonPDUSessionRequested(t *testing.T) {

	for _, table := range alwaysonPDUSessionRequestedTestTable {
		a := nasType.NewAlwaysonPDUSessionRequested(AlwaysonPDUSessionRequestedIeiInput)

		a.SetIei(AlwaysonPDUSessionRequestedIeiInput)
		a.SetAPSR(table.in.GetAPSR())

		assert.Equal(t, table.out.Octet, a.Octet)

	}
}
