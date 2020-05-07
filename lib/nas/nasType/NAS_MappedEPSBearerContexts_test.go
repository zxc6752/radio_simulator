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

func TestNasTypeNewMappedEPSBearerContexts(t *testing.T) {
	a := nasType.NewMappedEPSBearerContexts(nasMessage.PDUSessionModificationRequestMappedEPSBearerContextsType)
	assert.NotNil(t, a)

}

var nasTypeRegistrationRequestMappedEPSBearerContextsTable = []NasTypeIeiData{
	{nasMessage.PDUSessionModificationRequestMappedEPSBearerContextsType, nasMessage.PDUSessionModificationRequestMappedEPSBearerContextsType},
}

func TestNasTypeMappedEPSBearerContextsGetSetIei(t *testing.T) {
	a := nasType.NewMappedEPSBearerContexts(nasMessage.PDUSessionModificationRequestMappedEPSBearerContextsType)
	for _, table := range nasTypeRegistrationRequestMappedEPSBearerContextsTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

var nasTypeMappedEPSBearerContextsLenTable = []NasTypeLenUint16Data{
	{2, 2},
}

func TestNasTypeMappedEPSBearerContextsGetSetLen(t *testing.T) {
	a := nasType.NewMappedEPSBearerContexts(nasMessage.PDUSessionModificationRequestMappedEPSBearerContextsType)
	for _, table := range nasTypeMappedEPSBearerContextsLenTable {
		a.SetLen(table.in)
		assert.Equal(t, table.out, a.GetLen())
	}
}

type nasTypeMappedEPSBearerContextsMappedEPSBearerContextData struct {
	inLen uint16
	in    []uint8
	out   []uint8
}

var nasTypeMappedEPSBearerContextsMappedEPSBearerContextTable = []nasTypeMappedEPSBearerContextsMappedEPSBearerContextData{
	{2, []uint8{0xff, 0xff}, []uint8{0xff, 0xff}},
}

func TestNasTypeMappedEPSBearerContextsGetSetMappedEPSBearerContext(t *testing.T) {
	a := nasType.NewMappedEPSBearerContexts(nasMessage.PDUSessionModificationRequestMappedEPSBearerContextsType)
	for _, table := range nasTypeMappedEPSBearerContextsMappedEPSBearerContextTable {
		a.SetLen(table.inLen)
		a.SetMappedEPSBearerContext(table.in)
		assert.Equal(t, table.out, a.GetMappedEPSBearerContext())
	}
}

type testMappedEPSBearerContextsDataTemplate struct {
	inIei                     uint8
	inLen                     uint16
	inMappedEPSBearerContext  []uint8
	outIei                    uint8
	outLen                    uint16
	outMappedEPSBearerContext []uint8
}

var testMappedEPSBearerContextsTestTable = []testMappedEPSBearerContextsDataTemplate{
	{nasMessage.PDUSessionModificationRequestMappedEPSBearerContextsType, 2, []uint8{0xff, 0xff},
		nasMessage.PDUSessionModificationRequestMappedEPSBearerContextsType, 2, []uint8{0xff, 0xff}},
}

func TestNasTypeMappedEPSBearerContexts(t *testing.T) {

	for i, table := range testMappedEPSBearerContextsTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewMappedEPSBearerContexts(nasMessage.PDUSessionModificationRequestMappedEPSBearerContextsType)

		a.SetIei(table.inIei)
		a.SetLen(table.inLen)
		a.SetMappedEPSBearerContext(table.inMappedEPSBearerContext)

		assert.Equalf(t, table.outIei, a.Iei, "in(%v): out %v, actual %x", table.inIei, table.outIei, a.Iei)
		assert.Equalf(t, table.outLen, a.Len, "in(%v): out %v, actual %x", table.inLen, table.outLen, a.Len)
		assert.Equalf(t, table.outMappedEPSBearerContext, a.GetMappedEPSBearerContext(), "in(%v): out %v, actual %x", table.inMappedEPSBearerContext, table.outMappedEPSBearerContext, a.GetMappedEPSBearerContext())
	}
}
