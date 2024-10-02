package handlers

import (
	"echo-crate/internal/services"
	views "echo-crate/internal/views/todo_views"
	"net/http"
)

type TodoHandler struct {
	todoService services.TodoServiceInterface
}

func NewTodoHandler(todoService *services.TodoService) *TodoHandler {
	return &TodoHandler{
		todoService: todoService,
	}
}

func (h *TodoHandler) RenderPage(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoService.GetAllTodos()

	if err != nil {
		http.Error(w, "Failed to get all todos", http.StatusInternalServerError)
	}

	views.Index(todos).Render(r.Context(), w)
}

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
	}

	title := r.FormValue("title")
	todo, err := h.todoService.CreateTodo(title)

	if err != nil {
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
	}

	views.TodoView(todo).Render(r.Context(), w)
}
