package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ReviewId int64  `json:"id" gorm:"primaryKey"`
	Review   string `json:"review"`
}
