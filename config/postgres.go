package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "root"
	password = "root"
	dbname   = "postgre_db"
)

func InitializePostgres() (*sql.DB, error) {
	logger := *GetLogger("postgres")

	// Create DB connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Errorf("open postgres error: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Errorf("ping error: %v", err)
		return nil, err
	}

	logger.Info("Successfully connected!")
	createTableSQL := `
  CREATE TABLE IF NOT EXISTS content (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  type TEXT NOT NULL,
  genres TEXT NOT NULL,
  author TEXT NOT NULL,
  duration INT NOT NULL
  );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		logger.Errorf("failed to create table: %v", err)
		return nil, err
	}

	logger.Info("table 'content' created or already exists")
	return db, nil
}
