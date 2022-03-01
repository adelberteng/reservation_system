package main

import (
	"github.com/adelberteng/reservation_system/routers"
	"github.com/adelberteng/reservation_system/utils"
)

var (
	logger = utils.GetLogger()
)

func main() {
	router := routers.SetupRoute()
	router.Run()
}
