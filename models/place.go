package models

import (
// "errors"
)

type Region struct {
	Id         string `xorm:"not null pk autoincr INT(11)"`
	RegionName string `xorm:"not null unique VARCHAR(32)"`
}

func (r *Region) TableName() string {
	return "region_tbl"
}

type PlaceType struct {
	Id       string `xorm:"not null pk autoincr INT(11)"`
	TypeName string `xorm:"not null unique VARCHAR(32)"`
}

func (p *PlaceType) TableName() string {
	return "place_type_tbl"
}

type Place struct {
	Id        string `xorm:"not null pk autoincr INT(11)"`
	PlaceName string `xorm:"not null unique VARCHAR(32)"`
	OwnerId   string `xorm:"not null INT(11)"`
	RegionId  string `xorm:"not null INT(11)"`
	TypeId    string `xorm:"not null INT(11)"`
	Address   string `xorm:"not null VARCHAR(255)"`
	Capacity  string `xorm:"not null INT(11)"`
}

func (p *Place) TableName() string {
	return "place_tbl"
}

func init() {
	err := engine.Sync2(new(Region))
	err = engine.Sync2(new(PlaceType))
	err = engine.Sync2(new(Place))
	if err != nil {
		logger.Error(err)
	}
}

func ViewRegion() ([]map[string]string, error) {
	results, err := engine.QueryString("select id, region_name from region_tbl")
	if err != nil {
		return nil, err
	}

	return results, nil
}

func ViewPlaceType() ([]map[string]string, error) {
	results, err := engine.QueryString("select id, type_name from place_type_tbl")
	if err != nil {
		return nil, err
	}

	return results, nil
}

func ViewPlace() ([]map[string]string, error) {
	results, err := engine.QueryString("select * from place_tbl")
	if err != nil {
		return nil, err
	}

	return results, nil
}
