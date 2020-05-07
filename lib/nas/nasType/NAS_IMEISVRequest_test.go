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

var SecurityModeCommandIMEISVRequestTypeIeiInput uint8 = 0x0E

func TestNasTypeNewIMEISVRequest(t *testing.T) {
	a := nasType.NewIMEISVRequest(SecurityModeCommandIMEISVRequestTypeIeiInput)
	assert.NotNil(t, a)
}

var nasTypePDUSessionEstablishmentRequestIMEISVRequestTable = []NasTypeIeiData{
	{SecurityModeCommandIMEISVRequestTypeIeiInput, 0x0E},
}

func TestNasTypeIMEISVRequestGetSetIei(t *testing.T) {
	a := nasType.NewIMEISVRequest(SecurityModeCommandIMEISVRequestTypeIeiInput)
	for _, table := range nasTypePDUSessionEstablishmentRequestIMEISVRequestTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

type nasTypeIMEISVRequestIMEISVRequestValue struct {
	in  uint8
	out uint8
}

var nasTypeIMEISVRequestIMEISVRequestValueTable = []nasTypeIMEISVRequestIMEISVRequestValue{
	{0x07, 0x07},
}

func TestNasTypeIMEISVRequestGetSetIMEISVRequestValue(t *testing.T) {
	a := nasType.NewIMEISVRequest(SecurityModeCommandIMEISVRequestTypeIeiInput)
	for _, table := range nasTypeIMEISVRequestIMEISVRequestValueTable {
		a.SetIMEISVRequestValue(table.in)
		assert.Equal(t, table.out, a.GetIMEISVRequestValue())
	}
}

type testIMEISVRequestDataTemplate struct {
	inIei                uint8
	inIMEISVRequestValue uint8

	outIei                uint8
	outIMEISVRequestValue uint8
}

var iMEISVRequestTestTable = []testIMEISVRequestDataTemplate{
	{SecurityModeCommandIMEISVRequestTypeIeiInput, 0x07,
		0x0E, 0x07},
}

func TestNasTypeIMEISVRequest(t *testing.T) {

	for _, table := range iMEISVRequestTestTable {
		a := nasType.NewIMEISVRequest(SecurityModeCommandIMEISVRequestTypeIeiInput)

		a.SetIei(table.inIei)
		a.SetIMEISVRequestValue(table.inIMEISVRequestValue)

		assert.Equalf(t, table.outIei, a.GetIei(), "in(%v): out %v, actual %x", table.inIei, table.outIei, a.GetIei())
		assert.Equalf(t, table.outIMEISVRequestValue, a.GetIMEISVRequestValue(), "in(%v): out %v, actual %x", table.inIMEISVRequestValue, table.outIMEISVRequestValue, a.GetIMEISVRequestValue())
	}
}
