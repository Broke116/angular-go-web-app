package server

import (
	"encoding/json"
	"net/http"
)

// Error function is used return an error message
func Error(w http.ResponseWriter, statusCode int, message string) {
	JSON(w, statusCode, map[string]string{"error": message})
}

// JSON is used for returning an appropirate json response
func JSON(w http.ResponseWriter, statusCode int, content interface{}) {
	response, _ := json.Marshal(content)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
