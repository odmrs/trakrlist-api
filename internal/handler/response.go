package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func sendError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if msg == "" {
		return
	}

	errorMessage := map[string]string{
		"message": msg,
	}

	messageJson, err := json.Marshal(errorMessage)
	if err != nil {
		fmt.Println("marshal error")
	}

	fmt.Fprint(w, string(messageJson))
}

func sendSuccess(w http.ResponseWriter, op string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	errorMessage := map[string]interface{}{
		"message": fmt.Sprintf("operation from handler: %s sucessfull", op),
		"data":    data,
	}

	messageJson, err := json.Marshal(errorMessage)
	if err != nil {
		fmt.Println("marshal error")
	}

	fmt.Fprint(w, string(messageJson))
}
