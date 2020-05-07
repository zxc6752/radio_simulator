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
	"encoding/hex"
	"radio_simulator/lib/aper"
	"radio_simulator/lib/ngap/logger"
)

func BitStringToHex(bitString *aper.BitString) (hexString string) {
	hexString = hex.EncodeToString(bitString.Bytes)
	hexLen := (bitString.BitLength + 3) / 4
	hexString = hexString[:hexLen]
	return
}

func HexToBitString(hexString string, bitLength int) (bitString aper.BitString) {
	hexLen := len(hexString)
	if hexLen != (bitLength+3)/4 {
		logger.NgapLog.Warningln("hexLen[", hexLen, "] doesn't match bitLength[", bitLength, "]")
		return
	}
	if hexLen%2 == 1 {
		hexString += "0"
	}
	bitString.Bytes, _ = hex.DecodeString(hexString)
	bitString.BitLength = uint64(bitLength)
	mask := byte(0xff)
	mask = mask << uint(8-bitLength%8)
	if mask != 0 {
		bitString.Bytes[len(bitString.Bytes)-1] &= mask
	}
	return
}

func ByteToBitString(byteArray []byte, bitLength int) (bitString aper.BitString) {
	byteLen := (bitLength + 7) / 8
	if byteLen > len(byteArray) {
		logger.NgapLog.Warningln("bitLength[", bitLength, "] is beyond byteArray size[", len(byteArray), "]")
		return
	}
	bitString.Bytes = byteArray
	bitString.BitLength = uint64(bitLength)
	return
}
