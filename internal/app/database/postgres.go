package database

import (
	"gorm.io/gorm"
)

type Database struct {
	Conn *gorm.DB
}

func NewDBConnection() *Database {
	return &Database{}
}
