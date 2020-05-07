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

package milenage_test

import (
	"encoding/hex"
	"fmt"
	"radio_simulator/lib/milenage"
	"testing"
)

func TestGenerateOPC(t *testing.T) {
	_, _, OPC := make([]byte, 16), make([]byte, 16), make([]byte, 16)

	// OPC_str := ""
	// OPC, _ = hex.DecodeString(OPC_str)
	// fmt.Println("OPC", OPC)

	K_str := "3016ebeae2c45bd0060923dbbb402be6"
	K, _ := hex.DecodeString(K_str)

	OP_str := "00000000000000000000000000000000"
	OP, _ := hex.DecodeString(OP_str)
	fmt.Println("K:", K)

	fmt.Println("OP:", OP)

	milenage.GenerateOPC(K, OP, OPC)
}
