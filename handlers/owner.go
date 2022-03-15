package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/adelberteng/reservation_system/models"
	"github.com/adelberteng/reservation_system/utils"
)

func OwnerRegister(c *gin.Context) {
	var json map[string]string
	c.ShouldBindJSON(&json)

	companyName := json["company_name"]
	password := json["password"]
	phone := json["phone"]
	email := json["email"]
	if companyName == "" || password == "" || phone == "" || email == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "these fields can not be empty",
		})
		return
	}

	queryResult, err := engine.Table("owner_tbl").Where(
		"comany_name = ? or phone = ? or email = ? ", companyName, phone, email).QueryString()

	if queryResult != nil {
		var errMessage string
		for _, row := range queryResult {
			if row["comany_name"] == companyName {
				errMessage = "This company_name had been registered"
				break
			} else if row["phone"] == phone {
				errMessage = "This phone number had been registered"
				break
			} else if row["email"] == email {
				errMessage = "This email address had been registered"
				break
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"message": errMessage,
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprint(err),
		})
		return
	}

	passwordHash, err := utils.GeneratePasswordHash(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprint(err),
		})
		return
	}

	owner := models.Owner{CompanyName: companyName, PasswordHash: passwordHash, Phone: phone, Email: email}

	_, err = engine.Insert(&owner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprint(err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "owner register success.",
	})
}

func OwnerLogin(c *gin.Context) {
	var json map[string]string
	c.ShouldBindJSON(&json)

	companyName := json["company_name"]
	password := json["password"]

	owner, err := models.GetOwnerByName(companyName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprint(err),
		})
		return
	}

	if !utils.VerifyPassword(password, owner.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "the password is incorrect.",
		})
		return
	}

	jwtPayload := map[string]string{
		"id":           owner.Id,
		"company_name": owner.CompanyName,
	}
	jwt, err := utils.GenerateJWT(jwtPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprint(err),
		})
		return
	}

	c.Header("Authorization", jwt)
	c.JSON(http.StatusOK, gin.H{
		"message": "login success.",
	})
}
