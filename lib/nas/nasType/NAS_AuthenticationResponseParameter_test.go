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

func TestNasTypeNewAuthenticationResponseParameter(t *testing.T) {
	a := nasType.NewAuthenticationResponseParameter(nasMessage.AuthenticationResponseAuthenticationResponseParameterType)
	assert.NotNil(t, a)

}

var nasTypeAuthenticationResponseAuthenticationResponseParameterTable = []NasTypeIeiData{
	{nasMessage.AuthenticationResponseAuthenticationResponseParameterType, nasMessage.AuthenticationResponseAuthenticationResponseParameterType},
}

func TestNasTypeAuthenticationResponseParameterGetSetIei(t *testing.T) {
	a := nasType.NewAuthenticationResponseParameter(nasMessage.AuthenticationResponseAuthenticationResponseParameterType)
	for _, table := range nasTypeAuthenticationResponseAuthenticationResponseParameterTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

var nasTypeAuthenticationResponseAuthenticationResponseParameterLenTable = []NasTypeLenuint8Data{
	{16, 16},
}

func TestNasTypeAuthenticationResponseParameterGetSetLen(t *testing.T) {
	a := nasType.NewAuthenticationResponseParameter(nasMessage.AuthenticationResponseAuthenticationResponseParameterType)
	for _, table := range nasTypeAuthenticationResponseAuthenticationResponseParameterLenTable {
		a.SetLen(table.in)
		assert.Equal(t, table.out, a.GetLen())
	}
}

type nasTypeAuthenticationResponseParameterOctetData struct {
	inLen uint8
	in    [16]uint8
	out   [16]uint8
}

var nasTypeAuthenticationResponseParameterOctetTable = []nasTypeAuthenticationResponseParameterOctetData{
	{16, [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

func TestNasTypeAuthenticationResponseParameterGetSetRES(t *testing.T) {
	a := nasType.NewAuthenticationResponseParameter(nasMessage.AuthenticationResponseAuthenticationResponseParameterType)
	for _, table := range nasTypeAuthenticationResponseParameterOctetTable {
		a.SetLen(table.inLen)
		a.SetRES(table.in)
		assert.Equal(t, table.out, a.GetRES())
	}
}

type testAuthenticationResponseParameterDataTemplate struct {
	in  nasType.AuthenticationResponseParameter
	out nasType.AuthenticationResponseParameter
}

var authenticationResponseParameterTestData = []nasType.AuthenticationResponseParameter{
	{nasMessage.AuthenticationResponseAuthenticationResponseParameterType, 16, [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

var authenticationResponseParameterExpectedTestData = []nasType.AuthenticationResponseParameter{
	{nasMessage.AuthenticationResponseAuthenticationResponseParameterType, 16, [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

var authenticationResponseParameterTestTable = []testAuthenticationResponseParameterDataTemplate{
	{authenticationResponseParameterTestData[0], authenticationResponseParameterExpectedTestData[0]},
}

func TestNasTypeAuthenticationResponseParameter(t *testing.T) {

	for i, table := range authenticationResponseParameterTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewAuthenticationResponseParameter(nasMessage.AuthenticationResponseAuthenticationResponseParameterType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetRES(table.in.Octet)

		assert.Equalf(t, table.out.Iei, a.Iei, "in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		assert.Equalf(t, table.out.Len, a.Len, "in(%v): out %v, actual %x", table.in.Len, table.out.Len, a.Len)
		assert.Equalf(t, table.out.Octet, a.Octet, "in(%v): out %v, actual %x", table.in.Octet, table.out.Octet, a.Octet)

	}
}
