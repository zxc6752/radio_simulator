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

type nasTypeNewNgksiAndSpareHalfOctetData struct {
	inTsc                  uint8
	outTsc                 uint8
	inNASKeySetIdentifier  uint8
	outNASKeySetIdentifier uint8
	inSpareHalfOctet       uint8
	outSpareHalfOctet      uint8
}

var nasTypeNewNgksiAndSpareHalfOctetTable = []nasTypeNewNgksiAndSpareHalfOctetData{
	{0x1, 0x1, 0x7, 0x7, 0x7, 0x7},
}

func TestNasTypeNewSpareHalfOctetAndNgksi(t *testing.T) {
	a := nasType.NewSpareHalfOctetAndNgksi()
	assert.NotNil(t, a)
}

func TestNasTypeGetSetSpareHalfOctetAndNgksi(t *testing.T) {
	a := nasType.NewSpareHalfOctetAndNgksi()
	for _, table := range nasTypeNewNgksiAndSpareHalfOctetTable {
		a.SetTSC(table.inTsc)
		assert.Equal(t, table.outTsc, a.GetTSC())
		a.SetNasKeySetIdentifiler(table.inNASKeySetIdentifier)
		assert.Equal(t, table.outNASKeySetIdentifier, a.GetNasKeySetIdentifiler())

		a.SetSpareHalfOctet(table.inSpareHalfOctet)
		assert.Equal(t, table.outSpareHalfOctet, a.GetSpareHalfOctet())

	}
}
