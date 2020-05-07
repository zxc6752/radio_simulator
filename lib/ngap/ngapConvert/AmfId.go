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

func AmfIdToNgap(amfId string) (regionId, setId, ptrId aper.BitString) {
	regionId = HexToBitString(amfId[:2], 8)
	setId = HexToBitString(amfId[2:5], 10)
	tmpByte, err := hex.DecodeString(amfId[4:])
	if err != nil {
		logger.NgapLog.Warningln("AmfId From Models To NGAP Error: ", err.Error())
		return
	}
	shiftByte, err := aper.GetBitString(tmpByte, 2, 6)
	if err != nil {
		logger.NgapLog.Warningln("AmfId From Models To NGAP Error: ", err.Error())
		return
	}
	ptrId.BitLength = 6
	ptrId.Bytes = shiftByte
	return
}

func AmfIdToModels(regionId, setId, ptrId aper.BitString) (amfId string) {
	regionHex := BitStringToHex(&regionId)
	tmpByte := []byte{setId.Bytes[0], (setId.Bytes[1] & 0xc0) | (ptrId.Bytes[0] >> 2)}
	restHex := hex.EncodeToString(tmpByte)
	amfId = regionHex + restHex
	return
}
