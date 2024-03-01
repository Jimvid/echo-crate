package router

import (
	"fmt"
	"net/http"

	layout "echo-crate/internal/views/layouts"

	"github.com/a-h/templ"
	"github.com/jmoiron/sqlx"
)

func RenderHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, http.StatusNotFound)
		return
	}

	component := layout.Base("home")
	err := component.Render(r.Context(), w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func errorHandler(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404 Not Found")
	}
}

func New(db *sqlx.DB) http.Handler {
	mux := http.NewServeMux()

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	// Application routes
	mux.Handle("GET /templ", templ.Handler(layout.Base("home")))
	mux.HandleFunc("GET /", RenderHome)

	return mux
}
