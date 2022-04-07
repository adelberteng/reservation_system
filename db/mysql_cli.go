package db

import (
	"fmt"

	"github.com/adelberteng/reservation_system/utils"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	conf             = utils.Config.DB
	logger           = utils.Logger
	dataSourceString string
	Engine           *xorm.Engine
)

func init() {
	dataSourceString = fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8",
		conf.User,
		conf.Password,
		conf.Endpoint,
		conf.Port,
		conf.Database,
	)

	var err error

	Engine, err = xorm.NewEngine(conf.Type, dataSourceString)
	if err != nil {
		logger.Error(err)
	}
}
