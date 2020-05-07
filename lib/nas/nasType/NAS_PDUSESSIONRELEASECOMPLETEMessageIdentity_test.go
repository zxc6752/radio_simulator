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

func TestNasTypeNewPDUSESSIONRELEASECOMPLETEMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONRELEASECOMPLETEMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypePDUSESSIONRELEASECOMPLETEMessageIdentity struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONRELEASECOMPLETEMessageIdentityTable = []nasTypePDUSESSIONRELEASECOMPLETEMessageIdentity{
	{0x03, 0x03},
}

func TestNasTypePDUSESSIONRELEASECOMPLETEMessageIdentityGetSetMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONRELEASECOMPLETEMessageIdentity()
	for _, table := range nasTypePDUSESSIONRELEASECOMPLETEMessageIdentityTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}

type PDUSESSIONRELEASECOMPLETEMessageIdentityTestDataTemplate struct {
	in  nasType.PDUSESSIONRELEASECOMPLETEMessageIdentity
	out nasType.PDUSESSIONRELEASECOMPLETEMessageIdentity
}

var PDUSESSIONRELEASECOMPLETEMessageIdentityTestData = []nasType.PDUSESSIONRELEASECOMPLETEMessageIdentity{
	{0x03},
}

var PDUSESSIONRELEASECOMPLETEMessageIdentityExpectedTestData = []nasType.PDUSESSIONRELEASECOMPLETEMessageIdentity{
	{0x03},
}

var PDUSESSIONRELEASECOMPLETEMessageIdentityTable = []PDUSESSIONRELEASECOMPLETEMessageIdentityTestDataTemplate{
	{PDUSESSIONRELEASECOMPLETEMessageIdentityTestData[0], PDUSESSIONRELEASECOMPLETEMessageIdentityExpectedTestData[0]},
}

func TestNasTypePDUSESSIONRELEASECOMPLETEMessageIdentity(t *testing.T) {

	for _, table := range PDUSESSIONRELEASECOMPLETEMessageIdentityTable {

		a := nasType.NewPDUSESSIONRELEASECOMPLETEMessageIdentity()

		a.SetMessageType(table.in.GetMessageType())
		assert.Equal(t, table.out.GetMessageType(), a.GetMessageType())
	}
}
