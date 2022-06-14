package handlers

import (
	"database/sql"

	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

type phandler struct {
	DB *sql.DB
}

func pNew(db *sql.DB) phandler {
	return phandler{db}
}
