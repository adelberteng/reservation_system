package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/adelberteng/reservation_system/models"
)

func AddPlaceType(c *gin.Context) {
    var json map[string]string
	c.ShouldBindJSON(&json)

    typeName := json["type_name"]
	if typeName == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "these fields can not be empty",
		})
		return
	}

    placeType := models.PlaceType{TypeName: typeName}
	_, err := engine.Insert(&placeType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprint(err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "region adding success.",
	})


}