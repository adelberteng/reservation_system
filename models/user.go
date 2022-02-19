package models

import (
	// "fmt"
	"errors"
	"fmt"
	// "time"
	"golang.org/x/crypto/bcrypt"

	"github.com/adelberteng/reservation_system/db"
	"github.com/adelberteng/reservation_system/utils"
)

var (
	logger = utils.GetLogger()
	engine = db.GetSQLClient()
)

type User struct {
	Uid          string `xorm:"not null pk autoincr INT(11)"`
	Name         string `xorm:"not null unique VARCHAR(32)"`
	PasswordHash string `xorm:"not null VARCHAR(128)"`
	Phone        string `xorm:"not null unique CHAR(10)"`
	Email        string `xorm:"not null unique VARCHAR(255)"`
}

func init() {
	err := engine.Sync2(new(User))
	if err != nil {
		logger.Error(err)
	}
}

func (u *User) TableName() string {
	return "user_tbl"
}

func GeneratePasswordHash(password string) (string, error) {
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

func Register(name, password, phone, email string) error {
	nameResult, err := engine.Table("user_tbl").Where("name = ?", name).QueryString()
	if err != nil {
		return err
	} else if nameResult != nil {
		return errors.New("This user name had been registered")
	}

	phoneResult, err := engine.Table("user_tbl").Where("phone = ?", phone).QueryString()
	if err != nil {
		return err
	} else if phoneResult != nil {
		return errors.New("This phone number had been registered")
	}

	emailResult, err := engine.Table("user_tbl").Where("email = ?", email).QueryString()
	if err != nil {
		return err
	} else if emailResult != nil {
		return errors.New("This email address had been registered")
	}

	// generate the password hash with salt.
	passwordHash, err := GeneratePasswordHash(password)
	if err != nil {
		return err
	}
	user := User{Name: name, PasswordHash: passwordHash, Phone: phone, Email: email}

	_, err = engine.Insert(&user)
	if err != nil {
		return err
	}

	return nil
}

func Login(name, password string) (User, error) {
	res, err := engine.Table("user_tbl").Where("name = ?", name).QueryString()
	passwordHash := res[0]["password_hash"]
	if err != nil {
		fmt.Println(err)
	}

	isCorrect := verifyPassword(password, passwordHash)
	if !isCorrect {
		return User{}, errors.New("password is incorrect, please try again.")
	}

	user := User{
		Uid:   res[0]["uid"],
		Name:  res[0]["name"],
		Phone: res[0]["phone"],
		Email: res[0]["email"],
	}

	return user, nil
}

