package services

import (
	"fmt"

	"gorm.io/gorm"
)

type Todo struct {
	ID        int    `gorm:"primarkey;size:16"`
	Title     string `gorm:"title"`
	Completed bool   `gorm:"bool"`
}

type TodoStorage struct {
	db *gorm.DB
}

func NewTodoStorage(db *gorm.DB) *TodoStorage {
	return &TodoStorage{
		db: db,
	}
}

func (s *TodoStorage) GetAllTodos() ([]Todo, error) {
	todos := []Todo{}

	if err := s.db.Find(&todos).Error; err != nil {
		return todos, err
	}

	return todos, nil
}

func (s *TodoStorage) CreateTodo(title string) (Todo, error) {
	todo := &Todo{
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
