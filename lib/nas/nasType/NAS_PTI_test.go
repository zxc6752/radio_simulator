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

func TestNasTypeNewPTI(t *testing.T) {
	a := nasType.NewPTI()
	assert.NotNil(t, a)
}

type nasTypePTI struct {
	in  uint8
	out uint8
}

var nasTypePTITable = []nasTypePTI{
	{0x03, 0x03},
}

func TestNasTypePTIGetSetPDUSessionIdentity(t *testing.T) {
	a := nasType.NewPTI()
	for _, table := range nasTypePTITable {
		a.SetPTI(table.in)
		assert.Equal(t, table.out, a.GetPTI())
	}
}

type PTITestDataTemplate struct {
	in  nasType.PTI
	out nasType.PTI
}

var PTITestData = []nasType.PTI{
	{0x03},
}

var PTIExpectedTestData = []nasType.PTI{
	{0x03},
}

var PTITable = []PTITestDataTemplate{
	{PTITestData[0], PTIExpectedTestData[0]},
}

func TestNasTypePTI(t *testing.T) {

	for _, table := range PTITable {

		a := nasType.NewPTI()

		a.SetPTI(table.in.GetPTI())
		assert.Equal(t, table.out.GetPTI(), a.GetPTI())
	}
}
