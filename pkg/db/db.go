package db

import (
	"database/sql"
	"fmt"
	"log"

	"mymodule/pkg/models"

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

	db.AutoMigrate(&models.Genre{})
	db.AutoMigrate(&models.Review{})
	db.AutoMigrate(&models.Book{})

	return db
}

func createConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open the connection
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}
