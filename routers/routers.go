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

	owner := router.Group("/owner")
	{
		owner.POST("/register", handlers.UserRegister)
		owner.POST("/login", handlers.UserLogin)
	}

	region := router.Group("/region")
	region.Use(middleware.VerifyToken)
	{
		region.POST("/add", handlers.AddRegion)
	}
	
	place := router.Group("/place")
	place.Use(middleware.VerifyToken)
	{
		place.POST("/add", handlers.AddPlace)
	}

	test := router.Group("/test")
	test.Use(middleware.VerifyToken)
	{
		test.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })
	}

	return router
}
