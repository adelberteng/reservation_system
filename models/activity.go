package models

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

func ViewActivity() ([]map[string]string, error) {
	results, err := engine.QueryString("select * from activity_tbl")
	if err != nil {
		return nil, err
	}

	return results, nil
}
