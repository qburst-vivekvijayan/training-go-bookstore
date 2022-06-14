package models

type Genre struct {
	Id     int64  `json:"id" gorm:"primaryKey;auto_increment;not_null"`
	Genre  string `json:"genre"`
	BookId int64  `json:"-"`
}
