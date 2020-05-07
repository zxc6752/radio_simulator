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

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"radio_simulator/lib/nas/nasType"
)

type PDUSessionEstablishmentReject struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.PDUSESSIONESTABLISHMENTREJECTMessageIdentity
	nasType.Cause5GSM
	*nasType.BackoffTimerValue
	*nasType.AllowedSSCMode
	*nasType.EAPMessage
	*nasType.ExtendedProtocolConfigurationOptions
}

func NewPDUSessionEstablishmentReject(iei uint8) (pDUSessionEstablishmentReject *PDUSessionEstablishmentReject) {
	pDUSessionEstablishmentReject = &PDUSessionEstablishmentReject{}
	return pDUSessionEstablishmentReject
}

const (
	PDUSessionEstablishmentRejectBackoffTimerValueType                    uint8 = 0x37
	PDUSessionEstablishmentRejectAllowedSSCModeType                       uint8 = 0x0F
	PDUSessionEstablishmentRejectEAPMessageType                           uint8 = 0x78
	PDUSessionEstablishmentRejectExtendedProtocolConfigurationOptionsType uint8 = 0x7B
)

func (a *PDUSessionEstablishmentReject) EncodePDUSessionEstablishmentReject(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSESSIONESTABLISHMENTREJECTMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.Cause5GSM.Octet)
	if a.BackoffTimerValue != nil {
		binary.Write(buffer, binary.BigEndian, a.BackoffTimerValue.GetIei())
		binary.Write(buffer, binary.BigEndian, a.BackoffTimerValue.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.BackoffTimerValue.Octet)
	}
	if a.AllowedSSCMode != nil {
		binary.Write(buffer, binary.BigEndian, &a.AllowedSSCMode.Octet)
	}
	if a.EAPMessage != nil {
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei())
		binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.EAPMessage.Buffer)
	}
	if a.ExtendedProtocolConfigurationOptions != nil {
		binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetIei())
		binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Buffer)
	}
}

func (a *PDUSessionEstablishmentReject) DecodePDUSessionEstablishmentReject(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONESTABLISHMENTREJECTMessageIdentity.Octet)
	binary.Read(buffer, binary.BigEndian, &a.Cause5GSM.Octet)
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		binary.Read(buffer, binary.BigEndian, &ieiN)
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case PDUSessionEstablishmentRejectBackoffTimerValueType:
			a.BackoffTimerValue = nasType.NewBackoffTimerValue(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Len)
			a.BackoffTimerValue.SetLen(a.BackoffTimerValue.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Octet)
		case PDUSessionEstablishmentRejectAllowedSSCModeType:
			a.AllowedSSCMode = nasType.NewAllowedSSCMode(ieiN)
			a.AllowedSSCMode.Octet = ieiN
		case PDUSessionEstablishmentRejectEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()])
		case PDUSessionEstablishmentRejectExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len)
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer[:a.ExtendedProtocolConfigurationOptions.GetLen()])
		default:
		}
	}
}
