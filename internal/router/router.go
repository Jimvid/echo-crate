package router

import (
	"net/http"

	"echo-crate/internal/handlers"
	"echo-crate/internal/services"
	page "echo-crate/internal/views/pages"

	"github.com/jmoiron/sqlx"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		page.NotFound().Render(r.Context(), w)
		return
	}

	page.Index().Render(r.Context(), w)
}

func New(db *sqlx.DB) http.Handler {
	mux := http.NewServeMux()

	// Serve static files
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", fs))

	// Routes
	mux.HandleFunc("GET /", Index)

	// Todos
	todoStorage := services.NewTodoStorage(db)
	todoHandler := handlers.NewTodoHandler(todoStorage)

	mux.HandleFunc("GET /todos", todoHandler.RenderPage)
	mux.HandleFunc("POST /create-todo", todoHandler.Create)

	return mux
}
