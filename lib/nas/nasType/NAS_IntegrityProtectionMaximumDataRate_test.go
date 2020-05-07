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

func TestNasTypeNewIntegrityProtectionMaximumDataRate(t *testing.T) {
	a := nasType.NewIntegrityProtectionMaximumDataRate(nasMessage.PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType)
	assert.NotNil(t, a)

}

var nasTypePDUSessionModificationRequestIntegrityProtectionMaximumDataRateTable = []NasTypeIeiData{
	{nasMessage.PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType, nasMessage.PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType},
}

func TestNasTypeIntegrityProtectionMaximumDataRateGetSetIei(t *testing.T) {
	a := nasType.NewIntegrityProtectionMaximumDataRate(nasMessage.PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType)
	for _, table := range nasTypePDUSessionModificationRequestIntegrityProtectionMaximumDataRateTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

type nasTypeIntegrityProtectionMaximumDataRateMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLinkData struct {
	in  uint8
	out uint8
}

var nasTypeIntegrityProtectionMaximumDataRateMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLinkTable = []nasTypeIntegrityProtectionMaximumDataRateMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLinkData{
	{0xff, 0xff},
}

func TestNasTypeIntegrityProtectionMaximumDataRateGetSetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink(t *testing.T) {
	a := nasType.NewIntegrityProtectionMaximumDataRate(nasMessage.PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType)
	for _, table := range nasTypeIntegrityProtectionMaximumDataRateMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLinkTable {
		a.SetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink(table.in)
		assert.Equal(t, table.out, a.GetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink())
	}
}

type nasTypeIntegrityProtectionMaximumDataRateMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLinkData struct {
	in  uint8
	out uint8
}

var nasTypeIntegrityProtectionMaximumDataRateMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLinkTable = []nasTypeIntegrityProtectionMaximumDataRateMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLinkData{
	{0xff, 0xff},
}

func TestNasTypeIntegrityProtectionMaximumDataRateGetSetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink(t *testing.T) {
	a := nasType.NewIntegrityProtectionMaximumDataRate(nasMessage.PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType)
	for _, table := range nasTypeIntegrityProtectionMaximumDataRateMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLinkTable {
		a.SetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink(table.in)
		assert.Equal(t, table.out, a.GetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink())
	}
}

type testIntegrityProtectionMaximumDataRateDataTemplate struct {
	inIei                                                             uint8
	inMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink    uint8
	inMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink  uint8
	outIei                                                            uint8
	outMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink   uint8
	outMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink uint8
}

var integrityProtectionMaximumDataRateTestTable = []testIntegrityProtectionMaximumDataRateDataTemplate{
	{nasMessage.PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType, 0xff, 0x11,
		nasMessage.PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType, 0xff, 0x11},
}

func TestNasTypeIntegrityProtectionMaximumDataRate(t *testing.T) {

	for i, table := range integrityProtectionMaximumDataRateTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewIntegrityProtectionMaximumDataRate(nasMessage.PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType)

		a.SetIei(table.inIei)
		a.SetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink(table.inMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink)
		a.SetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink(table.inMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink)

		assert.Equalf(t, table.outIei, a.Iei, "in(%v): out %v, actual %x", table.inIei, table.outIei, a.Iei)
		assert.Equalf(t, table.outMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink, a.GetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink(), "in(%v): out %v, actual %x", table.inMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink, table.outMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink, a.GetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForUpLink())
		assert.Equalf(t, table.outMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink, a.GetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink(), "in(%v): out %v, actual %x", table.inMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink, table.outMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink, a.GetMaximumDataRatePerUEForUserPlaneIntegrityProtectionForDownLink())

	}
}
