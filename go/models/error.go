package models

// Error is used to return error/response objects
type Error struct {
	Definition string `json:"definition,omnitype"`
	Statuscode int    `json:"statuscode,omnitype"`
}
