package main

import (
	"flag"
	"log"
	"net/http"

	"angular-go-web-app/go/routes/MemberAPI/v1"
	"angular-go-web-app/go/utils/middlewares"

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

var statusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})

func main() {
	flag.Parse()

	mw := middlewares.ChainMiddleware(middlewares.Logging, middlewares.Tracing)

	r := mux.NewRouter()
	r.Handle("/status", mw(statusHandler))
	r.HandleFunc("/", Index).Methods("GET")
	r.PathPrefix("/static").Handler(http.StripPrefix("/static",
		http.FileServer(http.Dir("./assets"))))
	r.HandleFunc("/v1/member", mw(apimember.GetMembersEndpoint)).Methods("GET")
	r.HandleFunc("/v1/insertMember", mw(apimember.InsertMemberEndpoint)).Methods("POST")
	r.HandleFunc("/v1/updateMember/{id}", mw(apimember.UpdateMemberEndpoint)).Methods("PUT")
	r.HandleFunc("/v1/member/{id}", mw(apimember.GetMemberEndpoint)).Methods("GET")
	r.HandleFunc("/v1/member", mw(apimember.DeleteMemberEndpoint)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(*addr, r))
}
