package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	BookId   int64    `json:"id" gorm:"primaryKey"`
	Title    string   `json:"title"`
	Quantity int      `json:"quantity"`
	GenreId  int      `json:"-" gorm:"unique"`
	Genre    Genre    `json:"genre" gorm:"foreignKey:GenreId;default:null"`
	Author   string   `json:"author"`
	ReviewId int      `json:"-" gorm:"unique"`
	Review   []Review `json:"review" gorm:"foreignKey:ReviewId;default:null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
