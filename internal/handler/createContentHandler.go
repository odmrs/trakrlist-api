package handler

import (
	"encoding/json"
	"net/http"
)

func CreateContentHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	request := CreateOpeningRequest{}
	err := decoder.Decode(&request)
	if err != nil {
		logger.Errorf("decode request error: %v", err)
	}

	logger.Infof("request received: %+v", request)

	if err = request.Validate(); err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		logger.Errorf("validate request error: %v", err.Error())
		return
	}

	// Query Insert
	query := `INSERT INTO content (name, type, genres, author, duration)
    VALUES ($1, $2, $3, $4, $5)
    `

	_, err = db.Exec(
		query,
		request.Name,
		request.Type,
		request.Genres,
		request.Author,
		request.Duration,
	)
	if err != nil {
		logger.Errorf("insert query error: %v", err)
		sendError(w, http.StatusInternalServerError, "error insert data on database")
		return
	}

	logger.Info("insert request JSON with success")

	sendSuccess(w, "create-content", request)
}
