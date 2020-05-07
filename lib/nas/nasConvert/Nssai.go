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
	"fmt"
	"radio_simulator/lib/nas/nasMessage"
	"radio_simulator/lib/nas/nasType"
	"radio_simulator/lib/openapi/models"
)

func RequestedNssaiToModels(nasNssai *nasType.RequestedNSSAI) (nssai []models.Snssai) {

	buf := nasNssai.GetSNSSAIValue()
	lengthOfBuf := int(nasNssai.GetLen())
	offset := 0
	for offset < lengthOfBuf {
		snssaiValue := buf[offset:]
		snssai, readLength := requestedSnssaiToModels(snssaiValue)
		nssai = append(nssai, snssai)
		offset += readLength
	}

	return

}

func requestedSnssaiToModels(buf []byte) (snssai models.Snssai, length int) {

	lengthOfSnssaiContents := buf[0]
	switch lengthOfSnssaiContents {
	case 0x01: // sst
		snssai.Sst = int32(buf[1])
		length = 2
	case 0x04: // sst + sd
		snssai.Sst = int32(buf[1])
		snssai.Sd = hex.EncodeToString(buf[2:5])
		length = 5
	default:
		fmt.Printf("Not Supported length: %d\n", lengthOfSnssaiContents)
	}

	return
}

func RejectedNssaiToNas(rejectedNssaiInPlmn []models.Snssai, rejectedNssaiInTa []models.Snssai) (rejectedNssaiNas nasType.RejectedNSSAI) {

	var byteArray []uint8
	for _, rejectedSnssai := range rejectedNssaiInPlmn {
		byteArray = append(byteArray, RejectedSnssaiToNas(rejectedSnssai, nasMessage.RejectedSnssaiCauseNotAvailableInCurrentPlmn)...)
	}
	for _, rejectedSnssai := range rejectedNssaiInTa {
		byteArray = append(byteArray, RejectedSnssaiToNas(rejectedSnssai, nasMessage.RejectedSnssaiCauseNotAvailableInCurrentRegistrationArea)...)
	}

	rejectedNssaiNas.SetLen(uint8(len(byteArray)))
	rejectedNssaiNas.SetRejectedNSSAIContents(byteArray)
	return
}
