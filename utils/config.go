package utils

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type Config struct {
	DBType string
	DBEndpoint string
	DBPort string
	DBUser string
	DBPassword string
	DBDatabase string
	AppPort string
	AppSecretKey string
	LogDir string
	LogFileName string
	LogLevel string
}

var Conf Config

func init() {
	cfg, err := ini.Load("conf/config.ini")
	if err != nil {
		fmt.Println("error happened when load config.")
		os.Exit(1)
	}

	Conf.DBType = cfg.Section("db").Key("db_type").String()
	Conf.DBEndpoint = cfg.Section("db").Key("endpoint").String()
	Conf.DBPort = cfg.Section("db").Key("port").String()
	Conf.DBUser = cfg.Section("db").Key("user").String()
	Conf.DBPassword = cfg.Section("db").Key("password").String()
	Conf.DBDatabase = cfg.Section("db").Key("db").String()

	Conf.AppPort = cfg.Section("app").Key("service_port").String()
	Conf.AppSecretKey = cfg.Section("app").Key("secret_key").String()

	Conf.LogDir = cfg.Section("log").Key("log_dir").String()
	Conf.LogFileName = cfg.Section("log").Key("log_file_name").String()
	Conf.LogLevel = cfg.Section("log").Key("log_level").String()
}