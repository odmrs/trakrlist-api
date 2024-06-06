package handler

import (
	"database/sql"

	"github.com/odmrs/trakrlist-api/config"
)

var (
	logger *config.Logger
	db     *sql.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db, _ = config.InitializePostgres()
}
