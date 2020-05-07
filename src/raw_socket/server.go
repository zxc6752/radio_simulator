package raw_socket

import (
// "encoding/binary"
// "fmt"
// "golang.org/x/net/ipv4"
// "net"
// "radio_simulator/src/logger"
// "radio_simulator/src/simulator_context"
// "syscall"
)

// var self *simulator_context.Simulator = simulator_context.Simulator_Self()

// func ListenRawSocket(ip string) *ipv4.RawConn {
// 	c, err := net.ListenPacket(fmt.Sprintf("ip4:%d", syscall.IPPROTO_IPIP), ip)
// 	if err != nil {
// 		logger.GtpLog.Errorf("Error: %+v", err.Error())
// 		return nil
// 	}

// 	conn, err := ipv4.NewRawConn(c)
// 	if err != nil {
// 		logger.GtpLog.Errorf("Error: %+v", err.Error())
// 		return nil
// 	}
// 	go serveRawSocket(conn)
// 	return conn
// }

// func serveRawSocket(conn *ipv4.RawConn) {
// 	defer func() {
// 		if conn != nil {
// 			conn.Close()
// 		}
// 	}()

// 	buf := make([]byte, 8096)
// 	for {
// 		n, err := conn.Read(buf)
// 		if err != nil {
// 			logger.GtpLog.Debugf("Error: %+v", err.Error())
// 			return
// 		}
// 		payload := buf[20:n]
// 		ip4header, err := ipv4.ParseHeader(payload[:20])
// 		if err != nil {
// 			logger.GtpLog.Errorf(err.Error())
// 			continue
// 		}
// 		ueIp := ip4header.Src.String()
// 		if sess, exist := self.SessPool[ueIp]; exist {
// 			binary.BigEndian.PutUint16(sess.GtpHdr[2:4], sess.GtpHdrLen-1+uint16(n)-20)
// 			conn, err := sess.GetGtpConn()
// 			if err != nil {
// 				logger.GtpLog.Errorf(err.Error())
// 				continue
// 			}
// 			conn.Write(append(sess.GtpHdr, payload...))
// 		}
// 	}
// }
