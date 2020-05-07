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

var RegistrationAcceptNSSAIInclusionModeTypeIeiInput uint8 = 0x0A

func TestNasTypeNewNSSAIInclusionMode(t *testing.T) {
	a := nasType.NewNSSAIInclusionMode(RegistrationAcceptNSSAIInclusionModeTypeIeiInput)
	assert.NotNil(t, a)
}

var nasTypeNSSAIInclusionModeRegistrationAcceptNSSAIInclusionModeTypeTable = []NasTypeIeiData{
	{RegistrationAcceptNSSAIInclusionModeTypeIeiInput, 0x0A},
}

func TestNasTypeNSSAIInclusionModeGetSetIei(t *testing.T) {
	a := nasType.NewNSSAIInclusionMode(RegistrationAcceptNSSAIInclusionModeTypeIeiInput)
	for _, table := range nasTypeNSSAIInclusionModeRegistrationAcceptNSSAIInclusionModeTypeTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

type nasTypeNSSAIInclusionMode struct {
	inIei                 uint8
	inNSSAIInclusionMode  uint8
	outIei                uint8
	outNSSAIInclusionMode uint8
}

var nasTypeNSSAIInclusionModeTable = []nasTypeNSSAIInclusionMode{
	{RegistrationAcceptNSSAIInclusionModeTypeIeiInput, 0x03,
		0x0A, 0x03},
}

func TestNasTypeNSSAIInclusionMode(t *testing.T) {
	a := nasType.NewNSSAIInclusionMode(RegistrationAcceptNSSAIInclusionModeTypeIeiInput)
	for _, table := range nasTypeNSSAIInclusionModeTable {

		a.SetNSSAIInclusionMode(table.inNSSAIInclusionMode)

		assert.Equal(t, table.outIei, a.GetIei())
		assert.Equal(t, table.outNSSAIInclusionMode, a.GetNSSAIInclusionMode())
	}
}
