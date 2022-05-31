package models

type Review struct {
	Id     int64  `json:"id" gorm:"primaryKey"`
	Review string `json:"review"`
}
