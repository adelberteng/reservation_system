package models

type Region struct {
	Id         string `xorm:"not null pk autoincr INT(11)"`
	RegionName string `xorm:"not null unique VARCHAR(32)"`
}

func init() {
	if err := engine.Sync2(new(Region)); err != nil {
		logger.Error(err)
	}
}

func (r *Region) TableName() string {
	return "region_tbl"
}
