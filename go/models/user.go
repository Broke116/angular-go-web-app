package models

import "gopkg.in/mgo.v2/bson"

// User is a struct holding the definition for user's that is allowed to use this API
type User struct {
	ID        bson.ObjectId `json:"id" bson:"id"`
	Username  string        `json:"username" bson:"username"`
	Password  string        `json:"password" bson:"password"`
	Email     string        `json:"email" bson:"email"`
	StartDate string        `json:"start_date" bson:"start_date"`
	Status    bool          `json:"status" bson:"status"`
}
