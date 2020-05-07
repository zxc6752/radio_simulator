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

func TestNasTypeNewRejectedNSSAI(t *testing.T) {
	a := nasType.NewRejectedNSSAI(nasMessage.RegistrationAcceptRejectedNSSAIType)
	assert.NotNil(t, a)

}

var nasTypeAuthenticationResultRejectedNSSAITable = []NasTypeIeiData{
	{nasMessage.RegistrationAcceptRejectedNSSAIType, nasMessage.RegistrationAcceptRejectedNSSAIType},
}

func TestNasTypeRejectedNSSAIGetSetIei(t *testing.T) {
	a := nasType.NewRejectedNSSAI(nasMessage.RegistrationAcceptRejectedNSSAIType)
	for _, table := range nasTypeAuthenticationResultRejectedNSSAITable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())

		// if a.GetIei() != table.out {
		// 	t.Errorf("in(%d): out %d, actual %d", table.in, table.out, a.GetIei())
		// }
	}
}

var nasTypeAuthenticationResultRejectedNSSAILenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeRejectedNSSAIGetSetLen(t *testing.T) {
	a := nasType.NewRejectedNSSAI(nasMessage.RegistrationAcceptRejectedNSSAIType)
	for _, table := range nasTypeAuthenticationResultRejectedNSSAILenTable {
		a.SetLen(table.in)
		assert.Equal(t, table.out, a.GetLen())
		// if a.GetLen() != table.out {
		// 	t.Errorf("in(%d): out %d, actual %d", table.in, table.out, a.GetLen())
		// }
	}
}

type nasTypeRejectedNSSAIContentsData struct {
	inLen uint8
	in    []uint8
	out   []uint8
}

var nasTypeRejectedNSSAIContentsTable = []nasTypeRejectedNSSAIContentsData{
	{2, []uint8{0x00, 0x01}, []uint8{0x00, 0x1}},
}

func TestNasTypeRejectedNSSAIGetSetRejectedNSSAIContents(t *testing.T) {
	a := nasType.NewRejectedNSSAI(nasMessage.RegistrationAcceptRejectedNSSAIType)
	for _, table := range nasTypeRejectedNSSAIContentsTable {
		a.SetLen(table.inLen)
		a.SetRejectedNSSAIContents(table.in)
		assert.Equalf(t, table.out, a.GetRejectedNSSAIContents(), "in(%v): out %v, actual %x", table.in, table.out, a.GetRejectedNSSAIContents())
	}
}

type testRejectedNSSAIDataTemplate struct {
	in  nasType.RejectedNSSAI
	out nasType.RejectedNSSAI
}

var RejectedNSSAITestData = []nasType.RejectedNSSAI{
	{nasMessage.RegistrationAcceptRejectedNSSAIType, 2, []byte{0x00, 0x00}},
}

var RejectedNSSAIExpectedTestData = []nasType.RejectedNSSAI{
	{nasMessage.RegistrationAcceptRejectedNSSAIType, 2, []byte{0x00, 0x00}},
}

var RejectedNSSAITestTable = []testRejectedNSSAIDataTemplate{
	{RejectedNSSAITestData[0], RejectedNSSAIExpectedTestData[0]},
}

func TestNasTypeRejectedNSSAI(t *testing.T) {

	for i, table := range RejectedNSSAITestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewRejectedNSSAI(nasMessage.RegistrationAcceptRejectedNSSAIType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetRejectedNSSAIContents(table.in.Buffer)

		assert.Equalf(t, table.out.Iei, a.Iei, "in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		assert.Equalf(t, table.out.Len, a.Len, "in(%v): out %v, actual %x", table.in.Len, table.out.Len, a.Len)
		assert.Equalf(t, table.out.Buffer, a.Buffer, "in(%v): out %v, actual %x", table.in.Buffer, table.out.Buffer, a.Buffer)

	}
}
