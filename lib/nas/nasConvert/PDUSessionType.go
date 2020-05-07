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
	"radio_simulator/lib/openapi/models"
)

func PDUSessionTypeToModels(nasPduSessType uint8) (pduSessType models.PduSessionType) {
	switch nasPduSessType {
	case nasMessage.PDUSessionTypeIPv4:
		pduSessType = models.PduSessionType_IPV4
	case nasMessage.PDUSessionTypeIPv6:
		pduSessType = models.PduSessionType_IPV6
	case nasMessage.PDUSessionTypeIPv4IPv6:
		pduSessType = models.PduSessionType_IPV4_V6
	case nasMessage.PDUSessionTypeUnstructured:
		pduSessType = models.PduSessionType_UNSTRUCTURED
	case nasMessage.PDUSessionTypeEthernet:
		pduSessType = models.PduSessionType_ETHERNET
	}

	return
}

func ModelsToPDUSessionType(pduSessType models.PduSessionType) (nasPduSessType uint8) {
	switch pduSessType {
	case models.PduSessionType_IPV4:
		nasPduSessType = nasMessage.PDUSessionTypeIPv4
	case models.PduSessionType_IPV6:
		nasPduSessType = nasMessage.PDUSessionTypeIPv6
	case models.PduSessionType_IPV4_V6:
		nasPduSessType = nasMessage.PDUSessionTypeIPv4IPv6
	case models.PduSessionType_UNSTRUCTURED:
		nasPduSessType = nasMessage.PDUSessionTypeUnstructured
	case models.PduSessionType_ETHERNET:
		nasPduSessType = nasMessage.PDUSessionTypeEthernet
	}
	return
}
