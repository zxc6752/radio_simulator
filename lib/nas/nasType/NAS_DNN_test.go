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

func TestNasTypeNewDNN(t *testing.T) {
	a := nasType.NewDNN(0)
	assert.NotNil(t, a)
}

var nasTypeDNNIeiTable = []NasTypeIeiData{
	{0, 0},
}

func TestNasTypDNNGetSetIei(t *testing.T) {
	a := nasType.NewDNN(0)
	for _, table := range nasTypeDNNIeiTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

var nasTypeDNNLenTable = []NasTypeLenuint8Data{
	{1, 1},
}

func TestNasTypeDNNGetSetLen(t *testing.T) {
	a := nasType.NewDNN(0)
	for _, table := range nasTypeDNNLenTable {
		a.SetLen(table.in)
		assert.Equal(t, table.out, a.GetLen())
	}
}

type nasTypetDNNData struct {
	inLen uint8
	in    []uint8
	out   []uint8
}

var nasTypeDNNTable = []nasTypetDNNData{
	{8, []uint8{0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74}, []uint8{0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74}},
}

func TestNasTypeDNNGetSetDNNValue(t *testing.T) {
	a := nasType.NewDNN(0)
	for _, table := range nasTypeDNNTable {
		a.SetLen(table.inLen)
		a.SetDNN(table.in)
		assert.Equalf(t, table.out, a.GetDNN(), "in(%v): out %v, actual %x", table.in, table.out, a.GetDNN())
	}
}

type testDNNDataTemplate struct {
	in  nasType.DNN
	out nasType.DNN
}

var DNNTestData = []nasType.DNN{
	{0, 7, []byte{0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74}}, //AuthenticationResult
}

var DNNExpectedTestData = []nasType.DNN{
	{0, 8, []byte{0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74}}, //AuthenticationResult
}

var DNNTestTable = []testDNNDataTemplate{
	{DNNTestData[0], DNNExpectedTestData[0]},
}

func TestNasTypeDNN(t *testing.T) {

	for i, table := range DNNTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewDNN(0) //AuthenticationResult

		a.SetIei(table.in.GetIei())
		a.SetLen(table.in.Len)
		a.SetDNN(table.in.Buffer)

		assert.Equalf(t, table.out.Iei, a.Iei, "in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		assert.Equalf(t, table.out.Len, a.Len, "in(%v): out %v, actual %x", table.in.Len, table.out.Len, a.Len)
		assert.Equalf(t, table.out.Buffer, a.Buffer, "in(%v): out %v, actual %x", table.in.Buffer, table.out.Buffer, a.Buffer)
		t.Log(table.out.Buffer, a.Buffer, "in(%v): out %v, actual %x", table.in.Buffer, table.out.Buffer, a.Buffer)
		t.Log(a.Len)

	}
}