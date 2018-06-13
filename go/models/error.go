package models

import (
	"fmt"
	"net/http"
)

// Error is used to return error/response objects
type Error struct {
	Definition string `json:"definition,omnitype"`
	Statuscode int    `json:"statuscode,omnitype"`
}

// CheckError is used to return an error
func CheckError(w http.ResponseWriter, err string, statusCode int) {
	fmt.Println(&Error{Definition: err, Statuscode: statusCode})
	return
}
