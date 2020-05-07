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

type PDUSessionReleaseCommand struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.PDUSESSIONRELEASECOMMANDMessageIdentity
	nasType.Cause5GSM
	*nasType.BackoffTimerValue
	*nasType.EAPMessage
	*nasType.ExtendedProtocolConfigurationOptions
}

func NewPDUSessionReleaseCommand(iei uint8) (pDUSessionReleaseCommand *PDUSessionReleaseCommand) {
	pDUSessionReleaseCommand = &PDUSessionReleaseCommand{}
	return pDUSessionReleaseCommand
}

const (
	PDUSessionReleaseCommandBackoffTimerValueType                    uint8 = 0x37
	PDUSessionReleaseCommandEAPMessageType                           uint8 = 0x78
	PDUSessionReleaseCommandExtendedProtocolConfigurationOptionsType uint8 = 0x7B
)

func (a *PDUSessionReleaseCommand) EncodePDUSessionReleaseCommand(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Write(buffer, binary.BigEndian, &a.PDUSESSIONRELEASECOMMANDMessageIdentity.Octet)
	binary.Write(buffer, binary.BigEndian, &a.Cause5GSM.Octet)
	if a.BackoffTimerValue != nil {
		binary.Write(buffer, binary.BigEndian, a.BackoffTimerValue.GetIei())
		binary.Write(buffer, binary.BigEndian, a.BackoffTimerValue.GetLen())
		binary.Write(buffer, binary.BigEndian, &a.BackoffTimerValue.Octet)
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

func (a *PDUSessionReleaseCommand) DecodePDUSessionReleaseCommand(byteArray *[]byte) {
	buffer := bytes.NewBuffer(*byteArray)
	binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PTI.Octet)
	binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONRELEASECOMMANDMessageIdentity.Octet)
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
		case PDUSessionReleaseCommandBackoffTimerValueType:
			a.BackoffTimerValue = nasType.NewBackoffTimerValue(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Len)
			a.BackoffTimerValue.SetLen(a.BackoffTimerValue.GetLen())
			binary.Read(buffer, binary.BigEndian, &a.BackoffTimerValue.Octet)
		case PDUSessionReleaseCommandEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len)
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer[:a.EAPMessage.GetLen()])
		case PDUSessionReleaseCommandExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len)
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer[:a.ExtendedProtocolConfigurationOptions.GetLen()])
		default:
		}
	}
}
