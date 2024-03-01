package database

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

var (
	dburl = os.Getenv("DB_URL")
)

func NewSQLiteConnection() *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", dburl)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
