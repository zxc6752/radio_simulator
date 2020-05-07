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

func TestNasTypeNewSpareHalfOctetAndAccessType(t *testing.T) {
	a := nasType.NewSpareHalfOctetAndAccessType()
	assert.NotNil(t, a)
}

type nasTypeAccessType struct {
	in  uint8
	out uint8
}

var nasTypeAccessTypeTable = []nasTypeAccessType{
	{0x03, 0x03},
}

func TestNasTypeGetSetAccessType(t *testing.T) {
	a := nasType.NewSpareHalfOctetAndAccessType()
	for _, table := range nasTypeAccessTypeTable {
		a.SetAccessType(table.in)
		assert.Equal(t, table.out, a.GetAccessType())
	}
}

type AccessTypeAndSpareHalfOctetTestDataTemplate struct {
	in  nasType.SpareHalfOctetAndAccessType
	out nasType.SpareHalfOctetAndAccessType
}

var accessTypeAndSpareHalfOctetTestData = []nasType.SpareHalfOctetAndAccessType{
	{0x03},
}

var accessTypeAndSpareHalfOctetExpectedTestData = []nasType.SpareHalfOctetAndAccessType{
	{0x03},
}

var accessTypeAndSpareHalfOctetTable = []AccessTypeAndSpareHalfOctetTestDataTemplate{
	{accessTypeAndSpareHalfOctetTestData[0], accessTypeAndSpareHalfOctetExpectedTestData[0]},
}

func TestNasTypeAccessTypeAndSpareHalfOctet(t *testing.T) {

	for i, table := range accessTypeAndSpareHalfOctetTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewSpareHalfOctetAndAccessType()

		a.SetAccessType(table.in.GetAccessType())
		assert.Equal(t, table.out.GetAccessType(), a.GetAccessType())
	}
}
