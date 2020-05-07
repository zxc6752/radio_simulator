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

func TestNasTypeNewRequestType(t *testing.T) {
	a := nasType.NewRequestType(nasMessage.ULNASTransportRequestTypeType)
	assert.NotNil(t, a)
}

var nasTypeRequestTypeIeiTable = []NasTypeIeiData{
	{0x08, 0x08},
}

func TestNasTypeRequestTypeGetSetIei(t *testing.T) {
	a := nasType.NewRequestType(nasMessage.ULNASTransportRequestTypeType)
	assert.NotNil(t, a)
	for _, table := range nasTypeRequestTypeIeiTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

type nasTypeRequestRequestTypeValueType struct {
	in  uint8
	out uint8
}

var nasTypeRequestTypeRequestTypeValueTable = []nasTypeRequestRequestTypeValueType{
	{0x03, 0x03},
}

func TestNasTypeRequestTypeGetSetRequestTypeValue(t *testing.T) {
	a := nasType.NewRequestType(nasMessage.ULNASTransportRequestTypeType)
	for _, table := range nasTypeRequestTypeRequestTypeValueTable {
		a.SetRequestTypeValue(table.in)
		assert.Equal(t, table.out, a.GetRequestTypeValue())
	}
}

type RequestTypeTestDataTemplate struct {
	in  nasType.RequestType
	out nasType.RequestType
}

var RequestTypeTestData = []nasType.RequestType{
	{nasMessage.ULNASTransportRequestTypeType + 0x01},
}

var RequestTypeExpectedTestData = []nasType.RequestType{
	{0x81},
}

var RequestTypeTable = []RequestTypeTestDataTemplate{
	{RequestTypeTestData[0], RequestTypeExpectedTestData[0]},
}

func TestNasTypeRequestType(t *testing.T) {

	for _, table := range RequestTypeTable {

		a := nasType.NewRequestType(nasMessage.ULNASTransportRequestTypeType)
		a.SetIei(0x08)
		a.SetRequestTypeValue(0x01)

		assert.Equal(t, table.out.Octet, a.Octet)

	}
}
