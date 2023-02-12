package utils

import (
	// golang package
	"log"
	"net/http"

	// external package
	"github.com/gorilla/mux"

	// internal package
	"github.com/arifinhermawan/bubi/internal/app/server"
)

// HandleRequest handles all incoming request to backend.
func HandleRequest(infra *server.Infra, handlers *server.Handlers) {
	router := mux.NewRouter().StrictSlash(true)

	handleGetRequest(infra, handlers, router)
	handlePatchRequest(infra, handlers, router)
	handlePostRequest(infra, handlers, router)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// handleGetRequest will handle request with type GET
func handleGetRequest(infra *server.Infra, handlers *server.Handlers, router *mux.Router) {
}

// handlePatchRequest will handle request with type PATCH
func handlePatchRequest(infra *server.Infra, handlers *server.Handlers, router *mux.Router) {
	// account
	router.HandleFunc("/account/update", infra.Auth.JWTAuthorization(handlers.Account.HandleUpdateUserAccount)).Methods("PATCH")
	router.HandleFunc("/account/update_password", infra.Auth.JWTAuthorization(handlers.Account.HandleUpdateUserPassword)).Methods("PATCH")
}

// handlePostRequest will handle request with type POST
func handlePostRequest(infra *server.Infra, handlers *server.Handlers, router *mux.Router) {
	// account
	router.HandleFunc("/account/login", handlers.Account.HandleUserLogIn).Methods("POST")
	router.HandleFunc("/account/logout", handlers.Account.HandlerUserLogOut).Methods("POST")
	router.HandleFunc("/account/signup", handlers.Account.HandleUserSignUp).Methods("POST")
}
