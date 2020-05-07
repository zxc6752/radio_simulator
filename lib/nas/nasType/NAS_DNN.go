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

package nasType

import (
	"radio_simulator/lib/util_3gpp"
)

// DNN 9.11.2.1A
// DNN Row, sBit, len = [0, 0], 8 , INF
type DNN struct {
	Iei    uint8
	Len    uint8
	Buffer []uint8
}

func NewDNN(iei uint8) (dNN *DNN) {
	dNN = &DNN{}
	dNN.SetIei(iei)
	return dNN
}

// DNN 9.11.2.1A
// Iei Row, sBit, len = [], 8, 8
func (a *DNN) GetIei() (iei uint8) {
	return a.Iei
}

// DNN 9.11.2.1A
// Iei Row, sBit, len = [], 8, 8
func (a *DNN) SetIei(iei uint8) {
	a.Iei = iei
}

// DNN 9.11.2.1A
// Len Row, sBit, len = [], 8, 8
func (a *DNN) GetLen() (len uint8) {
	return a.Len
}

// DNN 9.11.2.1A
// Len Row, sBit, len = [], 8, 8
func (a *DNN) SetLen(len uint8) {
	a.Len = len
	a.Buffer = make([]uint8, a.Len)
}

// DNN 9.11.2.1A
// DNN Row, sBit, len = [0, 0], 8 , INF
func (a *DNN) GetDNN() (dNN []uint8) {
	dnn := new(util_3gpp.Dnn)
	dnn.UnmarshalBinary(a.Buffer)
	return *dnn
}

// DNN 9.11.2.1A
// DNN Row, sBit, len = [0, 0], 8 , INF
func (a *DNN) SetDNN(dNN []uint8) {
	tmp := (util_3gpp.Dnn)(dNN)
	dnn, _ := tmp.MarshalBinary()
	a.Buffer = dnn
	a.Len = uint8(len(a.Buffer))
}
