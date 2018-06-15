package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Address defines the fields of address struct
type Address struct {
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
}

// Member defines the fields of member struct
type Member struct {
	Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Fullname   string        `json:"fullname"`
	Password   string        `json:"password"`
	Gender     string        `json:"gender"`
	Age        int           `json:"age"`
	Company    string        `json:"company"`
	Email      string        `json:"email"`
	Phone      string        `json:"phone"`
	InsertDate time.Time     `json:"insert_date" bson:"insert_date"`
	Status     string        `json:"status"`
	//Address    *Address `json:"address,omitempty"`
}

// Members is an array of Member
type Members []Member

//MemberService is an interface providing the logic to accessing data.
type MemberService interface {
	GetMembers(members Members) (*Members, error)
	GetMemberByID(id string) (*Member, error)
	InsertMember(member *Member) error
	UpdateMember(member *Member, _id string) error
	DeleteMember(id string) error
}
