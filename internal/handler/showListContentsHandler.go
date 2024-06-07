package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func ListAllContentHandler(w http.ResponseWriter, r *http.Request) {
	content, err := fetchContent(db)
	if err != nil {
		logger.Errorf("fetch content to json error: %v", err)
	}

	jsonData, err := json.Marshal(content)
	if err != nil {
		logger.Errorf("marshal to json content error: %v", err)
	}

	fmt.Println(string(jsonData))
	sendSuccess(w, "list-content", content)
}

func fetchContent(db *sql.DB) ([]Content, error) {
	var contents []Content

	rows, err := db.Query("SELECT id, name, type, genres, author, duration FROM content")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var content Content
		if err := rows.Scan(&content.ID, &content.Name, &content.Type, &content.Genres, &content.Author, &content.Duration); err != nil {
			return nil, err
		}

		contents = append(contents, content)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return contents, nil
}
