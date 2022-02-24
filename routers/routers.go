package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/adelberteng/reservation_system/handlers"
	"github.com/adelberteng/reservation_system/middleware"
)

func SetupRoute() *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("/register", handlers.UserRegister)
		user.POST("/login", handlers.UserLogin)
	}
	
	router.GET("/ping", middleware.VerifyToken, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
