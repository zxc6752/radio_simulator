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
	"radio_simulator/lib/openapi/models"
	"reflect"
)

// TS 24.501 9.11.3.9
func TaiListToNas(taiList []models.Tai) (taiListNas []uint8) {

	typeOfList := 0x00

	plmnId := taiList[0].PlmnId
	for _, tai := range taiList {
		if !reflect.DeepEqual(plmnId, tai.PlmnId) {
			typeOfList = 0x02
		}
	}

	numOfElementsNas := uint8(len(taiList)) - 1

	taiListNas = append(taiListNas, uint8(typeOfList<<5)+numOfElementsNas)

	switch typeOfList {
	case 0x00:
		plmnNas := PlmnIDToNas(*plmnId)
		taiListNas = append(taiListNas, plmnNas...)

		for _, tai := range taiList {
			tacBytes, _ := hex.DecodeString(tai.Tac)
			taiListNas = append(taiListNas, tacBytes...)
		}
	case 0x02:
		for _, tai := range taiList {
			plmnNas := PlmnIDToNas(*tai.PlmnId)
			tacBytes, _ := hex.DecodeString(tai.Tac)
			taiListNas = append(taiListNas, plmnNas...)
			taiListNas = append(taiListNas, tacBytes...)
		}
	}

	return
}
