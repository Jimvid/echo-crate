package todo

import (
	"net/http"
)

func AddTodoRoutes(mux *http.ServeMux, handler *TodoHandler) {
	mux.HandleFunc("POST /todos", handler.create)
}
