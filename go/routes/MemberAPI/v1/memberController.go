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

type memberController struct {
	memberService *member.MemberService
}

// MemberControllerConstructor is used to initialize member api and define its routes.
func MemberControllerConstructor(ms *mongo.MemberService, r *mux.Router) *mux.Router {
	memberController := memberController{ms}

	mw := middlewares.ChainMiddleware(middlewares.Logging, middlewares.Tracing)

	r.HandleFunc("/", mw(memberController.GetMembersEndpoint)).Methods("GET")
	r.HandleFunc("/insertMember", mw(memberController.InsertMemberEndpoint)).Methods("POST")
	r.HandleFunc("/updateMember/{id}", mw(memberController.UpdateMemberEndpoint)).Methods("PUT")
	r.HandleFunc("/{id}", mw(memberController.GetMemberEndpoint)).Methods("GET")
	r.HandleFunc("/{id}", mw(memberController.DeleteMemberEndpoint)).Methods("DELETE")

	return r
}

// GetMembersEndpoint returns all users.
func (mc *memberController) GetMembersEndpoint(w http.ResponseWriter, req *http.Request) {
	members, err := mc.memberService.GetMembers()

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
func (mc *memberController) GetMemberEndpoint(w http.ResponseWriter, req *http.Request) {
	// get id parameter
	vars := mux.Vars(req)
	id := vars["id"]

	// check if the id is valid
	if !bson.IsObjectIdHex(id) {
		models.CheckError(w, "Invalid ObjectId", http.StatusNotFound) // 404 status code
	}

	member, err := mc.memberService.GetMemberByID(id)
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
func (mc *memberController) InsertMemberEndpoint(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	var member models.Member
	if err := json.NewDecoder(req.Body).Decode(&member); err != nil { // decode body
		models.CheckError(w, err.Error(), http.StatusInternalServerError) // 500 status code
	}

	if err := mc.memberService.InsertMember(&member); err != nil {
		models.CheckError(w, err.Error(), http.StatusInternalServerError) // 500 status code
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return
}

// UpdateMemberEndpoint - v1/updateMember PUT - is used to update the related member
func (mc *memberController) UpdateMemberEndpoint(w http.ResponseWriter, req *http.Request) {
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

	if err := mc.memberService.UpdateMember(&member, id); err != nil {
		models.CheckError(w, err.Error(), http.StatusNotFound) // 500 status code
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // status code
	return
}

// DeleteMemberEndpoint - DELETE member/{id} / member - deletes the member
func (mc *memberController) DeleteMemberEndpoint(w http.ResponseWriter, req *http.Request) {
	// get id paramtere
	vars := mux.Vars(req)
	id := vars["id"]

	// check if the id is valid
	if !bson.IsObjectIdHex(id) {
		models.CheckError(w, "Invalid ObjectId", http.StatusNotFound) // 404 status code
	}

	if err := mc.memberService.DeleteMember(id); err != nil {
		models.CheckError(w, err.Error(), http.StatusNotFound) // 500 status code
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 status code
	return
}
