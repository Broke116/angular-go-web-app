package models

import "gopkg.in/mgo.v2/bson"

// Address defines the fields of address struct
type Address struct {
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
}

// Member defines the fields of member struct
type Member struct {
	ID         bson.ObjectId `json:"id" bson:"id`
	Fullname   string        `json:"fullname" bson:"fullname"`
	Gender     string        `json:"gender" bson:"gender"`
	Age        int           `json:"age" bson:"age"`
	Company    string        `json:"company" bson:"company"`
	Email      string        `json:"email" bson:"email"`
	Phone      string        `json:"phone" bson:"phone"`
	InsertDate string        `json:"insert_date" bson:"insert_date"`
	Status     bool          `json:"status" bson:"status"`
	//Address    *Address `json:"address,omitempty"`
}

// Members is an array of Member
type Members []Member
