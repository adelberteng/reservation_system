package utils

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type config struct {
	DB  *db
	App *app
	Log *log
}

type db struct {
	Type     string
	Endpoint string
	Port     string
	User     string
	Password string
	Database string
}

type app struct {
	Port      string
	SecretKey string
}

type log struct {
	Dir      string
	FileName string
	Level    string
}

var (
	once   sync.Once
	DB     db
	App    app
	Log    log
	Config = &config{
		DB:  &DB,
		App: &App,
		Log: &Log,
	}
)

func init() {
	once.Do(func() {
		viper.SetConfigFile("./conf/config.yaml")
		viper.SetConfigType("yaml")

		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("Fatal error config file: %w \n", err))
		}

		viper.Unmarshal(&Config)
	})
}
