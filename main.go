package main

import (
	"github.com/adelberteng/reservation_system/routers"
	"github.com/adelberteng/reservation_system/utils"
)

var (
	cfg    = utils.GetConfig()
	logger = utils.GetLogger()
)

func main() {
	router := routers.SetupRoute()
	router.Run()
}
