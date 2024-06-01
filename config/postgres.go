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

func InitializePostgres() error {
	logger := *GetLogger("postgres")

	// Create DB connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Errorf("open postgres error: %v", err)
		return err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Errorf("ping error: %v", err)
		return err
	}

	logger.Info("Successfully connected!")
	return nil
}
