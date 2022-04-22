package handlers

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/adelberteng/reservation_system/models"
)

func AddRegion(c *gin.Context) {
    var json map[string]string
	c.ShouldBindJSON(&json)

    regionName := json["region_name"]
	if regionName == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "the field can not be empty",
		})
		return
	}

	region := models.Region{RegionName: regionName}
	_, err := engine.Insert(&region)
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