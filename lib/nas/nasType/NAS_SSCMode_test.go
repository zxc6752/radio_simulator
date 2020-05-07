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

func TestNasTypeNewSSCMode(t *testing.T) {
	a := nasType.NewSSCMode(nasMessage.PDUSessionEstablishmentRequestSSCModeType)
	assert.NotNil(t, a)
}

var nasTypeSSCModeIeiTable = []NasTypeIeiData{
	{0x01, 0x01},
}

func TestNasTypeSSCModeGetSetIei(t *testing.T) {
	a := nasType.NewSSCMode(nasMessage.PDUSessionEstablishmentRequestSSCModeType)
	assert.NotNil(t, a)
	for _, table := range nasTypeSSCModeIeiTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

type nasTypeRequestSpareType struct {
	in  uint8
	out uint8
}

var nasTypeSSCModeSpareTable = []nasTypeRequestSpareType{
	{0x01, 0x01},
}

func TestNasTypeSSCModeGetSetSpare(t *testing.T) {
	a := nasType.NewSSCMode(nasMessage.PDUSessionEstablishmentRequestSSCModeType)
	for _, table := range nasTypeSSCModeSpareTable {
		a.SetSpare(table.in)
		assert.Equal(t, table.out, a.GetSpare())
	}
}

type nasTypeRequestSSCModeType struct {
	in  uint8
	out uint8
}

var nasTypeSSCModeSSCModeTable = []nasTypeRequestSSCModeType{
	{0x01, 0x01},
}

func TestNasTypeSSCModeGetSetSSCMode(t *testing.T) {
	a := nasType.NewSSCMode(nasMessage.PDUSessionEstablishmentRequestSSCModeType)
	for _, table := range nasTypeSSCModeSSCModeTable {
		a.SetSSCMode(table.in)
		assert.Equal(t, table.out, a.GetSSCMode())
	}
}

type SSCModeTestDataTemplate struct {
	in  nasType.SSCMode
	out nasType.SSCMode
}

var SSCModeTestData = []nasType.SSCMode{
	{nasMessage.PDUSessionEstablishmentRequestSSCModeType},
}

var SSCModeExpectedTestData = []nasType.SSCMode{
	{0x19},
}

var SSCModeTable = []SSCModeTestDataTemplate{
	{SSCModeTestData[0], SSCModeExpectedTestData[0]},
}

func TestNasTypeSSCMode(t *testing.T) {

	for _, table := range SSCModeTable {

		a := nasType.NewSSCMode(nasMessage.PDUSessionEstablishmentRequestSSCModeType)
		a.SetIei(0x01)
		a.SetSpare(0x01)
		a.SetSSCMode(0x01)

		assert.Equal(t, table.out.Octet, a.Octet)

	}
}
