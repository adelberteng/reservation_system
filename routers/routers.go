package routers

import (
	"github.com/gin-gonic/gin"

	handlers "github.com/adelberteng/reservation_system/handlers"
)

func SetupRoute() *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("/register", handlers.UserRegister)
		user.POST("/login", handlers.UserLogin)
	}

	return router
}
