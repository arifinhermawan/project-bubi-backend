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
func HandleRequest(handlers *server.Handlers) {
	router := mux.NewRouter().StrictSlash(true)

	handleGetRequest(handlers, router)
	handlePostRequest(handlers, router)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// handleGetRequest will handle request with type GET
func handleGetRequest(handlers *server.Handlers, router *mux.Router) {
}

// handlePostRequest will handle request with type POST
func handlePostRequest(handlers *server.Handlers, router *mux.Router) {
	// account
	router.HandleFunc("/account/signup", handlers.Account.HandleUserSignUp).Methods("POST")
}
