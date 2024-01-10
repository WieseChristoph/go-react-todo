package models

import (
	"database/sql"
	"time"
)

type Session struct {
	ID        int
	Token     string
	UserId    int
	ExpiresAt time.Time
}

func (session Session) IsExpired() bool {
	return session.ExpiresAt.Before(time.Now())
}

func (session Session) Save(db *sql.DB) (Session, error) {
	query := `INSERT INTO session (token, user_id, expires_at) VALUES ($1, $2, $3) RETURNING id`
	row := db.QueryRow(query, session.Token, session.UserId, session.ExpiresAt)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return Session{}, err
	}

	createdSession := Session{
		ID:        id,
		Token:     session.Token,
		UserId:    session.UserId,
		ExpiresAt: session.ExpiresAt,
	}

	return createdSession, nil
}

func GetSessionByToken(db *sql.DB, token string) (Session, error) {
	query := `SELECT * FROM session WHERE token = $1`
	row := db.QueryRow(query, token)

	var session Session
	err := row.Scan(&session.ID, &session.Token, &session.UserId, &session.ExpiresAt)
	if err != nil {
		return Session{}, err
	}

	return session, nil
}

func (session Session) Delete(db *sql.DB) error {
	query := `DELETE FROM session WHERE id = $1`
	_, err := db.Exec(query, session.ID)
	if err != nil {
		return err
	}

	return nil
}
