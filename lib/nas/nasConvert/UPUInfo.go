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
	"radio_simulator/lib/openapi/models"
)

//  subclause 9.11.3.53A in 3GPP TS 24.501
func UpuInfoToNas(upuInfo models.UpuInfo) (buf []uint8) {
	// set upu Header
	buf = append(buf, upuInfoGetHeader(upuInfo.UpuRegInd, upuInfo.UpuAckInd))
	// Set UPU-MAC-IAUSF
	byteArray, _ := hex.DecodeString(upuInfo.UpuMacIausf)
	buf = append(buf, byteArray...)
	// Set Counter UPU
	byteArray, _ = hex.DecodeString(upuInfo.CounterUpu)
	buf = append(buf, byteArray...)
	// Set UE parameters update list
	for _, data := range upuInfo.UpuDataList {
		if data.SecPacket != "" {
			buf = append(buf, 0x01)
			byteArray, _ = hex.DecodeString(data.SecPacket)
		} else {
			buf = append(buf, 0x02)
			byteArray = []byte{}
			for _, snssai := range data.DefaultConfNssai {
				snssaiData := SnssaiToNas(snssai)
				byteArray = append(byteArray, snssaiData...)
			}
		}
		buf = append(buf, uint8(len(byteArray)))
		buf = append(buf, byteArray...)
	}
	return
}

func upuInfoGetHeader(reg bool, ack bool) (buf uint8) {
	var regValue, ackValue uint8
	if reg {
		regValue = 1
	}
	if ack {
		ackValue = 1
	}
	buf = regValue<<2 + ackValue<<1
	return
}

func UpuAckToModels(buf []uint8) (string, error) {
	if (buf[0] != 0x01) || (len(buf) != 17) {
		return "", fmt.Errorf("NAS UPU Ack is not valid")
	}
	return hex.EncodeToString(buf[1:]), nil
}
