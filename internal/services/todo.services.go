package services

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type Todo struct {
	ID        int    `db:"id"`
	Title     string `db:"title"`
	Completed bool   `db:"completed"`
}

type TodoStorage struct {
	db *sqlx.DB
}

func NewTodoStorage(db *sqlx.DB) *TodoStorage {
	return &TodoStorage{
		db: db,
	}
}

func (s *TodoStorage) GetAllTodos() ([]Todo, error) {
	query := "SELECT * FROM todo ORDER BY id"
	todos := []Todo{}

	err := s.db.Select(&todos, query)

	if err != nil {
		return todos, err
	}

	return todos, nil
}

func (s *TodoStorage) CreateTodo(title string) (Todo, error) {
	todo := &Todo{
		Title:     title,
		Completed: false,
	}

	createdTodo := Todo{}

	query := "INSERT INTO todo (title, completed) VALUES ($1, $2) RETURNING id, title, completed"
	if err := s.db.QueryRowx(query, todo.Title, todo.Completed).StructScan(&createdTodo); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Inserted user: %+v\n", createdTodo)

	return createdTodo, nil
}
