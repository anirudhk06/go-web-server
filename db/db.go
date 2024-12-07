package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresStorage(conn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	return db, err
}
