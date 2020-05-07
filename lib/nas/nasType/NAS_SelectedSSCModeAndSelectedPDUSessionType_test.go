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

func TestNasTypeNewSelectedSSCModeAndSelectedPDUSessionType(t *testing.T) {
	a := nasType.NewSelectedSSCModeAndSelectedPDUSessionType()
	assert.NotNil(t, a)
}

type nasTypeSelectedSSCModeAndSelectedPDUSessionTypeSSCModeData struct {
	in  uint8
	out uint8
}

var nasTypeSelectedSSCModeAndSelectedPDUSessionTypeSSCModeTable = []nasTypeSelectedSSCModeAndSelectedPDUSessionTypeSSCModeData{
	{0x01, 0x01},
}

func TestNasTypeSelectedSSCModeAndSelectedPDUSessionTypeGetSetSSCMode(t *testing.T) {
	a := nasType.NewSelectedSSCModeAndSelectedPDUSessionType()
	for _, table := range nasTypeSelectedSSCModeAndSelectedPDUSessionTypeSSCModeTable {
		a.SetSSCMode(table.in)
		assert.Equal(t, table.out, a.GetSSCMode())
	}
}

type nasTypeSelectedPDUSessionTypeAndSelectedPDUSessionTypePDUSessionTypeData struct {
	in  uint8
	out uint8
}

var nasTypeSelectedPDUSessionTypeAndSelectedPDUSessionTypePDUSessionTypeTable = []nasTypeSelectedPDUSessionTypeAndSelectedPDUSessionTypePDUSessionTypeData{
	{0x01, 0x01},
}

func TestNasTypeSelectedPDUSessionTypeAndSelectedPDUSessionTypeGetSetPDUSessionType(t *testing.T) {
	a := nasType.NewSelectedSSCModeAndSelectedPDUSessionType()
	for _, table := range nasTypeSelectedPDUSessionTypeAndSelectedPDUSessionTypePDUSessionTypeTable {
		a.SetPDUSessionType(table.in)
		assert.Equal(t, table.out, a.GetPDUSessionType())
	}
}

type SelectedSSCModeAndSelectedPDUSessionTypeTestDataTemplate struct {
	in  nasType.SelectedSSCModeAndSelectedPDUSessionType
	out nasType.SelectedSSCModeAndSelectedPDUSessionType
}

var SelectedSSCModeAndSelectedPDUSessionTypeTestData = []nasType.SelectedSSCModeAndSelectedPDUSessionType{
	{0x00},
}

var SelectedSSCModeAndSelectedPDUSessionTypeExpectedTestData = []nasType.SelectedSSCModeAndSelectedPDUSessionType{
	{0x11},
}

var SelectedSSCModeAndSelectedPDUSessionTypeTable = []SelectedSSCModeAndSelectedPDUSessionTypeTestDataTemplate{
	{SelectedSSCModeAndSelectedPDUSessionTypeTestData[0], SelectedSSCModeAndSelectedPDUSessionTypeExpectedTestData[0]},
}

func TestNasTypeSelectedSSCModeAndSelectedPDUSessionType(t *testing.T) {

	for _, table := range SelectedSSCModeAndSelectedPDUSessionTypeTable {

		a := nasType.NewSelectedSSCModeAndSelectedPDUSessionType()
		a.SetSSCMode(0x01)
		a.SetPDUSessionType(0x01)

		assert.Equal(t, table.out.Octet, a.Octet)
	}
}
