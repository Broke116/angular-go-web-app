package main

import (
	"angular-go-web-app/go/server"
	"angular-go-web-app/go/services"
	"angular-go-web-app/go/utils/security"
	"flag"
	"log"
	"net/http"

	"angular-go-web-app/go/config/mongo"
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

	ms, err := mongo.NewSession("mongodb://db:27017") // do not hard code the server address

	if err != nil {
		log.Fatalln("unable to connect to mongodb")
	}

	defer ms.Close()

	h := security.Hash{}
	memberService := services.MemberServiceConstructor(ms.Copy(), "airline", "members", h)
	userService := services.UserServiceConstructor(ms.Copy(), "airline", "users", h)
	s := server.NewServer(memberService, userService)
	s.Start()

	/*r.Handle("/status", mw(statusHandler))
	r.HandleFunc("/", Index).Methods("GET")
	r.PathPrefix("/static").Handler(http.StripPrefix("/static",
		http.FileServer(http.Dir("./assets"))))*/
}
