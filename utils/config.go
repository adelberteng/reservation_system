package utils

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type Config struct {
	DB
	App
	Log
}

type DB struct {
	Type string
	Endpoint string
	Port string
	User string
	Password string
	DB string
}

type App struct {
	Port string
	SecretKey string
}

type Log struct {
	LogDir string
	LogFileName string
	Level string
}

var Conf Config

func init() {
	cfg, err := ini.Load("conf/config.ini")
	if err != nil {
		fmt.Println("error happened when load config.")
		os.Exit(1)
	}

	Conf.DB.Type = cfg.Section("db").Key("db_type").String()
	Conf.DB.Endpoint = cfg.Section("db").Key("endpoint").String()
	Conf.DB.Port = cfg.Section("db").Key("port").String()
	Conf.DB.User = cfg.Section("db").Key("user").String()
	Conf.DB.Password = cfg.Section("db").Key("password").String()
	Conf.DB.DB = cfg.Section("db").Key("db").String()

	Conf.App.Port = cfg.Section("app").Key("service_port").String()
	Conf.App.SecretKey = cfg.Section("app").Key("secret_key").String()

	Conf.Log.LogDir = cfg.Section("log").Key("log_dir").String()
	Conf.Log.LogFileName = cfg.Section("log").Key("log_file_name").String()
	Conf.Log.Level = cfg.Section("log").Key("log_level").String()
}