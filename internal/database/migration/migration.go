package database

import (
	"echo-crate/internal/services"

	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(services.Todo{})

	if err != nil {
		log.Fatal("failed to migrate database", err)
	}
}
