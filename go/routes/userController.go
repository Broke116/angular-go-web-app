package controller

import (
	"angular-go-web-app/go/models"
	"angular-go-web-app/go/services"
	"angular-go-web-app/go/utils/middlewares"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type userController struct {
	userService *services.UserService
}

// UserControllerConstructor is used to instantinate the user controller and create its endpoint paths
func UserControllerConstructor(us *services.UserService, r *mux.Router) *mux.Router {
	userController := userController{us}

	mw := middlewares.ChainMiddleware(middlewares.Logging, middlewares.Tracing)

	r.HandleFunc("/authenticate", mw(userController.Authenticate)).Methods("POST")
	r.HandleFunc("/createUser", mw(userController.CreateUserEndpoint)).Methods("POST")

	return r
}

func (uc *userController) Authenticate(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	var credential models.Credentials
	if err := json.NewDecoder(req.Body).Decode(&credential); err != nil { // decode body
		models.CheckError(w, err.Error(), http.StatusBadRequest) // 400 status code
	}

	user, err := uc.userService.Authenticate(credential)

	if err == nil {
		fmt.Println("Controller Authorized ", user)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		//w.Write(token)
		return
	}

	fmt.Println("Controller Unauthorized")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	return
}

// CreateUserEndpoint - POST / insertUser - creating a new user
func (uc *userController) CreateUserEndpoint(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	if req.Body == nil {
		//server.JSON(W, http.StatusNoContent, "Request must not be empty.")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("Request body must not be empty."))
		return
	}

	var user models.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil { // decode body
		models.CheckError(w, err.Error(), http.StatusBadRequest) // 400 status code
	}

	if err := uc.userService.CreateUser(&user); err != nil {
		models.CheckError(w, err.Error(), http.StatusInternalServerError) // 500 status code
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return
}
