package user

import (
	"angular-go-web-app/go/models"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserModel is the database representation of User object
type UserModel struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	Username  string        `bson:"username"`
	Password  string        `bson:"password"`
	Email     string        `bson:"email"`
	StartDate time.Time     `bson:"start_date"`
	Status    string        `bson:"status"`
}

// UserModels is an array of UserModel
type UserModels []UserModel

// UserModelIndex is used as an index.
func UserModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

// NewUserModel is for converting the model data in which the data comes from the api, to mongodb model.
func NewUserModel(u *models.User) *UserModel {
	return &UserModel{
		Id:        u.Id,
		Username:  u.Username,
		Password:  u.Password,
		Email:     u.Email,
		StartDate: bson.Now(),
		Status:    u.Status}
}

// ToUser is a method which is used for getting data from the database and pushing it to the api used to show data to the client.
func (u *UserModel) ToUser() *models.User {
	return &models.User{
		Id:        u.Id,
		Username:  u.Username,
		Password:  u.Password,
		Email:     u.Email,
		StartDate: u.StartDate,
		Status:    u.Status}
}
