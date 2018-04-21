package apimember

import (
	"angular-go-web-app/go/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var members = make(map[string]models.Member)

func init() {
	members["1"] = models.Member{ID: "1", Firstname: "Kara", Lastname: "Murat",
		Address: &models.Address{City: "Dublin", Country: "Ireland"}}

	members["2"] = models.Member{ID: "2", Firstname: "Arif", Lastname: "Isik"}
}

// GetMembersEndpoint returns all users.
func GetMembersEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(members)
	fmt.Println(req.Method, req.URL)
}

// GetMemberEndpoint returns a specific user
func GetMemberEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	member, _ := members[params["id"]]

	if member.ID != "" {
		fmt.Println(req.Method, req.URL)
		json.NewEncoder(w).Encode(member)
	} else {
		json.NewEncoder(w).Encode(&models.Error{Errortype: "member not found", Statuscode: 404})
	}
}

// CreateMemberEndpoint creates a new member
func CreateMemberEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	var member models.Member
	_ = json.NewDecoder(req.Body).Decode(&member)

	member.ID = params["id"]
	members[member.ID] = member

	fmt.Println(req.Method, req.URL)
	json.NewEncoder(w).Encode(members)
}

// DeleteMemberEndpoint deletes the member
func DeleteMemberEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := params["id"]

	if id != "" {
		delete(members, id)
		fmt.Println(req.Method, req.URL)
		json.NewEncoder(w).Encode(members)
	} else {
		json.NewEncoder(w).Encode(&models.Error{Errortype: "member not found", Statuscode: 404})
	}
}
