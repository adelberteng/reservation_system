package utils

import (
	"fmt"
	"os"
	"strings"
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
		viper.SetEnvKeyReplacer(strings.NewReplacer("_", ""))

		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("fatal error config file: %s \n", err)
			os.Exit(1)
		}

		viper.Unmarshal(&Config)
	})
}
