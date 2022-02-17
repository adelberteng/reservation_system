package db

import (
	"fmt"

	"github.com/adelberteng/reservation_system/utils"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	cfg              = utils.GetConfig()
	logger           = utils.GetLogger()
	dbType           string
	dbEndpoint       string
	dbPort           string
	dbUser           string
	dbPassword       string
	dbDB             string
	dataSourceString string
)

func init() {
	dbType = cfg.Section("db").Key("db_type").String()
	dbEndpoint = cfg.Section("db").Key("endpoint").String()
	dbPort = cfg.Section("db").Key("port").String()
	dbUser = cfg.Section("db").Key("user").String()
	dbPassword = cfg.Section("db").Key("password").String()
	dbDB = cfg.Section("db").Key("db").String()
	dataSourceString = fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8",
		dbUser,
		dbPassword,
		dbEndpoint,
		dbPort,
		dbDB,
	)
}

func GetSQLClient() *xorm.Engine {
	engine, err := xorm.NewEngine(dbType, dataSourceString)
	if err != nil {
		logger.Error(err)
	}

	return engine
}
