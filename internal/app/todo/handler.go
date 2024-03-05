package todo

import (
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

func (t *TodoHandler) create(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
	}

	// Get form values
	title := r.FormValue("title")
	description := r.FormValue("description")

	// Create the todo
	message, err := t.storage.createTodo(title, description)

	if err != nil {
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
	}

	fmt.Print(w, message)
}
