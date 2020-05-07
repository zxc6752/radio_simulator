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
	"radio_simulator/lib/nas/nasType"
	"testing"

	"github.com/stretchr/testify/assert"
)

var RegistrationRequestNoncurrentNativeNASKeySetIdentifierTypeIeiInput uint8 = 0x0C

func TestNasTypeNewNoncurrentNativeNASKeySetIdentifier(t *testing.T) {
	a := nasType.NewNoncurrentNativeNASKeySetIdentifier(RegistrationRequestNoncurrentNativeNASKeySetIdentifierTypeIeiInput)
	assert.NotNil(t, a)
}

var nasTypeConfigurationUpdateCommandNoncurrentNativeNASKeySetIdentifierTable = []NasTypeIeiData{
	{RegistrationRequestNoncurrentNativeNASKeySetIdentifierTypeIeiInput, 0x0C},
}

func TestNasTypeNoncurrentNativeNASKeySetIdentifierGetSetIei(t *testing.T) {
	a := nasType.NewNoncurrentNativeNASKeySetIdentifier(RegistrationRequestNoncurrentNativeNASKeySetIdentifierTypeIeiInput)
	for _, table := range nasTypeConfigurationUpdateCommandNoncurrentNativeNASKeySetIdentifierTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

type nasTypeNoncurrentNativeNASKeySetIdentifier struct {
	inIei                   uint8
	inTsc                   uint8
	inNasKeySetIdentifiler  uint8
	outIei                  uint8
	outTsc                  uint8
	outNasKeySetIdentifiler uint8
}

var nasTypeNoncurrentNativeNASKeySetIdentifierTable = []nasTypeNoncurrentNativeNASKeySetIdentifier{
	{RegistrationRequestNoncurrentNativeNASKeySetIdentifierTypeIeiInput, 0x01, 0x01,
		0x0C, 0x01, 0x01},
}

func TestNasTypeNoncurrentNativeNASKeySetIdentifier(t *testing.T) {
	a := nasType.NewNoncurrentNativeNASKeySetIdentifier(RegistrationRequestNoncurrentNativeNASKeySetIdentifierTypeIeiInput)
	for _, table := range nasTypeNoncurrentNativeNASKeySetIdentifierTable {
		a.SetTsc(table.inTsc)
		a.SetNasKeySetIdentifiler(table.inNasKeySetIdentifiler)

		assert.Equal(t, table.outIei, a.GetIei())
		assert.Equal(t, table.outTsc, a.GetTsc())
		assert.Equal(t, table.outNasKeySetIdentifiler, a.GetNasKeySetIdentifiler())
	}
}
