package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewPostgresDB(dataSource string) *DB {

	db, err := gorm.Open(postgres.Open(dataSource), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &DB{db}
}
