package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/adelberteng/reservation_system/utils"
)

func VerifyToken(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	token, err := utils.ParseJWT(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message":    "can not parse token.",
		})
		c.Abort()
		return
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message":    "token is invalid.",
		})
		c.Abort()
		return
	}
	
	c.Next()
}