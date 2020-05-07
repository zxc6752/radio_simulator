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
	"radio_simulator/lib/aper"
	"radio_simulator/lib/ngap/logger"
)

/*
RFC 5905 Section 6 https://tools.ietf.org/html/rfc5905#section-6

       0                   1                   2                   3
       0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
      +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
      |                            Seconds                            |
      +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
      |                            Fraction                           |
      +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

                             NTP Timestamp Format

   The 64-bit timestamp format is used in packet headers and other
   places with limited word size.  It includes a 32-bit unsigned seconds
   field spanning 136 years and a 32-bit fraction field resolving 232
   picoseconds.  The 32-bit short format is used in delay and dispersion
   header fields where the full resolution and range of the other
   formats are not justified.  It includes a 16-bit unsigned seconds
   field and a 16-bit fraction field.

   In the date and timestamp formats, the prime epoch, or base date of
   era 0, is 0 h 1 January 1900 UTC, when all bits are zero.  It should
   be noted that strictly speaking, UTC did not exist prior to 1 January
   1972, but it is convenient to assume it has existed for all eternity,
   even if all knowledge of historic leap seconds has been lost.  Dates
   are relative to the prime epoch; values greater than zero represent

*/
func TimeStampToInt32(timeStampNgap aper.OctetString) (timeStamp int32) {

	if len(timeStampNgap) != 4 {
		logger.NgapLog.Error("TimeStampToInt32: the size of OctetString is not 4")
	}

	timeStamp = int32(binary.BigEndian.Uint32(timeStampNgap))
	return
}

func TimeStampToNgap(timeStamp int32) (timeStampNgap aper.OctetString) {
	// TODO: finish this function when need
	return
}
