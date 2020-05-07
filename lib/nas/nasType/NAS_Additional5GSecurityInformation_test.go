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

func TestNasTypeNewAdditional5GSecurityInformation(t *testing.T) {
	a := nasType.NewAdditional5GSecurityInformation(0x36) // security mode command message
	assert.NotNil(t, a)
}

var nasTypeSecurityModeCommandAdditional5GSecurityInformationTable = []NasTypeIeiData{
	{0x36, 0x36},
}

func TestNasTypeAdditional5GSecurityInformationGetSetIei(t *testing.T) {
	a := nasType.NewAdditional5GSecurityInformation(0x36)
	for _, table := range nasTypeSecurityModeCommandAdditional5GSecurityInformationTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

var nasTypeSecurityModeCommandAdditional5GSecurityInformationLenTable = []NasTypeLenuint8Data{
	{2, 2},
}

func TestNasTypeAdditional5GSecurityInformationGetSetLen(t *testing.T) {
	a := nasType.NewAdditional5GSecurityInformation(0x36)
	for _, table := range nasTypeSecurityModeCommandAdditional5GSecurityInformationLenTable {
		a.SetLen(table.in)
		assert.Equal(t, table.out, a.GetLen())
	}
}

type NasTypeRINMRuint8Data struct {
	in  uint8
	out uint8
}

var nasTypeAdditional5GSecurityInformationRINMR = []NasTypeRINMRuint8Data{
	{0x1, 0x1},
	{0x0, 0x0},
}

func TestNasTypeAdditional5GSecurityInformationGetSetRINMR(t *testing.T) {
	a := nasType.NewAdditional5GSecurityInformation(0x36)
	for _, table := range nasTypeAdditional5GSecurityInformationRINMR {
		a.SetRINMR(table.in)
		assert.Equal(t, table.out, a.GetRINMR())
	}
}

type NasTypeHDPuint8Data struct {
	in  uint8
	out uint8
}

var nasTypeAdditional5GSecurityInformationHDP = []NasTypeHDPuint8Data{
	{0x1, 0x1},
	{0x0, 0x0},
}

func TestNasTypeAdditional5GSecurityInformationGetSetHDP(t *testing.T) {
	a := nasType.NewAdditional5GSecurityInformation(0x36)
	for _, table := range nasTypeAdditional5GSecurityInformationHDP {
		a.SetHDP(table.in)
		assert.Equal(t, table.out, a.GetHDP())
	}
}
