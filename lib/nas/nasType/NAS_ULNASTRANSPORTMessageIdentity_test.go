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

func TestNasTypeNewULNASTRANSPORTMessageIdentity(t *testing.T) {
	a := nasType.NewULNASTRANSPORTMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypeULNASTRANSPORTMessageIdentity struct {
	in  uint8
	out uint8
}

var nasTypeULNASTRANSPORTMessageIdentityTable = []nasTypeULNASTRANSPORTMessageIdentity{
	{0x03, 0x03},
}

func TestNasTypeULNASTRANSPORTMessageIdentityGetSetMessageType(t *testing.T) {
	a := nasType.NewULNASTRANSPORTMessageIdentity()
	for _, table := range nasTypeULNASTRANSPORTMessageIdentityTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}

type ULNASTRANSPORTMessageIdentityTestDataTemplate struct {
	in  nasType.ULNASTRANSPORTMessageIdentity
	out nasType.ULNASTRANSPORTMessageIdentity
}

var ULNASTRANSPORTMessageIdentityTestData = []nasType.ULNASTRANSPORTMessageIdentity{
	{0x03},
}

var ULNASTRANSPORTMessageIdentityExpectedTestData = []nasType.ULNASTRANSPORTMessageIdentity{
	{0x03},
}

var ULNASTRANSPORTMessageIdentityTable = []ULNASTRANSPORTMessageIdentityTestDataTemplate{
	{ULNASTRANSPORTMessageIdentityTestData[0], ULNASTRANSPORTMessageIdentityExpectedTestData[0]},
}

func TestNasTypeULNASTRANSPORTMessageIdentity(t *testing.T) {

	for _, table := range ULNASTRANSPORTMessageIdentityTable {

		a := nasType.NewULNASTRANSPORTMessageIdentity()

		a.SetMessageType(table.in.GetMessageType())
		assert.Equal(t, table.out.GetMessageType(), a.GetMessageType())
	}
}
