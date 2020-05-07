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
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"

	"radio_simulator/lib/nas/nasMessage"
	"radio_simulator/lib/nas/nasType"
	"radio_simulator/lib/openapi/models"
)

func ModelsToSessionAMBR(ambr *models.Ambr) (sessAmbr nasType.SessionAMBR) {
	var bitRate int64
	var bitRateBytes [2]byte

	fmt.Println(ambr)

	uplink := strings.Split(ambr.Uplink, " ")
	bitRate, _ = strconv.ParseInt(uplink[0], 10, 16)
	binary.LittleEndian.PutUint16(bitRateBytes[:], uint16(bitRate))
	sessAmbr.SetSessionAMBRForUplink(bitRateBytes)
	sessAmbr.SetUnitForSessionAMBRForUplink(strToAMBRUnit(uplink[1]))

	downlink := strings.Split(ambr.Downlink, " ")
	bitRate, _ = strconv.ParseInt(downlink[0], 10, 16)
	binary.LittleEndian.PutUint16(bitRateBytes[:], uint16(bitRate))
	sessAmbr.SetSessionAMBRForDownlink(bitRateBytes)
	sessAmbr.SetUnitForSessionAMBRForDownlink(strToAMBRUnit(downlink[1]))
	return
}

func strToAMBRUnit(unit string) uint8 {
	switch unit {
	case "bps":
		return nasMessage.SessionAMBRUnitNotUsed
	case "Kbps":
		return nasMessage.SessionAMBRUnit1Kbps
	case "Mbps":
		return nasMessage.SessionAMBRUnit1Mbps
	case "Gbps":
		return nasMessage.SessionAMBRUnit1Gbps
	case "Tbps":
		return nasMessage.SessionAMBRUnit1Tbps
	case "Pbps":
		return nasMessage.SessionAMBRUnit1Pbps
	}
	return nasMessage.SessionAMBRUnitNotUsed
}
