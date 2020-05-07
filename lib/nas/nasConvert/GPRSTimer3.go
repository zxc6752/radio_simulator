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
	"radio_simulator/lib/nas/nasMessage"
)

// TS 24.008 10.5.7.4a
func GPRSTimer3ToNas(timerValue int) (timerValueNas uint8) {

	if timerValue <= 2*31 {
		t := uint8(timerValue / 2)
		timerValueNas = (nasMessage.GPRSTimer3UnitMultiplesOf2Seconds << 5) + t
	} else if timerValue <= 30*31 {
		t := uint8(timerValue / 30)
		timerValueNas = (nasMessage.GPRSTimer3UnitMultiplesOf30Seconds << 5) + t
	} else if timerValue <= 60*31 {
		t := uint8(timerValue / 60)
		timerValueNas = (nasMessage.GPRSTimer3UnitMultiplesOf1Minute << 5) + t
	} else if timerValue <= 600*31 {
		t := uint8(timerValue / 600)
		timerValueNas = (nasMessage.GPRSTimer3UnitMultiplesOf10Minutes << 5) + t
	} else if timerValue <= 3600*31 {
		t := uint8(timerValue / 3600)
		timerValueNas = (nasMessage.GPRSTimer3UnitMultiplesOf1Hour << 5) + t
	} else {
		t := uint8(timerValue / (36000))
		timerValueNas = (nasMessage.GPRSTimer3UnitMultiplesOf10Hours << 5) + t
	}

	return
}
