package services

import (
	"angular-go-web-app/go/config/mongo"
	"angular-go-web-app/go/config/mongo/user"
	"angular-go-web-app/go/models"
	"angular-go-web-app/go/utils/security"
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserService provides the required logic to handle user operations
type UserService struct {
	collection *mgo.Collection
	hash       security.Hash
}

// UserServiceConstructor is used to instantinate the service of User
func UserServiceConstructor(session *mongo.Session, database string, collectionName string, hash security.Hash) *UserService {
	collection := session.GetCollection(database, collectionName)
	collection.EnsureIndex(user.UserModelIndex())
	return &UserService{collection, hash}
}

// Authenticate is used to authenticate a user by accessing to mongo database
func (us *UserService) Authenticate(c models.Credentials) (models.User, error) {
	model := user.UserModel{}
	err := us.collection.Find(bson.M{"email": c.Email}).One(&model)

	if model.Password != "" {
		err = us.hash.Compare(model.Password, c.Password)
	}

	if err != nil {
		fmt.Println("Service Unauthorized")
		return models.User{}, err
	}

	fmt.Println("Service Authorized")
	return models.User{
		Id:        model.Id,
		Email:     model.Email,
		Password:  "-",
		Username:  model.Username,
		Status:    model.Status,
		StartDate: model.StartDate}, err
}

// CreateUser is used to insert a user to the database.
func (us *UserService) CreateUser(u *models.User) error {
	u.Password = us.hash.Generate(u.Password)
	user := user.NewUserModel(u)
	fmt.Println("User service insert user ", user)

	return us.collection.Insert(&user)
}
