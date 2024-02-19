package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"echo-crate/internal/db"
	"echo-crate/internal/router"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
	db   database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		db:   database.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      router.New(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
