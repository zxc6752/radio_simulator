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

func TestNasTypeNewPDUSESSIONRELEASEREQUESTMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONRELEASEREQUESTMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypePDUSESSIONRELEASEREQUESTMessageIdentity struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONRELEASEREQUESTMessageIdentityTable = []nasTypePDUSESSIONRELEASEREQUESTMessageIdentity{
	{0x03, 0x03},
}

func TestNasTypePDUSESSIONRELEASEREQUESTMessageIdentityGetSetMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONRELEASEREQUESTMessageIdentity()
	for _, table := range nasTypePDUSESSIONRELEASEREQUESTMessageIdentityTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}

type PDUSESSIONRELEASEREQUESTMessageIdentityTestDataTemplate struct {
	in  nasType.PDUSESSIONRELEASEREQUESTMessageIdentity
	out nasType.PDUSESSIONRELEASEREQUESTMessageIdentity
}

var PDUSESSIONRELEASEREQUESTMessageIdentityTestData = []nasType.PDUSESSIONRELEASEREQUESTMessageIdentity{
	{0x03},
}

var PDUSESSIONRELEASEREQUESTMessageIdentityExpectedTestData = []nasType.PDUSESSIONRELEASEREQUESTMessageIdentity{
	{0x03},
}

var PDUSESSIONRELEASEREQUESTMessageIdentityTable = []PDUSESSIONRELEASEREQUESTMessageIdentityTestDataTemplate{
	{PDUSESSIONRELEASEREQUESTMessageIdentityTestData[0], PDUSESSIONRELEASEREQUESTMessageIdentityExpectedTestData[0]},
}

func TestNasTypePDUSESSIONRELEASEREQUESTMessageIdentity(t *testing.T) {

	for _, table := range PDUSESSIONRELEASEREQUESTMessageIdentityTable {

		a := nasType.NewPDUSESSIONRELEASEREQUESTMessageIdentity()

		a.SetMessageType(table.in.GetMessageType())
		assert.Equal(t, table.out.GetMessageType(), a.GetMessageType())
	}
}
