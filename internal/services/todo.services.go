package services

import (
	"echo-crate/internal/models"
	"echo-crate/internal/repository"
	"gorm.io/gorm"
)

type TodoServiceInterface interface {
	GetAllTodos() ([]models.Todo, error)
	CreateTodo(title string) (models.Todo, error)
}

type TodoService struct {
	repo *repository.TodoRepository
}

func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{
		repo: repository.NewTodoRepository(db),
	}
}

func (s *TodoService) GetAllTodos() ([]models.Todo, error) {
	todos, err := s.repo.GetAllTodos()

	if err != nil {
		return todos, err
	}

	return todos, nil
}

func (s *TodoService) CreateTodo(title string) (models.Todo, error) {
	todo, err := s.repo.CreateTodo(title)

	if err != nil {
		return todo, err
	}

	return todo, nil
}
