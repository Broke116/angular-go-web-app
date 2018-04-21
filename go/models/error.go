package models

// Error is used to return error/response objects
type Error struct {
	Errortype  string `json:"errortype,omnitype"`
	Statuscode int `json:"statuscode,omnitype"`
}
