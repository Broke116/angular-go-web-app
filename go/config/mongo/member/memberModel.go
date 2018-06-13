package member

import (
	"angular-go-web-app/go/models"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type memberModel struct {
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

type memberModels []memberModel

func memberModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"_id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

// newMemberModel is for converting the model data in which the data comes from the api, to mongodb model.
func newMemberModel(m *models.Member) *memberModel {
	return &memberModel{
		Fullname:   m.Fullname,
		Password:   m.Password,
		Gender:     m.Gender,
		Age:        m.Age,
		Company:    m.Company,
		Email:      m.Email,
		Phone:      m.Phone,
		InsertDate: bson.Now(),
		Status:     m.Status}
}

// toMember is a method which is used for getting data from the database and pushing it to the api used to show data to the client.
func (m *memberModel) toMember() *models.Member {
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

// toMemberArray is a method which is used for getting data from the database and pushing it to the api used to show data to the client.
//func (m *memberModels) toMemberArray() *models.Members {
/*members := []*models.Members{
	&memberModels {
		memberModel
	}
}*/
//	var members models.Members { &m }
//	return &members
//}
