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

func TestNasTypeNewSTATUSMessageIdentity5GMM(t *testing.T) {
	a := nasType.NewSTATUSMessageIdentity5GMM()
	assert.NotNil(t, a)
}

type nasTypeSTATUSMessageIdentity5GMM struct {
	in  uint8
	out uint8
}

var nasTypeSTATUSMessageIdentity5GMMTable = []nasTypeSTATUSMessageIdentity5GMM{
	{0x03, 0x03},
}

func TestNasTypeSTATUSMessageIdentity5GMMGetSetMessageType(t *testing.T) {
	a := nasType.NewSTATUSMessageIdentity5GMM()
	for _, table := range nasTypeSTATUSMessageIdentity5GMMTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}

type STATUSMessageIdentity5GMMTestDataTemplate struct {
	in  nasType.STATUSMessageIdentity5GMM
	out nasType.STATUSMessageIdentity5GMM
}

var STATUSMessageIdentity5GMMTestData = []nasType.STATUSMessageIdentity5GMM{
	{0x03},
}

var STATUSMessageIdentity5GMMExpectedTestData = []nasType.STATUSMessageIdentity5GMM{
	{0x03},
}

var STATUSMessageIdentity5GMMTable = []STATUSMessageIdentity5GMMTestDataTemplate{
	{STATUSMessageIdentity5GMMTestData[0], STATUSMessageIdentity5GMMExpectedTestData[0]},
}

func TestNasTypeSTATUSMessageIdentity5GMM(t *testing.T) {

	for _, table := range STATUSMessageIdentity5GMMTable {

		a := nasType.NewSTATUSMessageIdentity5GMM()

		a.SetMessageType(table.in.GetMessageType())
		assert.Equal(t, table.out.GetMessageType(), a.GetMessageType())
	}
}
