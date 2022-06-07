package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Id    int64  `json:"id" gorm:"primaryKey"`
	Genre string `json:"genre"`
}
