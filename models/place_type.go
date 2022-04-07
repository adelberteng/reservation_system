package models

type PlaceType struct {
	Id       string `xorm:"not null pk autoincr INT(11)"`
	TypeName string `xorm:"not null unique VARCHAR(32)"`
}

func init() {
	if err := engine.Sync2(new(PlaceType)); err != nil {
		logger.Error(err)
	}
}

func (p *PlaceType) TableName() string {
	return "place_type_tbl"
}
