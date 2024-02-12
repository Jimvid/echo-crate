package router

import (
	"github.com/a-h/templ"
	"github.com/jimvid/echo-crate/internal/templates"
	"net/http"
)

func RenderHome(w http.ResponseWriter, r *http.Request) {
	component := templates.Base("home")
	err := component.Render(r.Context(), w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func New() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /", templ.Handler(templates.Base("home")))
	mux.HandleFunc("GET /handler", RenderHome)

	return mux
}
