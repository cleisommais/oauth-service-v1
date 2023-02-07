// db/postgres.go
package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// CreateDBConnection creates a new database connection using the PostgreSQL database driver.
func CreatePostgresConnection() (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
						os.Getenv("DB_USER"),
						os.Getenv("DB_PASSWORD"),
						os.Getenv("DB_HOST"),
						os.Getenv("DB_PORT"),
						os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
