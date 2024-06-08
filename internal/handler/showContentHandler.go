package handler

import (
	"database/sql"
	"net/http"
)

func ShowContentHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("content")
	if id == "" {
		sendError(w, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	response := CreateOpeningRequest{}

	query := `SELECT name, type, genres,author, duration FROM content where id=$1`

	err := db.QueryRow(query, id).
		Scan(&response.Name, &response.Type, &response.Genres, &response.Author, &response.Duration)
	if err != nil {
		if err == sql.ErrNoRows {
			sendError(w, http.StatusNotFound, "no ID founded")
			return
		}
		logger.Errorf("show row by id error: %v", err)
		sendError(w, http.StatusInternalServerError, "show by id error")
		return
	}

	sendSuccess(w, "show-content", response) // Query Delete
}
