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

func TestNasTypeNewAuthenticationFailureParameter(t *testing.T) {
	a := nasType.NewAuthenticationFailureParameter(nasMessage.AuthenticationFailureAuthenticationFailureParameterType)
	assert.NotNil(t, a)

}

var nasTypeAuthenticationResultAuthenticationFailureParameterTable = []NasTypeIeiData{
	{nasMessage.AuthenticationFailureAuthenticationFailureParameterType, nasMessage.AuthenticationFailureAuthenticationFailureParameterType},
}

func TestNasTypeAuthenticationFailureParameterGetSetIei(t *testing.T) {
	a := nasType.NewAuthenticationFailureParameter(nasMessage.AuthenticationFailureAuthenticationFailureParameterType)
	for _, table := range nasTypeAuthenticationResultAuthenticationFailureParameterTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

var nasTypeAuthenticationResultAuthenticationFailureParameterLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeAuthenticationFailureParameterGetSetLen(t *testing.T) {
	a := nasType.NewAuthenticationFailureParameter(nasMessage.AuthenticationFailureAuthenticationFailureParameterType)
	for _, table := range nasTypeAuthenticationResultAuthenticationFailureParameterLenTable {
		a.SetLen(table.in)
		assert.Equal(t, table.out, a.GetLen())
	}
}

type nasTypeOctetData struct {
	inLen uint8
	in    [14]uint8
	out   [14]uint8
}

var nasTypeOctetTable = []nasTypeOctetData{
	{14, [14]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, [14]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

func TestNasTypeAuthenticationFailureParameterGetSetOctet(t *testing.T) {
	a := nasType.NewAuthenticationFailureParameter(nasMessage.AuthenticationFailureAuthenticationFailureParameterType)
	for _, table := range nasTypeOctetTable {
		a.SetLen(table.inLen)
		a.SetAuthenticationFailureParameter(table.in)
		assert.Equal(t, table.out, a.GetAuthenticationFailureParameter())
	}
}

type testAuthenticationFailureParameterDataTemplate struct {
	in  nasType.AuthenticationFailureParameter
	out nasType.AuthenticationFailureParameter
}

var authenticationFailureParameterTestData = []nasType.AuthenticationFailureParameter{
	{nasMessage.AuthenticationFailureAuthenticationFailureParameterType, 14, [14]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

var authenticationFailureParameterExpectedTestData = []nasType.AuthenticationFailureParameter{
	{nasMessage.AuthenticationFailureAuthenticationFailureParameterType, 14, [14]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

var authenticationFailureParameterTestTable = []testAuthenticationFailureParameterDataTemplate{
	{authenticationFailureParameterTestData[0], authenticationFailureParameterExpectedTestData[0]},
}

func TestNasTypeAuthenticationFailureParameter(t *testing.T) {

	for i, table := range authenticationFailureParameterTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewAuthenticationFailureParameter(nasMessage.AuthenticationFailureAuthenticationFailureParameterType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetAuthenticationFailureParameter(table.in.Octet)

		assert.Equalf(t, table.out.Iei, a.Iei, "in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		assert.Equalf(t, table.out.Len, a.Len, "in(%v): out %v, actual %x", table.in.Len, table.out.Len, a.Len)
		assert.Equalf(t, table.out.Octet, a.Octet, "in(%v): out %v, actual %x", table.in.Octet, table.out.Octet, a.Octet)

	}
}
