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
	"radio_simulator/lib/openapi/models"
)

func LadnToModels(buf []uint8) (dnnValues []string) {

	for bufOffset := 1; bufOffset < len(buf); {
		lenOfDnn := int(buf[bufOffset])
		dnn := string(buf[bufOffset : bufOffset+lenOfDnn])
		dnnValues = append(dnnValues, dnn)
		bufOffset += lenOfDnn
	}

	return
}

func LadnToNas(dnn string, taiLists []models.Tai) (ladnNas []uint8) {

	dnnNas := []byte(dnn)

	ladnNas = append(ladnNas, uint8(len(dnnNas)))
	ladnNas = append(ladnNas, dnnNas...)

	taiListNas := TaiListToNas(taiLists)
	ladnNas = append(ladnNas, uint8(len(taiListNas)))
	ladnNas = append(ladnNas, taiListNas...)
	return
}
