package apimember

import (
	"angular-go-web-app/go/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	server     = "mongodb://db:27017"
	database   = "airline"
	collection = "members"
)

// MongoSession is a struct holding the session of Mongodb
type MongoSession struct {
	session *mgo.Session
}

var mongoSession = MongoSession{}

func init() {
	session, err := mgo.Dial(server)
	if err != nil {
		panic(err)
	}

	mongoSession.session = session
}

// GetMembersEndpoint returns all users.
func GetMembersEndpoint(w http.ResponseWriter, req *http.Request) {
	col := mongoSession.session.DB(database).C(collection)
	var members []models.Member
	col.Find(bson.M{}).All(&members)

	fmt.Println(req.Method, req.URL)

	json.NewEncoder(w).Encode(members)
	w.Header().Set("Content-Type", "application/json;")
	return
}

// GetMemberEndpoint - GET / member - returns a specific user
func GetMemberEndpoint(w http.ResponseWriter, req *http.Request) {
	// get id paramtere
	vars := mux.Vars(req)
	id := vars["id"]

	// check if the id is valid
	if !bson.IsObjectIdHex(id) {
		CheckError(w, "Invalid ObjectId", http.StatusNotFound) // 404 status code
	}

	member := bson.ObjectIdHex(id)

	if err := mongoSession.session.DB(database).C(collection).FindId(member); err != nil {
		CheckError(w, "Error when getting the member", http.StatusInternalServerError) // 500 status code
	}

	fmt.Println(req.Method, req.URL, "/", id)
	data, _ := json.Marshal(member)

	json.NewEncoder(w).Encode(member)
	w.Header().Set("Content-Type", "application/json;")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// InsertMemberEndpoint - POST / insertMember - creates a new member
func InsertMemberEndpoint(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	var member models.Member
	if err := json.NewDecoder(req.Body).Decode(&member); err != nil { // decode body
		CheckError(w, err.Error(), http.StatusInternalServerError) // 500 status code
	}

	member.ID = bson.NewObjectId()
	if err := mongoSession.session.DB(database).C(collection).Insert(&member); err != nil {
		CheckError(w, err.Error(), http.StatusInternalServerError) // 500 status code
	}

	fmt.Println(req.Method, req.URL)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 status code
	return
}

// UpdateMemberEndpoint - v1/updateMember PUT - is used to update the related member
func UpdateMemberEndpoint(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	var member models.Member
	if err := json.NewDecoder(req.Body).Decode(&member); err != nil {
		CheckError(w, err.Error(), http.StatusInternalServerError) // 500 status code
	}

	member.ID = bson.NewObjectId()
	if err := mongoSession.session.DB(database).C(collection).UpdateId(member.ID, member); err != nil {
		CheckError(w, err.Error(), http.StatusInternalServerError) // 500 status code
	}

	fmt.Println(req.Method, req.URL)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK) // status code
	return
}

// DeleteMemberEndpoint - DELETE member/{id} / member - deletes the member
func DeleteMemberEndpoint(w http.ResponseWriter, req *http.Request) {
	// get id paramtere
	vars := mux.Vars(req)
	id := vars["id"]

	// check if the id is valid
	if !bson.IsObjectIdHex(id) {
		CheckError(w, "Invalid ObjectId", http.StatusNotFound) // 404 status code
	}

	member := bson.ObjectIdHex(id)

	if err := mongoSession.session.DB(database).C(collection).Remove(member); err != nil {
		CheckError(w, err.Error(), http.StatusInternalServerError) // 500 status code
	}

	fmt.Println(req.Method, req.URL)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK) // 200 status code
	return
}

// CheckError is used to handle endpoint errors
func CheckError(w http.ResponseWriter, err string, statusCode int) {
	fmt.Println(&models.Error{Definition: err, Statuscode: statusCode})
	w.WriteHeader(statusCode)
	return
}
