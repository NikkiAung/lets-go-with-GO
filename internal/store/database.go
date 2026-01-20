package store

import (
	"database/sql"
	"fmt"
)

func Open() (*sql.DB, error) {
	connectionString := "host=localhost user=user password=pass dbname=postgres port=5432 sslmode=disable"

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, fmt.Errorf("db: open error %w", err)
	}

	fmt.Println("Database connected!")

	return db, nil
}
