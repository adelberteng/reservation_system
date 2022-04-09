package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AddPlace(c *gin.Context) {
	var json map[string]string
	c.ShouldBindJSON(&json)

	// placeName := json["place_name"]
	// region := json["region"]
	// placeType := json["place_type"]
	// address := json["address"]
	// capacity := json["capacity"]


}
