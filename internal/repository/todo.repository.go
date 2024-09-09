package repository

import (
	"echo-crate/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (s *TodoRepository) GetAllTodos() ([]models.Todo, error) {
	todos := []models.Todo{}

	if err := s.db.Find(&todos).Error; err != nil {
		return todos, err
	}

	return todos, nil
}

func (s *TodoRepository) CreateTodo(title string) (models.Todo, error) {
	todo := &models.Todo{
		Title:     title,
		Completed: false,
	}

	createdTodo := *todo

	if err := s.db.Create(todo).Error; err != nil {
		return createdTodo, err
	}

	fmt.Printf("Inserted user: %+v\n", createdTodo)

	return createdTodo, nil
}
