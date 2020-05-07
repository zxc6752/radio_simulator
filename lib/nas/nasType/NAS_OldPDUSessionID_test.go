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

func TestNasTypeNewOldPDUSessionID(t *testing.T) {
	a := nasType.NewOldPDUSessionID(nasMessage.ULNASTransportOldPDUSessionIDType)
	assert.NotNil(t, a)

}

var nasTypeOldPDUSessionIDULNASTransportOldPDUSessionIDTypeTable = []NasTypeIeiData{
	{nasMessage.ULNASTransportOldPDUSessionIDType, nasMessage.ULNASTransportOldPDUSessionIDType},
}

func TestNasTypeOldPDUSessionIDGetSetIei(t *testing.T) {
	a := nasType.NewOldPDUSessionID(nasMessage.ULNASTransportOldPDUSessionIDType)
	for _, table := range nasTypeOldPDUSessionIDULNASTransportOldPDUSessionIDTypeTable {
		a.SetIei(table.in)
		assert.Equal(t, table.out, a.GetIei())
	}
}

type nasTypeOldPDUSessionIDPduSessionIdentity2Value struct {
	in  uint8
	out uint8
}

var nasTypeOldPDUSessionIDPduSessionIdentity2ValueTable = []nasTypeOldPDUSessionIDPduSessionIdentity2Value{
	{0xff, 0xff},
}

func TestNasTypeOldPDUSessionIDGetSetOldPDUSessionID(t *testing.T) {
	a := nasType.NewOldPDUSessionID(nasMessage.ULNASTransportOldPDUSessionIDType)
	for _, table := range nasTypeOldPDUSessionIDPduSessionIdentity2ValueTable {
		a.SetOldPDUSessionID(table.in)
		assert.Equal(t, table.out, a.GetOldPDUSessionID())
	}
}

type nasTypeOldPDUSessionID struct {
	inIei                       uint8
	inPduSessionIdentity2Value  uint8
	outIei                      uint8
	outPduSessionIdentity2Value uint8
}

var nasTypeOldPDUSessionIDTable = []nasTypeOldPDUSessionID{
	{nasMessage.ULNASTransportOldPDUSessionIDType, 0xff,
		nasMessage.ULNASTransportOldPDUSessionIDType, 0xff},
}

func TestNasTypeOldPDUSessionID(t *testing.T) {
	a := nasType.NewOldPDUSessionID(nasMessage.ULNASTransportOldPDUSessionIDType)
	for _, table := range nasTypeOldPDUSessionIDTable {
		a.SetIei(table.inIei)
		a.SetOldPDUSessionID(table.inPduSessionIdentity2Value)
		assert.Equal(t, table.outIei, a.GetIei())
		assert.Equal(t, table.outPduSessionIdentity2Value, a.GetOldPDUSessionID())
	}
}
