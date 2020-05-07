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

func TestNasTypeNewLocalTimeZone(t *testing.T) {
	a := nasType.NewLocalTimeZone(nasMessage.ConfigurationUpdateCommandLocalTimeZoneType)
	assert.NotNil(t, a)

}

var nasTypeConfigurationUpdateCommandLocalTimeZoneTable = []NasTypeIeiData{
	{nasMessage.ConfigurationUpdateCommandLocalTimeZoneType, nasMessage.ConfigurationUpdateCommandLocalTimeZoneType},
}

func TestNasTypeLocalTimeZoneGetSetIei(t *testing.T) {
	a := nasType.NewLocalTimeZone(nasMessage.ConfigurationUpdateCommandLocalTimeZoneType)
	for _, table := range nasTypeConfigurationUpdateCommandLocalTimeZoneTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

type nasTypeLocalTimeZoneTimeZoneData struct {
	in  uint8
	out uint8
}

var nasTypeLocalTimeZoneOctetTable = []nasTypeLocalTimeZoneTimeZoneData{
	{0xff, 0xff},
}

func TestNasTypeLocalTimeZoneGetSetTimeZone(t *testing.T) {
	a := nasType.NewLocalTimeZone(nasMessage.ConfigurationUpdateCommandLocalTimeZoneType)
	for _, table := range nasTypeLocalTimeZoneOctetTable {
		a.SetTimeZone(table.in)
		assert.Equal(t, table.out, a.GetTimeZone())
	}
}

type testLocalTimeZoneDataTemplate struct {
	in  nasType.LocalTimeZone
	out nasType.LocalTimeZone
}

var LocalTimeZoneTestData = []nasType.LocalTimeZone{
	{nasMessage.ConfigurationUpdateCommandLocalTimeZoneType, 0xff},
}

var LocalTimeZoneExpectedTestData = []nasType.LocalTimeZone{
	{nasMessage.ConfigurationUpdateCommandLocalTimeZoneType, 0xff},
}

var LocalTimeZoneTestTable = []testLocalTimeZoneDataTemplate{
	{LocalTimeZoneTestData[0], LocalTimeZoneExpectedTestData[0]},
}

func TestNasTypeLocalTimeZone(t *testing.T) {

	for i, table := range LocalTimeZoneTestTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewLocalTimeZone(nasMessage.ConfigurationUpdateCommandLocalTimeZoneType)

		a.SetIei(table.in.GetIei())
		a.SetTimeZone(table.in.Octet)

		assert.Equalf(t, table.out.Iei, a.Iei, "in(%v): out %v, actual %x", table.in.Iei, table.out.Iei, a.Iei)
		assert.Equalf(t, table.out.Octet, a.Octet, "in(%v): out %v, actual %x", table.in.Octet, table.out.Octet, a.Octet)

	}
}
