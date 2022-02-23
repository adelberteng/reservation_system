package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/adelberteng/reservation_system/utils"
)

func Verify(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	token, err := utils.ParseJWT(tokenStr)
	if err != nil {
		c.Set("verify_status", "can not parse token.")
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if !token.Valid {
		c.Set("verify_status", "token is invalid.")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	
	c.Next()
}