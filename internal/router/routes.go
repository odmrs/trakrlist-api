package router

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func initializeRoutes(v1 *http.ServeMux) {
	v1.HandleFunc("GET /api/v1/ping", ping)

	v1.HandleFunc("GET /api/v1/content", ping)
	v1.HandleFunc("POST /api/v1/content", ping)
	v1.HandleFunc("PUT /api/v1/content", ping)
	v1.HandleFunc("DELETE /api/v1/content", ping)
	v1.HandleFunc("GET /api/v1/contents", ping)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pong := map[string]string{
		"message": "pong",
	}

	pongJson, err := json.Marshal(pong)
	if err != nil {
		fmt.Println("marshal error")
	}

	fmt.Fprint(w, string(pongJson))
}
