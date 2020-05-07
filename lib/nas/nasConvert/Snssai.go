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

package nasConvert

import (
	"encoding/hex"
	"radio_simulator/lib/nas/nasType"
	"radio_simulator/lib/openapi/models"
)

func SnssaiToModels(nasSnssai *nasType.SNSSAI) (snssai models.Snssai) {
	sD := nasSnssai.GetSD()
	snssai.Sd = hex.EncodeToString([]uint8(sD[:]))
	snssai.Sst = int32(nasSnssai.GetSST())
	return
}

func SnssaiToNas(snssai models.Snssai) (buf []uint8) {
	if snssai.Sd == "" {
		buf = append(buf, 0x01)
		buf = append(buf, uint8(snssai.Sst))
	} else {
		buf = append(buf, 0x04)
		buf = append(buf, uint8(snssai.Sst))
		byteArray, _ := hex.DecodeString(snssai.Sd)
		buf = append(buf, byteArray...)
	}
	return
}

func RejectedSnssaiToNas(snssai models.Snssai, rejectCause uint8) (rejectedSnssai []uint8) {

	if snssai.Sd == "" {
		rejectedSnssai = append(rejectedSnssai, (0x01<<4)+rejectCause)
		rejectedSnssai = append(rejectedSnssai, uint8(snssai.Sst))
	} else {
		rejectedSnssai = append(rejectedSnssai, (0x04<<4)+rejectCause)
		rejectedSnssai = append(rejectedSnssai, uint8(snssai.Sst))
		sDBytes, _ := hex.DecodeString(snssai.Sd)
		rejectedSnssai = append(rejectedSnssai, sDBytes...)
	}

	return
}
