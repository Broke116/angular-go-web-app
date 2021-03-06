package services

import (
	"angular-go-web-app/go/config/mongo"
	"angular-go-web-app/go/config/mongo/member"
	"angular-go-web-app/go/models"
	"angular-go-web-app/go/utils/security"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MemberService is used store collection information.
type MemberService struct {
	collection *mgo.Collection
	hash       security.Hash
}

// MemberServiceConstructor is an instance of the MemberService
func MemberServiceConstructor(session *mongo.Session, database string, collectionName string, hash security.Hash) *MemberService {
	collection := session.GetCollection(database, collectionName)
	collection.EnsureIndex(member.MemberModelIndex())
	return &MemberService{collection, hash}
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
	member := member.MemberModel{}
	err := ms.collection.FindId(bson.ObjectIdHex(id)).One(&member)
	return member.ToMember(), err
}

// InsertMember is a method of MemberService
func (ms *MemberService) InsertMember(m *models.Member) error {
	member := member.NewMemberModel(m)
	fmt.Println("Member service insert member ", member)

	member.Password = ms.hash.Generate(m.Password)

	return ms.collection.Insert(&member)
}

// UpdateMember is a method of MemberService
func (ms *MemberService) UpdateMember(m *models.Member, _id string) error {
	member := member.NewMemberModel(m)
	member.Password = ms.hash.Generate(m.Password)
	return ms.collection.UpdateId(bson.ObjectIdHex(_id), member)
}

// DeleteMember is a method of MemberService
func (ms *MemberService) DeleteMember(memberID string) error {
	return ms.collection.RemoveId(bson.ObjectIdHex(memberID))
}
