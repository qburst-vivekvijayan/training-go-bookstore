package models

type Review struct {
	ReviewId int64  `json:"id" gorm:"primaryKey;auto_increment;not_null"`
	Review   string `json:"review"`
	BookId   int64  `json:"-"`
}
