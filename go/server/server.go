package server

import (
	"angular-go-web-app/go/routes"
	"angular-go-web-app/go/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//var addr = flag.String("addr", ":9090", "http service address")

// Server holds the values for server instantination
type Server struct {
	router *mux.Router
}

// NewServer is used to create a server
func NewServer(ms *services.MemberService, us *services.UserService) *Server {
	s := Server{router: mux.NewRouter()}
	controller.MemberControllerConstructor(ms, s.newSubRouter("/v1/member"))
	controller.UserControllerConstructor(us, s.newSubRouter("/v1/user"))
	return &s
}

// Start is used to start a server
func (s *Server) Start() {
	//flag.Parse()
	log.Println("Listening on port 9090")

	if err := http.ListenAndServe(":9090", s.router); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}

}

func (s *Server) newSubRouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}
