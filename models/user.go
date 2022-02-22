package models

import (
	"errors"
	"fmt"

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

func Register(name, password, phone, email string) error {
	queryResult, err := engine.Table("user_tbl").Where(
		"name = ? and phone = ? and email = ? ", name, phone, email).QueryString()

	var record map[string]string
	if queryResult != nil {
		record = queryResult[0]
	}

	if record["name"] == name {
		return errors.New("This user name had been registered")
	} else if record["phone"] == phone {
		return errors.New("This phone number had been registered")
	} else if record["email"] == email {
		return errors.New("This email address had been registered")
	} else if err != nil {
		return err
	}

	passwordHash, err := utils.GeneratePasswordHash(password)
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

	isCorrect := utils.VerifyPassword(password, passwordHash)
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
