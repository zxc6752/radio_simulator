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

func TestNasTypeNewAuthenticationParameterAUTN(t *testing.T) {
	a := nasType.NewAuthenticationParameterAUTN(nasMessage.AuthenticationRequestAuthenticationParameterAUTNType)
	assert.NotNil(t, a)

}

var nasTypeAuthenticationResultAuthenticationParameterAUTNTable = []NasTypeIeiData{
	{nasMessage.AuthenticationRequestAuthenticationParameterAUTNType, nasMessage.AuthenticationRequestAuthenticationParameterAUTNType},
}

func TestNasTypeAuthenticationParameterAUTNGetSetIei(t *testing.T) {
	a := nasType.NewAuthenticationParameterAUTN(nasMessage.AuthenticationRequestAuthenticationParameterAUTNType)
	for _, table := range nasTypeAuthenticationResultAuthenticationParameterAUTNTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

var nasTypeAuthenticationResultAuthenticationParameterAUTNLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeAuthenticationParameterAUTNGetSetLen(t *testing.T) {
	a := nasType.NewAuthenticationParameterAUTN(nasMessage.AuthenticationRequestAuthenticationParameterAUTNType)
	for _, table := range nasTypeAuthenticationResultAuthenticationParameterAUTNLenTable {
		a.SetLen(table.in)
		assert.Equal(t, table.out, a.GetLen())
	}
}

type nasTypeAuthenticationParameterAUTNOctetData struct {
	inLen uint8
	in    [16]uint8
	out   [16]uint8
}

var nasTypeAuthenticationParameterAUTNOctetTable = []nasTypeAuthenticationParameterAUTNOctetData{
	{16, [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

func TestNasTypeAuthenticationParameterAUTNGetSetAUTN(t *testing.T) {
	a := nasType.NewAuthenticationParameterAUTN(nasMessage.AuthenticationRequestAuthenticationParameterAUTNType)
	for _, table := range nasTypeAuthenticationParameterAUTNOctetTable {
		a.SetLen(table.inLen)
		a.SetAUTN(table.in)
		assert.Equal(t, table.out, a.GetAUTN())
	}
}

type testAuthenticationParameterAUTNDataTemplate struct {
	in  nasType.AuthenticationParameterAUTN
	out nasType.AuthenticationParameterAUTN
}

var authenticationParameterAUTNTestData = []nasType.AuthenticationParameterAUTN{
	{nasMessage.AuthenticationRequestAuthenticationParameterAUTNType, 16, [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

var authenticationParameterAUTNExpectedTestData = []nasType.AuthenticationParameterAUTN{
	{nasMessage.AuthenticationRequestAuthenticationParameterAUTNType, 16, [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

var authenticationParameterAUTNTestTable = []testAuthenticationParameterAUTNDataTemplate{
	{authenticationParameterAUTNTestData[0], authenticationParameterAUTNExpectedTestData[0]},
}

func TestNasTypeAuthenticationParameterAUTN(t *testing.T) {

	for i, table := range authenticationParameterAUTNTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewAuthenticationParameterAUTN(nasMessage.AuthenticationRequestAuthenticationParameterAUTNType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetAUTN(table.in.Octet)

		assert.Equalf(t, table.out.Iei, a.Iei, "in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		assert.Equalf(t, table.out.Len, a.Len, "in(%v): out %v, actual %x", table.in.Len, table.out.Len, a.Len)
		assert.Equalf(t, table.out.Octet, a.Octet, "in(%v): out %v, actual %x", table.in.Octet, table.out.Octet, a.Octet)

	}
}
