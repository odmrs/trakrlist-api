package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
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
