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

func TestNasTypeNewSecurityModeCompleteMessageIdentity(t *testing.T) {
	a := nasType.NewSecurityModeCompleteMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypeSecurityModeCompleteMessageIdentity struct {
	in  uint8
	out uint8
}

var nasTypeSecurityModeCompleteMessageIdentityTable = []nasTypeSecurityModeCompleteMessageIdentity{
	{0x03, 0x03},
}

func TestNasTypeSecurityModeCompleteMessageIdentityGetSetMessageType(t *testing.T) {
	a := nasType.NewSecurityModeCompleteMessageIdentity()
	for _, table := range nasTypeSecurityModeCompleteMessageIdentityTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}

type SecurityModeCompleteMessageIdentityTestDataTemplate struct {
	in  nasType.SecurityModeCompleteMessageIdentity
	out nasType.SecurityModeCompleteMessageIdentity
}

var SecurityModeCompleteMessageIdentityTestData = []nasType.SecurityModeCompleteMessageIdentity{
	{0x03},
}

var SecurityModeCompleteMessageIdentityExpectedTestData = []nasType.SecurityModeCompleteMessageIdentity{
	{0x03},
}

var SecurityModeCompleteMessageIdentityTable = []SecurityModeCompleteMessageIdentityTestDataTemplate{
	{SecurityModeCompleteMessageIdentityTestData[0], SecurityModeCompleteMessageIdentityExpectedTestData[0]},
}

func TestNasTypeSecurityModeCompleteMessageIdentity(t *testing.T) {

	for _, table := range SecurityModeCompleteMessageIdentityTable {

		a := nasType.NewSecurityModeCompleteMessageIdentity()

		a.SetMessageType(table.in.GetMessageType())
		assert.Equal(t, table.out.GetMessageType(), a.GetMessageType())
	}
}
