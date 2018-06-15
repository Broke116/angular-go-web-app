package member

import (
	"angular-go-web-app/go/models"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MemberModel is the definition for mongodb object for member
type MemberModel struct {
	Id         bson.ObjectId `bson:"_id,omitempty"`
	Fullname   string        `bson:"fullname"`
	Password   string        `bson:"password"`
	Gender     string        `bson:"gender"`
	Age        int           `bson:"age"`
	Company    string        `bson:"company"`
	Email      string        `bson:"email"`
	Phone      string        `bson:"phone"`
	InsertDate time.Time     `bson:"insert_date"`
	Status     string        `bson:"status"`
}

// MemberModels is an array of MemberModels
type MemberModels []MemberModel

// MemberModelIndex is used as an index.
func MemberModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"_id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

// NewMemberModel is for converting the model data in which the data comes from the api, to mongodb model.
func NewMemberModel(m *models.Member) *MemberModel {
	return &MemberModel{
		Fullname:   m.Fullname,
		Gender:     m.Gender,
		Age:        m.Age,
		Company:    m.Company,
		Email:      m.Email,
		Phone:      m.Phone,
		InsertDate: bson.Now(),
		Status:     m.Status}
}

// ToMember is a method which is used for getting data from the database and pushing it to the api used to show data to the client.
func (m *MemberModel) ToMember() *models.Member {
	return &models.Member{
		Id:         m.Id,
		Fullname:   m.Fullname,
		Password:   m.Password,
		Gender:     m.Gender,
		Age:        m.Age,
		Company:    m.Company,
		Email:      m.Email,
		Phone:      m.Phone,
		InsertDate: m.InsertDate,
		Status:     m.Status}
}
