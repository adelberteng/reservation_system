package main

import (
	"fmt"

	"github.com/adelberteng/reservation_system/models"
	// "github.com/adelberteng/reservation_system/utils"
)

func main() {
	err := models.Register("aa", "123456", "0912345677", "aa@gmail.com")
	if err != nil {
		fmt.Println(err)
	}

	// user, err := models.Login("aaa", "123456")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%+v \n", user)

	// payload := map[string]string{"name": "aaaa"}
	// s, err := utils.GenerateJWT(payload)
	// if err != nil {
	// 	fmt.Println(err)
	// }


	// p, err := utils.ParseJWT(s)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%+v \n", p)
}
