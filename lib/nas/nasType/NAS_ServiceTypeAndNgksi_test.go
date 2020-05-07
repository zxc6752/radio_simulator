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

type nasTypeNgksiAndServiceTypeData struct {
	inTsc                  uint8
	outTsc                 uint8
	inNASKeySetIdentifier  uint8
	outNASKeySetIdentifier uint8
	inServiceTypeValue     uint8
	outServiceTypeValue    uint8
}

var nasTypeNgksiAndServiceTypeTable = []nasTypeNgksiAndServiceTypeData{
	{0x01, 0x01, 0x07, 0x07, 0x7, 0x07},
}

func TestNasTypeNewServiceTypeAndNgksi(t *testing.T) {
	a := nasType.NewServiceTypeAndNgksi()
	assert.NotNil(t, a)
}

func TestNasTypeGetSetNgksiAndServiceType(t *testing.T) {
	a := nasType.NewServiceTypeAndNgksi()
	for _, table := range nasTypeNgksiAndServiceTypeTable {
		a.SetTSC(table.inTsc)
		assert.Equal(t, table.outTsc, a.GetTSC())
		// a.SetTSC(0)
		a.SetNasKeySetIdentifiler(table.inNASKeySetIdentifier)
		assert.Equal(t, table.outNASKeySetIdentifier, a.GetNasKeySetIdentifiler())

		a.SetServiceTypeValue(table.inServiceTypeValue)
		assert.Equal(t, table.outServiceTypeValue, a.GetServiceTypeValue())

	}
}
