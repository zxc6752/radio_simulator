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

func TestNasTypeNewCause5GSM(t *testing.T) {
	a := nasType.NewCause5GSM(nasMessage.PDUSessionReleaseCompleteCause5GSMType)
	assert.NotNil(t, a)

}

var nasTypePDUSessionReleaseCompleteCause5GSMTable = []NasTypeIeiData{
	{nasMessage.PDUSessionReleaseCompleteCause5GSMType, nasMessage.PDUSessionReleaseCompleteCause5GSMType},
}

func TestNasTypeCause5GSMGetSetIei(t *testing.T) {
	a := nasType.NewCause5GSM(nasMessage.PDUSessionReleaseCompleteCause5GSMType)
	for _, table := range nasTypePDUSessionReleaseCompleteCause5GSMTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

type nasTypeCause5GSMCauseValueData struct {
	in  uint8
	out uint8
}

var nasTypeCause5GSMOctetTable = []nasTypeCause5GSMCauseValueData{
	{0xff, 0xff},
}

func TestNasTypeCause5GSMGetSetCauseValue(t *testing.T) {
	a := nasType.NewCause5GSM(nasMessage.PDUSessionReleaseCompleteCause5GSMType)
	for _, table := range nasTypeCause5GSMOctetTable {
		a.SetCauseValue(table.in)
		assert.Equal(t, table.out, a.GetCauseValue())
	}
}

type testCause5GSMDataTemplate struct {
	in  nasType.Cause5GSM
	out nasType.Cause5GSM
}

var cause5GSMTestData = []nasType.Cause5GSM{
	{nasMessage.PDUSessionReleaseCompleteCause5GSMType, 0xff},
}

var cause5GSMExpectedTestData = []nasType.Cause5GSM{
	{nasMessage.PDUSessionReleaseCompleteCause5GSMType, 0xff},
}

var cause5GSMTestTable = []testCause5GSMDataTemplate{
	{cause5GSMTestData[0], cause5GSMExpectedTestData[0]},
}

func TestNasTypeCause5GSM(t *testing.T) {

	for i, table := range cause5GSMTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewCause5GSM(nasMessage.PDUSessionReleaseCompleteCause5GSMType)

		a.SetIei(table.in.GetIei())
		a.SetCauseValue(table.in.Octet)

		assert.Equalf(t, table.out.Iei, a.Iei, "in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		assert.Equalf(t, table.out.Octet, a.Octet, "in(%v): out %v, actual %x", table.in.Octet, table.out.Octet, a.Octet)

	}
}
