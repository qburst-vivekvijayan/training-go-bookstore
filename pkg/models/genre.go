package models

type Genre struct {
	Id    int64  `json:"id" gorm:"primaryKey"`
	Genre string `json:"genre"`
}
