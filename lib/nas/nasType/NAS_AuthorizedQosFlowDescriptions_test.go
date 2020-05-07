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

func TestNasTypeNewAuthorizedQosFlowDescriptions(t *testing.T) {
	a := nasType.NewAuthorizedQosFlowDescriptions(nasMessage.PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType)
	assert.NotNil(t, a)
}

var nasTypePDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsTable = []NasTypeIeiData{
	{nasMessage.PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType, nasMessage.PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType},
}

func TestNasTypeAuthorizedQosFlowDescriptionsGetSetIei(t *testing.T) {
	a := nasType.NewAuthorizedQosFlowDescriptions(nasMessage.PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType)
	for _, table := range nasTypePDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

var nasTypePDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsLenTable = []NasTypeLenUint16Data{
	{12, 12},
}

func TestNasTypeAuthorizedQosFlowDescriptionsGetSetLen(t *testing.T) {
	a := nasType.NewAuthorizedQosFlowDescriptions(nasMessage.PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType)
	for _, table := range nasTypePDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsLenTable {
		a.SetLen(table.in)
		assert.Equal(t, table.out, a.GetLen())
	}
}

type nasTypeQoSFlowDescription struct {
	inLen uint16
	in    []uint8
	out   []uint8
}

var nasTypeQoSFlowDescriptionTable = []nasTypeQoSFlowDescription{
	{2, []uint8{0x00, 0x01}, []uint8{0x00, 0x01}},
}

func TestNasTypeAuthorizedQosFlowDescriptionsGetSetQoSFlowDescription(t *testing.T) {
	a := nasType.NewAuthorizedQosFlowDescriptions(nasMessage.PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType)
	for _, table := range nasTypeQoSFlowDescriptionTable {
		a.SetLen(table.inLen)
		a.SetQoSFlowDescriptions(table.in)
		assert.Equalf(t, table.out, a.GetQoSFlowDescriptions(), "in(%v): out %v, actual %x", table.in, table.out, a.GetQoSFlowDescriptions())
	}
}

type testAuthorizedQosFlowDescriptionsDataTemplate struct {
	in  nasType.AuthorizedQosFlowDescriptions
	out nasType.AuthorizedQosFlowDescriptions
}

var AuthorizedQosFlowDescriptionsTestData = []nasType.AuthorizedQosFlowDescriptions{
	{nasMessage.PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType, 2, []uint8{0x00, 0x01}},
}

var AuthorizedQosFlowDescriptionsExpectedTestData = []nasType.AuthorizedQosFlowDescriptions{
	{nasMessage.PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType, 2, []uint8{0x00, 0x01}},
}

var AuthorizedQosFlowDescriptionsTable = []testAuthorizedQosFlowDescriptionsDataTemplate{
	{AuthorizedQosFlowDescriptionsTestData[0], AuthorizedQosFlowDescriptionsExpectedTestData[0]},
}

func TestNasTypeAuthorizedQosFlowDescriptions(t *testing.T) {
	for i, table := range AuthorizedQosFlowDescriptionsTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewAuthorizedQosFlowDescriptions(nasMessage.PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetQoSFlowDescriptions(table.in.Buffer)

		assert.Equalf(t, table.out.Iei, a.Iei, "in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		assert.Equalf(t, table.out.Len, a.Len, "in(%v): out %v, actual %x", table.in.Len, table.out.Len, a.Len)
		assert.Equalf(t, table.out.Buffer, a.Buffer, "in(%v): out %v, actual %x", table.in.Buffer, table.out.Buffer, a.Buffer)
	}

}
