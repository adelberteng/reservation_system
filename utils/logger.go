package utils

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var Logger logrus.Logger

func init() {
	logDir := Config.Log.Dir
	logName := Config.Log.FileName
	logLevel := Config.Log.Level

	once.Do(func() {
		Logger = *logrus.New()
		os.MkdirAll(logDir, 0766)
		logFile, err := os.OpenFile(logDir+"/"+logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			Logger.Fatalf("log file open error : %v", err)
		}
		Logger.Out = io.MultiWriter(os.Stdout, logFile)
		Logger.Formatter = &logrus.JSONFormatter{}
		logrusLevel, err := logrus.ParseLevel(logLevel)
		if err != nil {
			Logger.Fatalf("log level error : %v", err)
		}
		Logger.Level = logrusLevel
	})
}
