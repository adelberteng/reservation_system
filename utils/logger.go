package utils

import (
	"os"
	"io"
	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func init() {
	logDir := LoggerConf.Dir
	logName := LoggerConf.FileName
	logLevel := LoggerConf.Level

	os.MkdirAll(logDir, 0766)
	logFile, err := os.OpenFile(logDir+"/"+logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if  err != nil {
		Logger.Fatalf("log file open error : %v", err)
	}

	Logger.Out = io.MultiWriter(os.Stdout, logFile)
	Logger.Formatter = &logrus.JSONFormatter{}
	logrusLevel, err := logrus.ParseLevel(logLevel)
	if err != nil {
		Logger.Fatalf("log level error : %v", err)
	}
	Logger.Level = logrusLevel
}
