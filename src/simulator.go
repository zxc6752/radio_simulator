package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"radio_simulator/lib/MongoDBLibrary"
	"radio_simulator/lib/path_util"
	"radio_simulator/src/factory"
	"radio_simulator/src/logger"
	"radio_simulator/src/simulator_context"
	"radio_simulator/src/simulator_init"
	"radio_simulator/src/simulator_util"
	"radio_simulator/src/tcp_server"
	"syscall"
)

var config string

var self *simulator_context.Simulator = simulator_context.Simulator_Self()

func Initailize() {
	flag.StringVar(&config, "simcfg", path_util.Gofree5gcPath("radio_simulator/config/rancfg.conf"), "ran simulator config file")
	flag.Parse()

	factory.InitConfigFactory(config)
	config := factory.SimConfig
	if config.Logger.DebugLevel != "" {
		level, err := logrus.ParseLevel(config.Logger.DebugLevel)
		if err == nil {
			logger.SetLogLevel(level)
		}
	}
	logger.SetReportCaller(config.Logger.ReportCaller)

	MongoDBLibrary.SetMongoDB(config.DBName, config.DBUrl)
}

func Terminate() {
	logger.SimulatorLog.Infof("Terminating Simulator...")

	// TODO: Send UE Deregistration to AMF
	logger.SimulatorLog.Infof("Clear UE DB...")

	simulator_util.ClearDB()

	logger.SimulatorLog.Infof("Close SCTP Connection...")

	for _, ran := range self.RanPool {
		logger.SimulatorLog.Infof("Ran[%s] Connection Close", ran.RanSctpUri)
		ran.SctpConn.Close()
	}

	logger.SimulatorLog.Infof("Close TCP Server...")

	if self.TcpServer != nil {
		self.TcpServer.Close()
	}

	logger.SimulatorLog.Infof("Clean Ue IP Addr in IP tables")

	// for key, conn := range self.GtpConnPool {
	// 	logger.InitLog.Infof("GTP[%s] Connection Close", key)
	// 	conn.Close()
	// }
	for _, ue := range self.UeContextPool {
		for _, sess := range ue.PduSession {
			if sess.UeIp != "" {
				_, err := exec.Command("ip", "addr", "del", sess.UeIp, "dev", "lo").Output()
				if err != nil {
					logger.SimulatorLog.Errorf("Delete ue addr failed[%s]", err.Error())
				}
			}
		}
	}

	// logger.SimulatorLog.Infof("Close Raw Socket...")
	// if self.ListenRawConn != nil {
	// 	self.ListenRawConn.Close()
	// }

	logger.SimulatorLog.Infof("Simulator terminated")

}

func main() {
	Initailize()
	simulator_util.ParseRanContext()
	simulator_util.ParseTunDev()

	path, err := filepath.Abs(filepath.Dir(config))
	if err != nil {
		logger.SimulatorLog.Errorf(err.Error())
	}
	simulator_util.ParseUeData(path+"/", factory.SimConfig.UeInfoFile)
	simulator_util.InitUeToDB()
	for _, ran := range self.RanPool {
		simulator_init.RanStart(ran)
	}

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChannel
		Terminate()
		os.Exit(0)
	}()
	// Raw Socket Server
	// self.ListenRawConn = raw_socket.ListenRawSocket(factory.SimConfig.ListenIp)
	// TCP server for cli test UE
	tcp_server.StartTcpServer()
	simulator_util.ClearDB()
}
