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

func TestNasTypeNewSMPDUDNRequestContainer(t *testing.T) {
	a := nasType.NewSMPDUDNRequestContainer(nasMessage.PDUSessionEstablishmentRequestSMPDUDNRequestContainerType)
	assert.NotNil(t, a)

}

var nasTypeSMPDUDNRequestContainerTable = []NasTypeIeiData{
	{nasMessage.PDUSessionEstablishmentRequestSMPDUDNRequestContainerType, nasMessage.PDUSessionEstablishmentRequestSMPDUDNRequestContainerType},
}

func TestNasTypeSMPDUDNRequestContainerGetSetIei(t *testing.T) {
	a := nasType.NewSMPDUDNRequestContainer(nasMessage.PDUSessionEstablishmentRequestSMPDUDNRequestContainerType)
	for _, table := range nasTypeSMPDUDNRequestContainerTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

var nasTypeSMPDUDNRequestContainerLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeSMPDUDNRequestContainerGetSetLen(t *testing.T) {
	a := nasType.NewSMPDUDNRequestContainer(nasMessage.PDUSessionEstablishmentRequestSMPDUDNRequestContainerType)
	for _, table := range nasTypeSMPDUDNRequestContainerLenTable {
		a.SetLen(table.in)
		assert.Equal(t, table.out, a.GetLen())
	}
}

type nasTypeSMPDUDNRequestContainerDNSpecificIdentityData struct {
	inLen uint8
	in    []uint8
	out   []uint8
}

var nasTypeSMPDUDNRequestContainerDNSpecificIdentityTable = []nasTypeSMPDUDNRequestContainerDNSpecificIdentityData{
	{2, []uint8{0x01, 0x01}, []uint8{0x01, 0x01}},
}

func TestNasTypeSMPDUDNRequestContainerGetSetDNSpecificIdentity(t *testing.T) {
	a := nasType.NewSMPDUDNRequestContainer(nasMessage.PDUSessionEstablishmentRequestSMPDUDNRequestContainerType)
	for _, table := range nasTypeSMPDUDNRequestContainerDNSpecificIdentityTable {
		a.SetLen(table.inLen) // fix it, set input length
		a.SetDNSpecificIdentity(table.in)
		assert.Equalf(t, table.out, a.GetDNSpecificIdentity(), "in(%v): out %v, actual %x", table.in, table.out, a.GetDNSpecificIdentity())
	}
}

type testSMPDUDNRequestContainerDataTemplate struct {
	in  nasType.SMPDUDNRequestContainer
	out nasType.SMPDUDNRequestContainer
}

var SMPDUDNRequestContainerTestData = []nasType.SMPDUDNRequestContainer{
	{nasMessage.PDUSessionEstablishmentRequestSMPDUDNRequestContainerType, 2, []uint8{}},
}

var SMPDUDNRequestContainerExpectedTestData = []nasType.SMPDUDNRequestContainer{
	{nasMessage.PDUSessionEstablishmentRequestSMPDUDNRequestContainerType, 2, []uint8{0x01, 0x01}},
}

var SMPDUDNRequestContainerTestTable = []testSMPDUDNRequestContainerDataTemplate{
	{SMPDUDNRequestContainerTestData[0], SMPDUDNRequestContainerExpectedTestData[0]},
}

func TestNasTypeSMPDUDNRequestContainer(t *testing.T) {

	for i, table := range SMPDUDNRequestContainerTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewSMPDUDNRequestContainer(nasMessage.PDUSessionEstablishmentRequestSMPDUDNRequestContainerType)

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetDNSpecificIdentity([]uint8{0x01, 0x01})

		assert.Equalf(t, table.out.Iei, a.Iei, "in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		assert.Equalf(t, table.out.Len, a.Len, "in(%v): out %v, actual %x", table.in.Len, table.out.Len, a.Len)
		assert.Equalf(t, table.out.Buffer, a.Buffer, "in(%v): out %v, actual %x", table.in.Buffer, table.out.Buffer, a.Buffer)

	}
}
