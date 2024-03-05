package todo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type todoDB struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}

type TodoStorage struct {
	db *sqlx.DB
}

func NewTodoStorage(db *sqlx.DB) *TodoStorage {
	return &TodoStorage{
		db: db,
	}
}

func (s *TodoStorage) createTodo(title string, description string) (string, error) {
	fmt.Print(title)
	fmt.Print(description)
	todo := &todoDB{
		Title:       title,
		Description: description,
	}

	query := "INSERT INTO todo (title, description) VALUES (:title, :description)"
	result, err := s.db.NamedExec(query, todo)
	if err != nil {
		return "", err
	}

	lastID, err := result.LastInsertId()

	if err != nil {
		return "", err
	}

	message := fmt.Sprintf("Insert operation successful. Last inserted ID: %d", lastID)
	return message, nil
}
