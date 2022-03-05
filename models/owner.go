package models

import (
	"errors"
)

type Owner struct {
	Id           string `xorm:"not null pk autoincr INT(11)"`
	CompanyName  string `xorm:"not null unique VARCHAR(32)"`
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

func (o *Owner) TableName() string {
	return "owner_tbl"
}

func GetOwnerByName(companyName string) (Owner, error) {
	queryResult, err := engine.Table("owner_tbl").Where("company_name = ? ", companyName).QueryString()
	if queryResult == nil {
		return Owner{}, errors.New("the owner is not exist.")
	} else if err != nil {
		return Owner{}, err
	}

	owner := Owner{
		Id:           queryResult[0]["id"],
		CompanyName:  companyName,
		PasswordHash: queryResult[0]["password_hash"],
		Phone:        queryResult[0]["phone"],
		Email:        queryResult[0]["email"],
	}

	return owner, nil
}
