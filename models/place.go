package models

type Place struct {
	Id        string `xorm:"not null pk autoincr INT(11)"`
	PlaceName string `xorm:"not null unique VARCHAR(32)"`
	OwnerId   string `xorm:"not null INT(11)"`
	RegionId  string `xorm:"not null INT(11)"`
	TypeId    string `xorm:"not null INT(11)"`
	Address   string `xorm:"not null VARCHAR(255)"`
	Capacity  string `xorm:"not null INT(11)"`
}

func init() {
	if err := engine.Sync2(new(Place)); err != nil {
		logger.Error(err)
	}
}

func (p *Place) TableName() string {
	return "place_tbl"
}
