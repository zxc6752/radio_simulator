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

func TestNasTypeNewPDUSessionID(t *testing.T) {
	a := nasType.NewPDUSessionID()
	assert.NotNil(t, a)

}

var nasTypePDUSessionIDULNASTransportOldPDUSessionIDTypeTable = []NasTypeIeiData{
	{nasMessage.ULNASTransportOldPDUSessionIDType, nasMessage.ULNASTransportOldPDUSessionIDType},
}

func TestNasTypePDUSessionIDGetSetIei(t *testing.T) {
	a := nasType.NewPDUSessionID()
	for _, table := range nasTypePDUSessionIDULNASTransportOldPDUSessionIDTypeTable {
		a.SetPDUSessionID(table.in)
		assert.Equal(t, table.out, a.GetPDUSessionID())
	}
}

type nasTypePDUSessionIDPduSessionIdentity2ValueData struct {
	in  uint8
	out uint8
}

var nasTypePDUSessionIDPduSessionIdentity2ValueTable = []nasTypePDUSessionIDPduSessionIdentity2ValueData{
	{0xff, 0xff},
}

func TestNasTypePDUSessionIDGetSetPduSessionIdentity2Value(t *testing.T) {
	a := nasType.NewPDUSessionID()
	for _, table := range nasTypePDUSessionIDPduSessionIdentity2ValueTable {
		a.SetPDUSessionID(table.in)
		assert.Equal(t, table.out, a.GetPDUSessionID())
	}
}

type testPDUSessionIDDataTemplate struct {
	inPduSessionIdentity2Value  uint8
	outPduSessionIdentity2Value uint8
}

var testPDUSessionIDTestTable = []testPDUSessionIDDataTemplate{
	{0x0f, 0x0f},
}

func TestNasTypePDUSessionID(t *testing.T) {

	for i, table := range testPDUSessionIDTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewPDUSessionID()
		a.SetPDUSessionID(table.inPduSessionIdentity2Value)

		assert.Equalf(t, table.outPduSessionIdentity2Value, a.GetPDUSessionID(), "in(%v): out %v, actual %x", table.inPduSessionIdentity2Value, table.outPduSessionIdentity2Value, a.GetPDUSessionID())
	}
}
