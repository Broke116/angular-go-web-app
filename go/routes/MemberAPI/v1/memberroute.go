package apimember

import (
	"angular-go-web-app/go/config/mongo/member"
	mongo "angular-go-web-app/go/config/mongo/member"
	"angular-go-web-app/go/models"
	"angular-go-web-app/go/utils/middlewares"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2/bson"
)

const (
	server     = "mongodb://db:27017"
	database   = "airline"
	collection = "members"
)

// MongoSession is a struct holding the session of Mongodb
/*type MongoSession struct {
	session *mgo.Session
}

var mongoSession = MongoSession{}

func init() {
	session, err := mgo.Dial(server) // TO DO add timeout
	if err != nil {
		log.Fatalln("Session could not be opened. Failed to connect to server.")
	}

	mongoSession.session = session
}*/

type memberRouter struct {
	memberService *member.MemberService
}

// NewMemberRouter is used to define routes.
func NewMemberRouter(ms *mongo.MemberService, r *mux.Router) *mux.Router {
	memberRouter := memberRouter{ms}

	mw := middlewares.ChainMiddleware(middlewares.Logging, middlewares.Tracing)

	r.HandleFunc("/", mw(memberRouter.GetMembersEndpoint)).Methods("GET")
	r.HandleFunc("/insertMember", mw(memberRouter.InsertMemberEndpoint)).Methods("POST")
	r.HandleFunc("/updateMember/{id}", mw(memberRouter.UpdateMemberEndpoint)).Methods("PUT")
	r.HandleFunc("/{id}", mw(memberRouter.GetMemberEndpoint)).Methods("GET")
	r.HandleFunc("/{id}", mw(memberRouter.DeleteMemberEndpoint)).Methods("DELETE")

	return r
}

// GetMembersEndpoint returns all users.
func (mr *memberRouter) GetMembersEndpoint(w http.ResponseWriter, req *http.Request) {
	members, err := mr.memberService.GetMembers()

	if err != nil {
		models.CheckError(w, "Error when getting the members", http.StatusNotFound) // 404 status code
	}

	response, _ := json.Marshal(members)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return
}

// GetMemberEndpoint - GET / member - returns a specific user
func (mr *memberRouter) GetMemberEndpoint(w http.ResponseWriter, req *http.Request) {
	// get id parameter
	vars := mux.Vars(req)
	id := vars["id"]

	// check if the id is valid
	if !bson.IsObjectIdHex(id) {
		models.CheckError(w, "Invalid ObjectId", http.StatusNotFound) // 404 status code
	}

	member, err := mr.memberService.GetMemberByID(id)
	if err != nil {
		models.CheckError(w, "Error when getting the member", http.StatusNotFound) // 404 status code
	}

	response, _ := json.Marshal(member)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return
}

// InsertMemberEndpoint - POST / insertMember - creates a new member
func (mr *memberRouter) InsertMemberEndpoint(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	var member models.Member
	if err := json.NewDecoder(req.Body).Decode(&member); err != nil { // decode body
		models.CheckError(w, err.Error(), http.StatusInternalServerError) // 500 status code
	}

	if err := mr.memberService.InsertMember(&member); err != nil {
		models.CheckError(w, err.Error(), http.StatusInternalServerError) // 500 status code
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return
}

// UpdateMemberEndpoint - v1/updateMember PUT - is used to update the related member
func (mr *memberRouter) UpdateMemberEndpoint(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	// get id parameter
	vars := mux.Vars(req)
	id := vars["id"]

	// check if the id is valid
	if !bson.IsObjectIdHex(id) {
		models.CheckError(w, "Invalid ObjectId", http.StatusNotFound) // 404 status code
	}

	var member models.Member
	if err := json.NewDecoder(req.Body).Decode(&member); err != nil {
		models.CheckError(w, err.Error(), http.StatusInternalServerError) // 500 status code
	}

	if err := mr.memberService.UpdateMember(&member, id); err != nil {
		models.CheckError(w, err.Error(), http.StatusNotFound) // 500 status code
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // status code
	return
}

// DeleteMemberEndpoint - DELETE member/{id} / member - deletes the member
func (mr *memberRouter) DeleteMemberEndpoint(w http.ResponseWriter, req *http.Request) {
	// get id paramtere
	vars := mux.Vars(req)
	id := vars["id"]

	// check if the id is valid
	if !bson.IsObjectIdHex(id) {
		models.CheckError(w, "Invalid ObjectId", http.StatusNotFound) // 404 status code
	}

	if err := mr.memberService.DeleteMember(id); err != nil {
		models.CheckError(w, err.Error(), http.StatusNotFound) // 500 status code
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 status code
	return
}
