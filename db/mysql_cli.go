package db

import (
	"fmt"

	"github.com/adelberteng/reservation_system/utils"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	conf = utils.Conf
	logger           = utils.Logger
	dataSourceString string
	Engine *xorm.Engine
)

func init() {
	dataSourceString = fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8",
		conf.DBUser,
		conf.DBPassword,
		conf.DBEndpoint,
		conf.DBPort,
		conf.DBDatabase,
	)

	var err error

	Engine, err = xorm.NewEngine(conf.DBType, dataSourceString)
	if err != nil {
		logger.Error(err)
	}
}
