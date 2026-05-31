package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {

	connStr := `
	host=localhost
	port=5432
	user=admin
	password=password
	dbname=hospital_analytics
	sslmode=disable
	`

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	fmt.Println("Database ping successful")

	return db, nil
}
