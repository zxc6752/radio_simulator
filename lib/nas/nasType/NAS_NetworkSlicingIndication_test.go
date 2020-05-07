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

var ConfigurationUpdateCommandNetworkSlicingIndicationTypeIeiInput uint8 = 0x09

func TestNasTypeNewNetworkSlicingIndication(t *testing.T) {
	a := nasType.NewNetworkSlicingIndication(ConfigurationUpdateCommandNetworkSlicingIndicationTypeIeiInput)
	assert.NotNil(t, a)
}

var nasTypeConfigurationUpdateCommandNetworkSlicingIndicationTable = []NasTypeIeiData{
	{ConfigurationUpdateCommandNetworkSlicingIndicationTypeIeiInput, 0x09},
}

func TestNasTypeNetworkSlicingIndicationGetSetIei(t *testing.T) {
	a := nasType.NewNetworkSlicingIndication(ConfigurationUpdateCommandNetworkSlicingIndicationTypeIeiInput)
	for _, table := range nasTypeConfigurationUpdateCommandNetworkSlicingIndicationTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

type nasTypeNetworkSlicingIndication struct {
	inDCNI   uint8
	outDCNI  uint8
	inNSSCI  uint8
	outNSSCI uint8
	outIei   uint8
}

var nasTypeNetworkSlicingIndicationTable = []nasTypeNetworkSlicingIndication{
	{0x01, 0x01, 0x01, 0x01, 0x09},
}

func TestNasTypeNetworkSlicingIndication(t *testing.T) {
	a := nasType.NewNetworkSlicingIndication(ConfigurationUpdateCommandNetworkSlicingIndicationTypeIeiInput)
	for _, table := range nasTypeNetworkSlicingIndicationTable {
		a.SetDCNI(table.inDCNI)
		a.SetNSSCI(table.inNSSCI)

		assert.Equal(t, table.outIei, a.GetIei())
		assert.Equal(t, table.outDCNI, a.GetDCNI())
		assert.Equal(t, table.outNSSCI, a.GetNSSCI())
	}
}
