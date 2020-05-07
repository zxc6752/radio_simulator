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

func TestNasTypeNewSequenceNumber(t *testing.T) {
	a := nasType.NewSequenceNumber()
	assert.NotNil(t, a)
}

type nasTypeSequenceNumber struct {
	in  uint8
	out uint8
}

var nasTypeSequenceNumberTable = []nasTypeSequenceNumber{
	{0x03, 0x03},
}

func TestNasTypeSequenceNumberGetSetSQN(t *testing.T) {
	a := nasType.NewSequenceNumber()
	for _, table := range nasTypeSequenceNumberTable {
		a.SetSQN(table.in)
		assert.Equal(t, table.out, a.GetSQN())
	}
}

type SequenceNumberTestDataTemplate struct {
	in  nasType.SequenceNumber
	out nasType.SequenceNumber
}

var SequenceNumberTestData = []nasType.SequenceNumber{
	{0x03},
}

var SequenceNumberExpectedTestData = []nasType.SequenceNumber{
	{0x03},
}

var SequenceNumberTable = []SequenceNumberTestDataTemplate{
	{SequenceNumberTestData[0], SequenceNumberExpectedTestData[0]},
}

func TestNasTypeSequenceNumber(t *testing.T) {

	for _, table := range SequenceNumberTable {

		a := nasType.NewSequenceNumber()

		a.SetSQN(table.in.GetSQN())
		assert.Equal(t, table.out.GetSQN(), a.GetSQN())
	}
}
