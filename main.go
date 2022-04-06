package main

import (
	"github.com/adelberteng/reservation_system/routers"
	"github.com/adelberteng/reservation_system/utils"
)

var (
	log  = utils.Logger
	appConf = utils.AppConf
)

func main() {
	log.Info("logrus")
	router := routers.SetupRoute()
	router.Run(":" + appConf.Port)
}
