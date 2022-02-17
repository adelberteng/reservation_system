package models

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/adelberteng/reservation_system/db"
	"github.com/adelberteng/reservation_system/utils"
)

var (
	logger = utils.GetLogger()
	engine = db.GetSQLClient()
)

type User struct {
	Uid          int    `xorm:"not null pk autoincr INT(11)"`
	Name         string `xorm:"not null VARCHAR(32)"`
	PasswordHash string `xorm:"not null VARCHAR(128)"`
	Phone        string `xorm:"not null CHAR(10)"`
	Email        string `xorm:"not null VARCHAR(255)"`
}

func init() {
	err := engine.Sync(new(User))
	if err != nil {
		logger.Error(err)
	}
}

func (u *User) TableName() string {
	return "user_tbl"
}

func hashingPassword(password string) (string, error) {
	pw := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(pw, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}

func verifyPassword(password, passwordHash string) bool {
	pw := []byte(password)
	hash := []byte(passwordHash)
	
	err := bcrypt.CompareHashAndPassword(hash, pw)
	if err != nil {
		return false
	} 

	return true
}

func RegisterUser(name, password, phone, email string) error {
	passwordHash, err := hashingPassword(password)
	if err != nil {
		return err
	}
	user := User{
		Name: name, 
		PasswordHash: passwordHash, 
		Phone: phone, 
		Email: email,
	}

	_, err = engine.Insert(&user)
	if err != nil {
		return err
	}

	return nil
}

		
