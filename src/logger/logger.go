package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
)

var log *logrus.Logger
var SimulatorLog *logrus.Entry
var UtilLog *logrus.Entry
var InitLog *logrus.Entry
var NgapLog *logrus.Entry
var GtpLog *logrus.Entry
var HandlerLog *logrus.Entry
var NasLog *logrus.Entry
var TcpServerLog *logrus.Entry
var ContextLog *logrus.Entry

func init() {
	log = logrus.New()
	log.SetReportCaller(true)

	log.Formatter = &logrus.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           "",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			orgFilename, _ := os.Getwd()
			repopath := orgFilename
			repopath = strings.Replace(repopath, "/bin", "", 1)
			filename := strings.Replace(f.File, repopath, "", -1)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}

	SimulatorLog = log.WithFields(logrus.Fields{"SimulatorLog": "SimulatorLog"})
	NgapLog = log.WithFields(logrus.Fields{"SimulatorLog": "NgapLog"})
	GtpLog = log.WithFields(logrus.Fields{"SimulatorLog": "GtpLog"})
	HandlerLog = log.WithFields(logrus.Fields{"SimulatorLog": "HandlerLog"})
	InitLog = log.WithFields(logrus.Fields{"SimulatorLog": "InitLog"})
	UtilLog = log.WithFields(logrus.Fields{"SimulatorLog": "UtilLog"})
	NasLog = log.WithFields(logrus.Fields{"SimulatorLog": "NasLog"})
	TcpServerLog = log.WithFields(logrus.Fields{"SimulatorLog": "TcpServerLog"})
	ContextLog = log.WithFields(logrus.Fields{"SimulatorLog": "ContextLog"})
}

func SetLogLevel(level logrus.Level) {
	SimulatorLog.Infoln("set log level :", level)
	log.SetLevel(level)
}

func SetReportCaller(bool bool) {
	SimulatorLog.Infoln("set report call :", bool)
	log.SetReportCaller(bool)
}
