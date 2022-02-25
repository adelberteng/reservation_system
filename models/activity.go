package models

import ()

type Activity struct {
	Id        string `xorm:"not null pk autoincr INT(11)"`
	ClientId  string `xorm:"not null INT(11)"`
	PlaceId   string `xorm:"not null INT(11)"`
	StartTime string `xorm:"not null time.Time`
	EndTime   string `xorm:"not null time.Time`
}

func (a *Activity) TableName() string {
	return "activity_tbl"
}
