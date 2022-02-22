package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/adelberteng/reservation_system/handlers"
)

func SetupRoute() *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("/signup", handlers.Register)
		user.POST("/login", handlers.Login)
	}

	return router
}
