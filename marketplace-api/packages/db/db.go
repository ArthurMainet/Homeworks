package db

import (
	"Email-API/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDB(config *config.Config) *DB {
	db, err := gorm.Open(postgres.Open(config.DB.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &DB{
		DB: db,
	}
}
