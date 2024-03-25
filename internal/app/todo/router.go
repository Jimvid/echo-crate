package todo

import (
	"net/http"
)

func AddTodoRoutes(mux *http.ServeMux, handler *TodoHandler) {
	mux.HandleFunc("GET /todos", handler.renderPage)
	mux.HandleFunc("POST /create-todo", handler.create)
}
