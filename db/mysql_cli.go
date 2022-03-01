package db

import (
	"fmt"

	"github.com/adelberteng/reservation_system/utils"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	conf = utils.Conf
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
	dataSourceString = fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8",
		conf.DB.User,
		conf.DB.Password,
		conf.DB.Endpoint,
		conf.DB.Port,
		conf.DB.DB,
	)
}

func GetSQLClient() *xorm.Engine {
	engine, err := xorm.NewEngine(conf.DB.Type, dataSourceString)
	if err != nil {
		logger.Error(err)
	}

	return engine
}
