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

package ngapConvert

import (
	"encoding/binary"
	"radio_simulator/lib/ngap/ngapType"
)

func PortNumberToInt(port ngapType.PortNumber) (portInt32 int32) {
	portInt32 = int32(binary.BigEndian.Uint16(port.Value))
	return
}

func PortNumberToNgap(portInt32 int32) (port ngapType.PortNumber) {
	port.Value = make([]byte, 2)
	binary.BigEndian.PutUint16(port.Value, uint16(portInt32))
	return
}
