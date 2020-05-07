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
	"encoding/hex"
	"radio_simulator/lib/ngap/ngapType"
	"radio_simulator/lib/openapi/models"
)

func SNssaiToModels(ngapSnssai ngapType.SNSSAI) (modelsSnssai models.Snssai) {
	modelsSnssai.Sst = int32(ngapSnssai.SST.Value[0])
	if ngapSnssai.SD != nil {
		modelsSnssai.Sd = hex.EncodeToString(ngapSnssai.SD.Value)
	}
	return
}

func SNssaiToNgap(modelsSnssai models.Snssai) (ngapSnssai ngapType.SNSSAI) {
	ngapSnssai.SST.Value = []byte{byte(modelsSnssai.Sst)}

	if modelsSnssai.Sd != "" {
		ngapSnssai.SD = new(ngapType.SD)
		ngapSnssai.SD.Value, _ = hex.DecodeString(modelsSnssai.Sd)
	}
	return
}
