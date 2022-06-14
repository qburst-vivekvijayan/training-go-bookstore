package models

type Book struct {
	Id       int64    `json:"id" gorm:"primaryKey;auto_increment;not_null"`
	Title    string   `json:"title"`
	Quantity int      `json:"quantity"`
	GenreId  int      `json:"-" gorm:"unique"`
	Genre    Genre    `json:"genre,omitempty" gorm:"foreignKey:GenreId;default:null"`
	Author   string   `json:"author"`
	ReviewId int      `json:"-" gorm:"unique"`
	Review   []Review `json:"review,omitempty" gorm:"foreignKey:ReviewId;default:null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
