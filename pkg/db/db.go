package db

import (
	"log"

	"github.com/qburst-vivekvijayan/training-go-bookstore.git/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "qburst123"
	dbname   = "book_store"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:qburst123@localhost:5432/book_store"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}
