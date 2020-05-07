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

package ngapSctp

import (
	"encoding/binary"
	"errors"
	"net"
	"strings"
	"unsafe"

	"git.cs.nctu.edu.tw/calee/sctp"

	"radio_simulator/lib/ngap/logger"
)

var clientNum int
var NGAP_PPID uint32 = 60

func init() {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	switch buf {
	// little endian
	case [2]byte{0xCD, 0xAB}:
		tmp := make([]byte, 4)
		binary.BigEndian.PutUint32(tmp, NGAP_PPID)
		NGAP_PPID = binary.LittleEndian.Uint32(tmp)
	// big endian
	case [2]byte{0xAB, 0xCD}:
	}
}

// ConnData structure that pair the connection and data together
type ConnData struct {
	remoteAddr string
	data       []byte
	err        error
}

// GetError Return the Error of ConnData structure
func (cd ConnData) GetError() error {
	return cd.err
}

// GetRAddr Return the Remote Addr of ConnData structure
func (cd ConnData) GetRAddr() string {
	return cd.remoteAddr
}

// GetData Return the Data of ConnData structure
func (cd ConnData) GetData() []byte {
	return cd.data
}

// Server - Init SCTP Server, Set initial value / resource
func Server(addrStr string) *sctp.SCTPListener {
	port := 38412
	ips := []net.IPAddr{}

	for _, i := range strings.Split(addrStr, ",") {
		if a, err := net.ResolveIPAddr("ip", i); err == nil {
			logger.NgapLog.Debugf("Resolved address '%s' to %s\n", i, a)
			ips = append(ips, *a)
		} else {
			logger.NgapLog.Errorf("Error resolving address '%s': %v\n", i, err)
		}
	}

	addr := &sctp.SCTPAddr{
		IPAddrs: ips,
		Port:    port,
	}

	//ln, err := sctp.ListenSCTP("sctp", addr)
	ln, err := sctp.ListenSCTPExt("sctp", addr,
		sctp.InitMsg{NumOstreams: 3, MaxInstreams: 5, MaxAttempts: 4, MaxInitTimeout: 8})
	if err != nil {
		logger.NgapLog.Errorf("failed to listen: %v", err)
	}
	logger.NgapLog.Infof("Listen on %s", ln.Addr())

	return ln
}

// Accept - Accepting SCTP socket
func Accept(sctpLn *sctp.SCTPListener) (*sctp.SCTPConn, error) {
	sctpConn, err := sctpLn.AcceptSCTP()
	if err != nil {
		logger.NgapLog.Errorf("failed to accept: %v", err)
		return nil, err
	}
	info, _ := sctpConn.GetDefaultSentParam()
	info.PPID = NGAP_PPID
	err = sctpConn.SetDefaultSentParam(info)
	if err != nil {
		logger.NgapLog.Errorf("failed to accept: %v", err)
		return nil, err
	}
	err = sctpConn.SubscribeEvents(sctp.SCTP_EVENT_DATA_IO)
	if err != nil {
		logger.NgapLog.Errorf("failed to accept: %v", err)
		return nil, err
	}

	logger.NgapLog.Debugf("Accepted Connection from RemoteAddr: %s", sctpConn.RemoteAddr())

	// wconn := sctp.NewSCTPSndRcvInfoWrappedConn(conn.(*sctp.SCTPConn))
	clientNum++
	logger.NgapLog.Debugf("A new Connection %d.\n", clientNum)

	return sctpConn, nil
}

// Start - Start SCTP read channel
func Start(conn *sctp.SCTPConn, readChan chan ConnData) {
	defer closeConnection(conn)
	raddr := conn.RemoteAddr()
	if raddr == nil {
		// conn error
		return
	}
	raddrStr := raddr.String()
	for {
		buffer := make([]byte, 8192)
		n, info, err := conn.SCTPRead(buffer)
		if err != nil {
			logger.NgapLog.Debugf("Error %v", err)
			readChan <- ConnData{remoteAddr: raddrStr, data: nil, err: err}
			break
		} else if info == nil || info.PPID != NGAP_PPID {
			logger.NgapLog.Warnf("Recv SCTP PPID != 60")
			continue
		}
		logger.NgapLog.Debugf("Read: %s, %s, %x", raddrStr, string(buffer[:n]), buffer[:n])
		readChan <- ConnData{remoteAddr: raddrStr, data: buffer[:n], err: nil}
	}
}

// SendMsg - used to send out message to SCTP connection
func SendMsg(conn net.Conn, msg []byte) error {
	if conn.RemoteAddr() == nil {
		// conn error
		return errors.New("Connection no Remote Address")
	}
	logger.NgapLog.Debugf("Write: %s, %s, %x", conn.RemoteAddr().String(), string(msg), msg)
	_, err := conn.Write(msg)
	if err != nil {
		logger.NgapLog.Errorf("Error %v", err)
		return err
	}
	return nil
}

func closeConnection(conn net.Conn) {

	conn.Close()
	clientNum--
	logger.NgapLog.Debugf("Now, %d connections is alive.\n", clientNum)

}

// Destroy - Destroy the SCTP Server Resource
func Destroy(ln *sctp.SCTPListener) error {
	logger.NgapLog.Infoln("Close listener")
	return ln.Close()
}
