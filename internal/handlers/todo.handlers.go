package handlers

import (
	"echo-crate/internal/services"
	views "echo-crate/internal/views/todo_views"
	"net/http"
)

type TodoHandler struct {
	storage *services.TodoStorage
}

func NewTodoHandler(storage *services.TodoStorage) *TodoHandler {
	return &TodoHandler{
		storage: storage,
	}
}

func (t *TodoHandler) RenderPage(w http.ResponseWriter, r *http.Request) {
	todos, err := t.storage.GetAllTodos()

	if err != nil {
		http.Error(w, "Failed to get all todos", http.StatusInternalServerError)
	}

	views.Index(todos).Render(r.Context(), w)
}

func (t *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
	}

	// Get form values
	title := r.FormValue("title")

	// Create the todo
	todo, err := t.storage.CreateTodo(title)

	if err != nil {
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
	}

	views.TodoView(todo).Render(r.Context(), w)
}
