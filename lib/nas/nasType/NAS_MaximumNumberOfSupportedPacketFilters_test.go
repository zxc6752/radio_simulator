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

func TestNasTypeNewMaximumNumberOfSupportedPacketFilters(t *testing.T) {
	a := nasType.NewMaximumNumberOfSupportedPacketFilters(nasMessage.PDUSessionModificationRequestMaximumNumberOfSupportedPacketFiltersType)
	assert.NotNil(t, a)

}

var nasTypePDUSessionModificationRequestMaximumNumberOfSupportedPacketFiltersTable = []NasTypeIeiData{
	{nasMessage.PDUSessionModificationRequestMaximumNumberOfSupportedPacketFiltersType, nasMessage.PDUSessionModificationRequestMaximumNumberOfSupportedPacketFiltersType},
}

func TestNasTypeMaximumNumberOfSupportedPacketFiltersGetSetIei(t *testing.T) {
	a := nasType.NewMaximumNumberOfSupportedPacketFilters(nasMessage.PDUSessionModificationRequestMaximumNumberOfSupportedPacketFiltersType)
	for _, table := range nasTypePDUSessionModificationRequestMaximumNumberOfSupportedPacketFiltersTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

type nasTypeMaximumNumberOfSupportedPacketFilters struct {
	in  uint16
	out uint16
}

var nasTypeMaximumNumberOfSupportedPacketFiltersTable = []nasTypeMaximumNumberOfSupportedPacketFilters{
	{0x0100, 0x0100},
}

func TestNasTypeMaximumNumberOfSupportedPacketFiltersGetSetMaximumNumberOfSupportedPacketFilters(t *testing.T) {
	a := nasType.NewMaximumNumberOfSupportedPacketFilters(nasMessage.PDUSessionModificationRequestMaximumNumberOfSupportedPacketFiltersType)
	for _, table := range nasTypeMaximumNumberOfSupportedPacketFiltersTable {
		a.SetMaximumNumberOfSupportedPacketFilters(table.in)
		assert.Equal(t, table.out, a.GetMaximumNumberOfSupportedPacketFilters())
	}
}
