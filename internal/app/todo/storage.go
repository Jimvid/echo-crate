package todo

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type todoDB struct {
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

func (s *TodoStorage) createTodo(title string) (string, error) {
	todo := &todoDB{
		Title:     title,
		Completed: false,
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS todo(
	   id SERIAL PRIMARY KEY,
	   title TEXT NOT NULL,
	   completed BOOLEAN
	   );`

	_, err := s.db.Exec(createTableSQL)

	if err != nil {
		log.Fatal(err)
	}

	query := "INSERT INTO todo (title, completed) VALUES (:title, :completed)"
	result, err := s.db.NamedExec(query, todo)

	if err != nil {
		return "", err
	}

	lastID, err := result.RowsAffected()

	if err != nil {
		return "", err
	}

	message := fmt.Sprintf("Insert operation successful. Last inserted ID: %d", lastID)
	return message, nil
}
