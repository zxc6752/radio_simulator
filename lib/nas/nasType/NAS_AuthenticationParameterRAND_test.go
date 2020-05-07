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

func TestNasTypeNewAuthenticationParameterRAND(t *testing.T) {
	a := nasType.NewAuthenticationParameterRAND(nasMessage.AuthenticationRequestAuthenticationParameterRANDType)
	assert.NotNil(t, a)

}

var nasTypeAuthenticationRequestAuthenticationParameterRANDTable = []NasTypeIeiData{
	{nasMessage.AuthenticationRequestAuthenticationParameterRANDType, nasMessage.AuthenticationRequestAuthenticationParameterRANDType},
}

func TestNasTypeAuthenticationParameterRANDGetSetIei(t *testing.T) {
	a := nasType.NewAuthenticationParameterRAND(nasMessage.AuthenticationRequestAuthenticationParameterRANDType)
	for _, table := range nasTypeAuthenticationRequestAuthenticationParameterRANDTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

type nasTypeAuthenticationParameterRANDOctetData struct {
	in  [16]uint8
	out [16]uint8
}

var nasTypeAuthenticationParameterRANDOctetTable = []nasTypeAuthenticationParameterRANDOctetData{
	{[16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

func TestNasTypeAuthenticationParameterRANDGetSetRANDValue(t *testing.T) {
	a := nasType.NewAuthenticationParameterRAND(nasMessage.AuthenticationRequestAuthenticationParameterRANDType)
	for _, table := range nasTypeAuthenticationParameterRANDOctetTable {
		a.SetRANDValue(table.in)
		assert.Equal(t, table.out, a.GetRANDValue())
	}
}

type testAuthenticationParameterRANDDataTemplate struct {
	in  nasType.AuthenticationParameterRAND
	out nasType.AuthenticationParameterRAND
}

var authenticationParameterRANDTestData = []nasType.AuthenticationParameterRAND{
	{nasMessage.AuthenticationRequestAuthenticationParameterRANDType, [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

var authenticationParameterRANDExpectedTestData = []nasType.AuthenticationParameterRAND{
	{nasMessage.AuthenticationRequestAuthenticationParameterRANDType, [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
}

var authenticationParameterRANDTestTable = []testAuthenticationParameterRANDDataTemplate{
	{authenticationParameterRANDTestData[0], authenticationParameterRANDExpectedTestData[0]},
}

func TestNasTypeAuthenticationParameterRAND(t *testing.T) {

	for i, table := range authenticationParameterRANDTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewAuthenticationParameterRAND(nasMessage.AuthenticationRequestAuthenticationParameterRANDType)

		a.SetIei(table.in.GetIei())
		a.SetRANDValue(table.in.Octet)

		assert.Equalf(t, table.out.Iei, a.Iei, "in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		assert.Equalf(t, table.out.Octet, a.Octet, "in(%v): out %v, actual %x", table.in.Octet, table.out.Octet, a.Octet)

	}
}
