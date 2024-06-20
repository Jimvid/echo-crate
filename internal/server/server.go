package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"echo-crate/internal/database"
	"echo-crate/internal/router"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

type Server struct {
	port int
	db   *gorm.DB
}

func New() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		db:   database.NewPostgresConnection(),
	}

	// Migration
	database.Migrate(NewServer.db)

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      router.New(NewServer.db),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
