package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User is a struct holding the definition for user's that is allowed to use this API
type User struct {
	Id        bson.ObjectId `json:"id" bson:"_id" form:"-"`
	Username  string        `json:"username" bson:"username" form:"username"`
	Password  string        `json:"password" bson:"password" form:"password"`
	Email     string        `json:"email" bson:"email"`
	StartDate time.Time     `json:"start_date" bson:"start_date"`
	Status    string        `json:"status" bson:"status"`
}

// Users is used store users inside an array.
type Users []User

// Credentials is used for logging process when a user sends their required details to login.
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserService is an interface providing the logic to accessing data.
type UserService interface {
	Authenticate(c Credentials) (User, error)
	CreateUser(user *User) error
	/*GetUserByID(id string) (*User, error)
	InsertUser(user *User) error
	UpdateUser(user *User, _id string) error
	DeleteUser(id string) error*/
}
