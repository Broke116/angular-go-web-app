package models

// Address defines the fields of address struct
type Address struct {
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
}

// Member defines the fields of member struct
type Member struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
