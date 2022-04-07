package main

import (
	"github.com/adelberteng/reservation_system/routers"
	"github.com/adelberteng/reservation_system/utils"
)

var (
	log  = utils.Logger
	appConf = utils.Config.App
)

func main() {
	router := routers.SetupRoute()
	router.Run(":" + appConf.Port)
}
