package main

import (
	"flag"
	"log"
	"net/http"

	"angular-go-web-app/go/routes/MemberAPI/v1"

	"github.com/gorilla/mux"
)

var addr = flag.String("addr", ":9090", "http service address")

//var tmpl = template.Must(template.ParseFiles("./assets/public/html/index.html"))

// Index returning the landing page
func Index(w http.ResponseWriter, req *http.Request) {
	/*data := models.IndexPageData{
		PageTitle: "API documentation",
		List: []models.ListItem{
			{Title: "Member", URL: "http://localhost:9090/v1/member"},
		},
	}*/
	/*err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}*/
}

func main() {
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/", Index).Methods("GET")
	router.Handle("/static/", http.StripPrefix("/static",
		http.FileServer(http.Dir("./assets"))))
	router.HandleFunc("/v1/member", apimember.GetMembersEndpoint).Methods("GET")
	router.HandleFunc("/v1/member/{id}", apimember.GetMemberEndpoint).Methods("GET")
	router.HandleFunc("/v1/member/{id}", apimember.CreateMemberEndpoint).Methods("POST")
	router.HandleFunc("/v1/member/{id}", apimember.DeleteMemberEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(*addr, router))
}
