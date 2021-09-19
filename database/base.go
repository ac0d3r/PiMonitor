package database

import (
	"log"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init(dsn string) error {
	var err error
	db, err = gorm.Open(sqlite.Open(strings.TrimSpace(dsn)), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := migrate(); err != nil {
		log.Printf("[InitDB] Migrate sqlite tables error: %v", err)
	}

	return nil
}

func migrate() error {
	return db.AutoMigrate(
		&User{},
	)
}
