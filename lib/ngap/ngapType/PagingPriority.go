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

package ngapType

import "radio_simulator/lib/aper"

// Need to import "free5gc/lib/aper" if it uses "aper"

const (
	PagingPriorityPresentPriolevel1 aper.Enumerated = 0
	PagingPriorityPresentPriolevel2 aper.Enumerated = 1
	PagingPriorityPresentPriolevel3 aper.Enumerated = 2
	PagingPriorityPresentPriolevel4 aper.Enumerated = 3
	PagingPriorityPresentPriolevel5 aper.Enumerated = 4
	PagingPriorityPresentPriolevel6 aper.Enumerated = 5
	PagingPriorityPresentPriolevel7 aper.Enumerated = 6
	PagingPriorityPresentPriolevel8 aper.Enumerated = 7
)

type PagingPriority struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:7"`
}
