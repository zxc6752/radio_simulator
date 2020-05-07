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

func TestNasTypeNewServiceRequestMessageIdentity(t *testing.T) {
	a := nasType.NewServiceRequestMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypeServiceRequestMessageIdentity struct {
	in  uint8
	out uint8
}

var nasTypeServiceRequestMessageIdentityTable = []nasTypeServiceRequestMessageIdentity{
	{0x03, 0x03},
}

func TestNasTypeServiceRequestMessageIdentityGetSetMessageType(t *testing.T) {
	a := nasType.NewServiceRequestMessageIdentity()
	for _, table := range nasTypeServiceRequestMessageIdentityTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}

type ServiceRequestMessageIdentityTestDataTemplate struct {
	in  nasType.ServiceRequestMessageIdentity
	out nasType.ServiceRequestMessageIdentity
}

var ServiceRequestMessageIdentityTestData = []nasType.ServiceRequestMessageIdentity{
	{0x03},
}

var ServiceRequestMessageIdentityExpectedTestData = []nasType.ServiceRequestMessageIdentity{
	{0x03},
}

var ServiceRequestMessageIdentityTable = []ServiceRequestMessageIdentityTestDataTemplate{
	{ServiceRequestMessageIdentityTestData[0], ServiceRequestMessageIdentityExpectedTestData[0]},
}

func TestNasTypeServiceRequestMessageIdentity(t *testing.T) {

	for _, table := range ServiceRequestMessageIdentityTable {

		a := nasType.NewServiceRequestMessageIdentity()

		a.SetMessageType(table.in.GetMessageType())
		assert.Equal(t, table.out.GetMessageType(), a.GetMessageType())
	}
}
