package member

import (
	"angular-go-web-app/go/config/mongo"
	"angular-go-web-app/go/models"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MemberService is used store collection information.
type MemberService struct {
	collection *mgo.Collection
}

// MemberServiceInstance is an instance of the MemberService
func MemberServiceInstance(session *mongo.Session, database string, collectionName string) *MemberService {
	collection := session.GetCollection(database, collectionName)
	collection.EnsureIndex(memberModelIndex())
	return &MemberService{collection}
}

// GetMembers is a method of MemberService
func (ms *MemberService) GetMembers() (*models.Members, error) {
	members := models.Members{}
	err := ms.collection.Find(bson.M{}).All(&members)
	if err != nil {
		fmt.Println("Error was occured when fetching the list of members ", err)
	}

	return &members, err
}

// GetMemberByID is a method of MemberService
func (ms *MemberService) GetMemberByID(id string) (*models.Member, error) {
	member := memberModel{}
	err := ms.collection.FindId(bson.ObjectIdHex(id)).One(&member)
	return member.toMember(), err
}

// InsertMember is a method of MemberService
func (ms *MemberService) InsertMember(m *models.Member) error {
	member := newMemberModel(m)
	fmt.Println("Member service insert member ", member)

	return ms.collection.Insert(&member)
}

// UpdateMember is a method of MemberService
func (ms *MemberService) UpdateMember(m *models.Member, _id string) error {
	member := newMemberModel(m)
	return ms.collection.UpdateId(bson.ObjectIdHex(_id), member)
}

// DeleteMember is a method of MemberService
func (ms *MemberService) DeleteMember(memberID string) error {
	return ms.collection.RemoveId(bson.ObjectIdHex(memberID))
}
