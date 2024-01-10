package utils

import (
	"database/sql"
	"log"
	"os"
	"slices"
)

func ConnectDB(driver string, connString string) (*sql.DB, error) {
	db, err := sql.Open(driver, connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CloseDB(db *sql.DB) error {
	return db.Close()
}

func MigrateDB(db *sql.DB) error {
	log.Println("Migrating database...")

	// Create migrations table if it doesn't exist
	createMigrationsTableQuery := `
		CREATE TABLE IF NOT EXISTS migration (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		);
	`
	_, err := db.Exec(createMigrationsTableQuery)
	if err != nil {
		return err
	}

	// Get all migrations from database (migration that have already been run)
	rows, err := db.Query(`
		SELECT name FROM migration;
	`)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Put migration names in slice
	var migrations []string
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return err
		}

		migrations = append(migrations, name)
	}

	path, err := os.Getwd()
	if err != nil {
		return err
	}

	// Get all migration files in migrations folder
	files, err := os.ReadDir(path + "/migrations")
	if err != nil {
		return err
	}

	for _, file := range files {
		// Skip file if it has already been run
		if slices.Contains(migrations, file.Name()) {
			continue
		}

		// Read migration file
		migration, err := os.ReadFile("./migrations/" + file.Name())
		if err != nil {
			return err
		}

		log.Println("Running migration: " + file.Name())

		// Run migration
		_, err = db.Exec(string(migration))
		if err != nil {
			return err
		}

		// Insert migration into database
		insertMigrationQuery := "INSERT INTO migration (name) VALUES ($1)"
		_, err = db.Exec(insertMigrationQuery, file.Name())
		if err != nil {
			return err
		}
	}

	log.Println("Migrations finished")

	return nil
}
