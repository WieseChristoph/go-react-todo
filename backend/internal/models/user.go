package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID         int       `json:"id,string"`
	Username   string    `json:"username"`
	GlobalName string    `json:"global_name"`
	Email      string    `json:"email"`
	Avatar     string    `json:"avatar"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
}

func GetAllUsers(db *sql.DB) ([]User, error) {
	query := `SELECT * FROM "user"`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.GlobalName, &user.Email, &user.Avatar, &user.UpdatedAt, &user.CreatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUserById(db *sql.DB, id int) (User, error) {
	query := `SELECT * FROM "user" WHERE id = $1`
	row := db.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.GlobalName, &user.Email, &user.Avatar, &user.UpdatedAt, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (user User) Save(db *sql.DB) (User, error) {
	query := `INSERT INTO "user" (id, username, global_name, email, avatar) VALUES ($1, $2, $3, $4, $5) ON CONFLICT (id) DO UPDATE SET username = $2, global_name = $3, email = $4, avatar = $5 RETURNING id`
	_, err := db.Exec(query, user.ID, user.Username, user.GlobalName, user.Email, user.Avatar)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (user User) Update(db *sql.DB) (User, error) {
	query := `UPDATE "user" SET username = $1, global_name = $2, email = $3, avatar = $4 WHERE id = $5`
	_, err := db.Exec(query, user.Username, user.GlobalName, user.Email, user.Avatar, user.ID)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (user User) Delete(db *sql.DB) error {
	query := `DELETE FROM "user" WHERE id = $1`
	_, err := db.Exec(query, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUserById(db *sql.DB, id int) error {
	query := `DELETE FROM "user" WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
