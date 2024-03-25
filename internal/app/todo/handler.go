package todo

import (
	views "echo-crate/internal/app/todo/views"
	"fmt"
	"net/http"
)

type TodoHandler struct {
	storage *TodoStorage
}

func NewTodoHandler(storage *TodoStorage) *TodoHandler {
	return &TodoHandler{
		storage: storage,
	}
}

func (t *TodoHandler) renderPage(w http.ResponseWriter, r *http.Request) {
	views.Index().Render(r.Context(), w)
}

func (t *TodoHandler) create(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
	}

	// Get form values
	title := r.FormValue("title")

	// Create the todo
	_, err = t.storage.createTodo(title)

	if err != nil {
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
	}

	views.Todo(title).Render(r.Context(), w)
}
