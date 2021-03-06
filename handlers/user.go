package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/adelberteng/reservation_system/models"
	"github.com/adelberteng/reservation_system/utils"
)

func UserRegister(c *gin.Context) {
	var json map[string]string
	c.ShouldBindJSON(&json)

	name := json["name"]
	password := json["password"]
	phone := json["phone"]
	email := json["email"]
	if name == "" || password == "" || phone == "" || email == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "these fields can not be empty",
		})
		return
	}

	queryResult, err := engine.Table("user_tbl").Where(
		"name = ? or phone = ? or email = ? ", name, phone, email).QueryString()

	if queryResult != nil {
		var errMessage string
		for _, row := range queryResult {
			if row["name"] == name {
				errMessage = "This user name had been registered"
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

	user := models.User{Name: name, PasswordHash: passwordHash, Phone: phone, Email: email}

	_, err = engine.Insert(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprint(err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user register success.",
	})
}

func UserLogin(c *gin.Context) {
	var json map[string]string
	c.ShouldBindJSON(&json)

	name := json["name"]
	password := json["password"]

	user, err := models.GetUserByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprint(err),
		})
		return
	}

	if !utils.VerifyPassword(password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "the password is incorrect.",
		})
		return
	}

	jwtPayload := map[string]string{
		"uid":  user.Uid,
		"name": user.Name,
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
