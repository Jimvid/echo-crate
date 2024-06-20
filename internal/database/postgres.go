package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection() *gorm.DB {
	dsn := fmt.Sprintf("%s sslmode=disable", os.Getenv("DB_DSN"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
