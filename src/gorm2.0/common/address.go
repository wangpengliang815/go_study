package common

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserId               int
	Province, City, Area string
}

func (Address) TableName() string {
	return "Address"
}
