package utils

import (
	"database/sql"
	"log"
)

func CleanExpiredSessions(db *sql.DB) func() {
	return func() {
		query := "DELETE FROM session WHERE expires_at < NOW()"
		_, err := db.Exec(query)
		if err != nil {
			log.Println(err)
		}
	}
}
