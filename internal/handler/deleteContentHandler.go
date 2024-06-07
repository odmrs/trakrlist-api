package handler

import (
	"database/sql"
	"net/http"
)

func DeleteContentHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("content")

	if id == "" {
		sendError(w, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	response := CreateOpeningRequest{}
	// Query Delete
	query := `DELETE from content WHERE id=$1 RETURNING name, type, genres, author, duration`
	err := db.QueryRow(query, id).
		Scan(&response.Name, &response.Type, &response.Genres, &response.Author, &response.Duration)
	if err != nil {
		if err == sql.ErrNoRows {
			sendError(w, http.StatusNotFound, "no ID founded")
			return
		}
		logger.Errorf("delete row by id error: %v", err)
		sendError(w, http.StatusInternalServerError, "delete by id error")
		return
	}

	sendSuccess(w, "delete-content", response)
}
