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

package ngap

import (
	"fmt"
	"radio_simulator/lib/aper"
	"radio_simulator/lib/ngap/ngapType"
	"reflect"
	"strings"
)

// Decoder is to decode raw data to NGAP pdu pointer with PER Aligned
func Decoder(b []byte) (pdu *ngapType.NGAPPDU, err error) {
	pdu = &ngapType.NGAPPDU{}

	err = aper.UnmarshalWithParams(b, pdu, "valueExt,valueLB:0,valueUB:2")
	return
}

// Encoder is to NGAP pdu to raw data with PER Aligned
func Encoder(pdu ngapType.NGAPPDU) ([]byte, error) {
	return aper.MarshalWithParams(pdu, "valueExt,valueLB:0,valueUB:2")
}

func PrintResult(v reflect.Value, layer int) (s string) {

	fieldType := v.Type()
	if v.Kind() == reflect.Ptr {
		s += ("&" + PrintResult(v.Elem(), layer))
		return
	}
	switch fieldType {
	case aper.OctetStringType:
		s = fmt.Sprintf("OctetString (0x%x)[%d]\n", v.Bytes(), len(v.Bytes()))
		return

	case aper.BitStringType:
		s = fmt.Sprintf("BitString (%0.8b)[%d]\n", v.Field(0).Bytes(), v.Field(1).Uint())
		return

	case aper.EnumeratedType:
		s = fmt.Sprintf("Enumerated(%d)\n", v.Uint())
		return
	}
	switch v.Kind() {
	case reflect.Struct:
		structType := fieldType
		s += fmt.Sprintf("{\n")
		end := strings.Repeat(" ", layer) + "}\n"
		layer += 2
		space := strings.Repeat(" ", layer)
		if structType.Field(0).Name == "Present" {
			present := v.Field(0).Int()
			s += (space + fmt.Sprintf("Present: %d\n", present))
			s += (space + fmt.Sprintf("%s : ", structType.Field(int(present)).Name))
			s += PrintResult(v.Field(int(present)), layer)
			s += end
			return

		}
		for i := 0; i < v.NumField(); i++ {
			// optional
			if v.Field(i).Type().Kind() == reflect.Ptr && v.Field(i).IsNil() {
				continue
			}

			s += (space + fmt.Sprintf("%s : ", structType.Field(i).Name))
			s += PrintResult(v.Field(i), layer)
		}
		s += end
	case reflect.Slice:
		s += fmt.Sprintf("[%d]{\n", v.Len())
		end := strings.Repeat(" ", layer) + "}\n"
		layer += 2
		space := strings.Repeat(" ", layer)
		for i := 0; i < v.Len(); i++ {
			s += space
			s += PrintResult(v.Index(i), layer)
			s += (space + ",\n")
		}
		s += end
	case reflect.String:
		s = fmt.Sprintf("PrintableString(\"%s\")\n", v.String())
	case reflect.Int32, reflect.Int64:
		s = fmt.Sprintf("INTEGER(%d)\n", v.Int())
	default:
		fmt.Printf("Type: %s does not handle", v.Type())

	}
	return
}
