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

func TestNasTypeNewAuthenticationFailureMessageIdentity(t *testing.T) {
	a := nasType.NewAuthenticationFailureMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypeMessageType struct {
	in  uint8
	out uint8
}

var nasTypeMessageTypeTable = []nasTypeMessageType{
	{0x03, 0x03},
}

func TestNasTypeGetSetMessageType(t *testing.T) {
	a := nasType.NewAuthenticationFailureMessageIdentity()
	for _, table := range nasTypeMessageTypeTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}

type AuthenticationFailureMessageIdentityTestDataTemplate struct {
	in  nasType.AuthenticationFailureMessageIdentity
	out nasType.AuthenticationFailureMessageIdentity
}

var authenticationFailureMessageIdentityTestData = []nasType.AuthenticationFailureMessageIdentity{
	{0x03},
}

var authenticationFailureMessageIdentityExpectedTestData = []nasType.AuthenticationFailureMessageIdentity{
	{0x03},
}

var authenticationFailureMessageIdentityTable = []AuthenticationFailureMessageIdentityTestDataTemplate{
	{authenticationFailureMessageIdentityTestData[0], authenticationFailureMessageIdentityExpectedTestData[0]},
}

func TestNasTypeAuthenticationFailureMessageIdentity(t *testing.T) {

	for i, table := range authenticationFailureMessageIdentityTable {
		t.Logf("Test Cnt:%d", i)
		a := nasType.NewAuthenticationFailureMessageIdentity()

		a.SetMessageType(table.in.GetMessageType())
		assert.Equal(t, table.out.GetMessageType(), a.GetMessageType())
	}
}
