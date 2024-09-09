package database

import (
	"echo-crate/internal/models"

	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(models.Todo{})

	if err != nil {
		log.Fatal("failed to migrate database", err)
	}
}
