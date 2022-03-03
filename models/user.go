package models

import (
	"errors"

	"github.com/adelberteng/reservation_system/db"
	"github.com/adelberteng/reservation_system/utils"
)

var (
	logger = utils.Logger
	engine = db.Engine
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

func GetUserByName(name string) (User, error) {
	queryResult, err := engine.Table("user_tbl").Where("name = ? ", name).QueryString()
	if queryResult == nil {
		return User{}, errors.New("the user is not exist.")
	} else if err != nil {
		return User{}, err
	}

	user := User{
		Uid:          queryResult[0]["uid"],
		Name:         name,
		PasswordHash: queryResult[0]["password_hash"],
		Phone:        queryResult[0]["phone"],
		Email:        queryResult[0]["email"],
	}

	return user, nil
}
