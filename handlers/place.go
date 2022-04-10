package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddPlace(c *gin.Context) {
	var json map[string]string
	c.ShouldBindJSON(&json)

	placeName := json["place_name"]
	region := json["region"]
	placeType := json["place_type"]
	address := json["address"]
	capacity := json["capacity"]
	if placeName == "" || region == "" || placeType == "" || address == "" || capacity == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "these fields can not be empty",
		})
		return
	}

	queryResult, err := engine.Table("place_tbl").Where("address = ? ", address).QueryString()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprint(err),
		})
		return
	}

	if queryResult != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "This place had been registered",
		})
		return
	}



}
