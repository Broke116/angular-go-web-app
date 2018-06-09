package apimember

import (
	"angular-go-web-app/go/models"
	"encoding/json"
	"fmt"
	"log"
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
	session, err := mgo.Dial(server) // TO DO add timeout
	if err != nil {
		log.Fatalln("Session could not be opened. Failed to connect to server.")
	}

	mongoSession.session = session
}

// GetMembersEndpoint returns all users.
func GetMembersEndpoint(w http.ResponseWriter, req *http.Request) {
	col := mongoSession.session.DB(database).C(collection)
	var members []models.Member
	col.Find(bson.M{}).All(&members)

	// fmt.Println(req.Method, req.URL)

	json.NewEncoder(w).Encode(members)
	w.Header().Set("Content-Type", "application/json;")
	return
}

// GetMemberEndpoint - GET / member - returns a specific user
func GetMemberEndpoint(w http.ResponseWriter, req *http.Request) {
	// get id parameter
	vars := mux.Vars(req)
	id := vars["id"]

	// check if the id is valid
	if !bson.IsObjectIdHex(id) {
		checkError(w, "Invalid ObjectId", http.StatusNotFound) // 404 status code
	}

	memberID := bson.ObjectIdHex(id)
	m := models.Member{}

	if err := mongoSession.session.DB(database).C(collection).Find(bson.D{{"id", memberID}}).One(&m); err != nil {
		checkError(w, "Error when getting the member", http.StatusNotFound) // 404 status code
	}

	// fmt.Println(req.Method, req.URL, "/", id)

	json.NewEncoder(w).Encode(m)
	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(http.StatusOK)
	return
}

// InsertMemberEndpoint - POST / insertMember - creates a new member
func InsertMemberEndpoint(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	var member models.Member
	if err := json.NewDecoder(req.Body).Decode(&member); err != nil { // decode body
		checkError(w, err.Error(), http.StatusInternalServerError) // 500 status code
	}

	member.ID = bson.NewObjectId()
	if err := mongoSession.session.DB(database).C(collection).Insert(&member); err != nil {
		checkError(w, err.Error(), http.StatusInternalServerError) // 500 status code
	}

	// fmt.Println(req.Method, req.URL)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 status code
	return
}

// UpdateMemberEndpoint - v1/updateMember PUT - is used to update the related member
func UpdateMemberEndpoint(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	// get id parameter
	vars := mux.Vars(req)
	id := vars["id"]

	// check if the id is valid
	if !bson.IsObjectIdHex(id) {
		checkError(w, "Invalid ObjectId", http.StatusNotFound) // 404 status code
	}

	var member models.Member
	if err := json.NewDecoder(req.Body).Decode(&member); err != nil {
		checkError(w, err.Error(), http.StatusInternalServerError) // 500 status code
	}

	member.ID = bson.NewObjectId()
	if err := mongoSession.session.DB(database).C(collection).UpdateId(member.ID, member); err != nil {
		checkError(w, err.Error(), http.StatusInternalServerError) // 500 status code
	}

	// fmt.Println(req.Method, req.URL)
	w.Header().Set("Content-Type", "application/json")
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
		checkError(w, "Invalid ObjectId", http.StatusNotFound) // 404 status code
	}

	member := bson.ObjectIdHex(id)

	if err := mongoSession.session.DB(database).C(collection).Remove(bson.D{{"id", member}}); err != nil {
		checkError(w, err.Error(), http.StatusInternalServerError) // 500 status code
	}

	// fmt.Println(req.Method, req.URL)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 status code
	return
}

func checkError(w http.ResponseWriter, err string, statusCode int) {
	fmt.Println(&models.Error{Definition: err, Statuscode: statusCode})
	return
}
