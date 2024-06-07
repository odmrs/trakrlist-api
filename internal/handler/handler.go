package handler

import (
	"database/sql"

	"github.com/odmrs/trakrlist-api/config"
)

var (
	logger *config.Logger
	db     *sql.DB
)

type Content struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Genres   string `json:"genres"`
	Author   string `json:"author"`
	Duration int    `json:"duration"`
}

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db, _ = config.InitializePostgres()
}
