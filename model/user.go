package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" form:"name" json:"name"`
	Telephone string `gorm:"varchar(11; not null; unique)" form:"telephone" json:"telephone"`
	Password string `gorm:"size:255;not null" form:"password" json:"password"`
}
