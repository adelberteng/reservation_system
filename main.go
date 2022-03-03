package main

import (
	"github.com/adelberteng/reservation_system/routers"
	"github.com/adelberteng/reservation_system/utils"
)

var (
	logger = utils.Logger
	appConf = utils.AppConf
)

func main() {
	router := routers.SetupRoute()
	router.Run(appConf.Port)
}
