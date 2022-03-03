package utils

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type DBConfig struct {
	Type string
	Endpoint string
	Port string
	User string
	Password string
	Database string
}

type AppConfig struct {
	Port string
	SecretKey string
}

type LoggerConfig struct {
	Dir string
	FileName string
	Level string
}

var (
	DBConf DBConfig
	AppConf AppConfig
	LoggerConf LoggerConfig
)

func init() {
	cfg, err := ini.Load("conf/config.ini")
	if err != nil {
		fmt.Println("error happened when load config.")
		os.Exit(1)
	}

	DBConf.Type = cfg.Section("db").Key("db_type").String()
	DBConf.Endpoint = cfg.Section("db").Key("endpoint").String()
	DBConf.Port = cfg.Section("db").Key("port").String()
	DBConf.User = cfg.Section("db").Key("user").String()
	DBConf.Password = cfg.Section("db").Key("password").String()
	DBConf.Database = cfg.Section("db").Key("db").String()

	AppConf.Port = cfg.Section("app").Key("service_port").String()
	AppConf.SecretKey = cfg.Section("app").Key("secret_key").String()

	LoggerConf.Dir = cfg.Section("log").Key("log_dir").String()
	LoggerConf.FileName = cfg.Section("log").Key("log_file_name").String()
	LoggerConf.Level = cfg.Section("log").Key("log_level").String()
}