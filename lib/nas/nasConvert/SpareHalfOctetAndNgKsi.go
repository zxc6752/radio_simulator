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
	"radio_simulator/lib/nas/nasType"
	"radio_simulator/lib/openapi/models"
)

func SpareHalfOctetAndNgksiToModels(ngKsiNas nasType.SpareHalfOctetAndNgksi) (ngKsiModels models.NgKsi) {

	switch ngKsiNas.GetTSC() {
	case nasMessage.TypeOfSecurityContextFlagNative:
		ngKsiModels.Tsc = models.ScType_NATIVE
	case nasMessage.TypeOfSecurityContextFlagMapped:
		ngKsiModels.Tsc = models.ScType_MAPPED
	}

	ngKsiModels.Ksi = int32(ngKsiNas.GetNasKeySetIdentifiler())
	return
}

func SpareHalfOctetAndNgksiToNas(ngKsiModels models.NgKsi) (ngKsiNas nasType.SpareHalfOctetAndNgksi) {

	switch ngKsiModels.Tsc {
	case models.ScType_NATIVE:
		ngKsiNas.SetTSC(nasMessage.TypeOfSecurityContextFlagNative)
	case models.ScType_MAPPED:
		ngKsiNas.SetTSC(nasMessage.TypeOfSecurityContextFlagMapped)
	}

	ngKsiNas.SetSpareHalfOctet(0)
	ngKsiNas.SetNasKeySetIdentifiler(uint8(ngKsiModels.Ksi))
	return
}
