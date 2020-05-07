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
	"radio_simulator/lib/nas/logger"
)

// TS 24.008 10.5.7.4, TS 24.501 9.11.2.4
// the unit of timerValue is second
func GPRSTimer2ToNas(timerValue int) (timerValueNas uint8) {

	timerValueNas = 0

	if timerValue <= 64 {
		if timerValue%2 != 0 {
			logger.ConvertLog.Error("timer Value is not multiples of 2 seconds")
			return
		}
		timerValueNas = uint8(timerValue / 2)
	} else {
		t := uint8(timerValue / 60) // t is multiples of 1 min
		if t <= 31 {
			timerValueNas = (timerValueNas | 0x20) + t
		} else {
			if t%6 != 0 {
				logger.ConvertLog.Error("timer Value is not multiples of decihours")
				return
			}
			t = t / 6
			timerValueNas = (timerValueNas | 0x40) + t
		}
	}

	return
}
