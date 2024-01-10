package models

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID          int       `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserId      int       `json:"user_id"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

func GetAllTodosByUser(db *sql.DB, userId int) ([]Todo, error) {
	query := "SELECT * FROM todo WHERE user_id = $1 ORDER BY created_at DESC"
	rows, err := db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []Todo{}

	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.UserId, &todo.UpdatedAt, &todo.CreatedAt)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func GetTodoById(db *sql.DB, id int) (Todo, error) {
	query := "SELECT * FROM todo WHERE id = $1"
	row := db.QueryRow(query, id)

	var todo Todo
	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.UserId, &todo.UpdatedAt, &todo.CreatedAt)
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}

func (todo Todo) Save(db *sql.DB) (Todo, error) {
	query := "INSERT INTO todo (title, description, status, user_id) VALUES ($1, $2, $3, $4) RETURNING id"
	row := db.QueryRow(query, todo.Title, todo.Description, todo.Status, todo.UserId)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return Todo{}, err
	}

	createdTodo := Todo{
		ID:          id,
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
	}

	return createdTodo, nil
}

func (todo Todo) Update(db *sql.DB) (Todo, error) {
	query := "UPDATE todo SET title = $1, description = $2, status = $3 WHERE id = $4"
	_, err := db.Exec(query, todo.Title, todo.Description, todo.Status, todo.ID)
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}

func (todo Todo) Delete(db *sql.DB) error {
	query := "DELETE FROM todo WHERE id = $1"
	_, err := db.Exec(query, todo.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTodoById(db *sql.DB, id int) error {
	query := "DELETE FROM todo WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
