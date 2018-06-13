package main

import (
	"angular-go-web-app/go/server"
	"flag"
	"log"
	"net/http"

	"angular-go-web-app/go/config/mongo"
	"angular-go-web-app/go/config/mongo/member"
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

	ms, err := mongo.NewSession("mongodb://db:27017")

	if err != nil {
		log.Fatalln("unable to connect to mongodb")
	}

	defer ms.Close()

	memberService := member.MemberServiceInstance(ms.Copy(), "airline", "members")
	s := server.NewServer(memberService)
	s.Start()

	/*r.Handle("/status", mw(statusHandler))
	r.HandleFunc("/", Index).Methods("GET")
	r.PathPrefix("/static").Handler(http.StripPrefix("/static",
		http.FileServer(http.Dir("./assets"))))*/
}
