package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
	dataStore  DataStore
}

func NewServer(listenAddr string, dataStore DataStore) *Server {

	return &Server{
		listenAddr: listenAddr,
		dataStore:  dataStore,
	}
}

func (s *Server) Run() {
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/account/", makeHTTPHandleFunc(s.getAllAccountsHandler)).Methods("GET")
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.getAccountHandler)).Methods("GET")
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.updateAccountHandler)).Methods("PUT")
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.createAccountHandler)).Methods("POST")
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.deleteAccountHandler)).Methods("DELETE")

	log.Println("Server is currently running on ", s.listenAddr)

	log.Fatal(http.ListenAndServe(s.listenAddr, router))

}

func (s *Server) createAccountHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) getAllAccountsHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) getAccountHandler(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	acc := NewAccount("Adam", "Smith")

	idInt, _ := strconv.Atoi(id)

	acc.SetID(idInt)

	renderJSON(w, http.StatusOK, acc)

	return nil
}

func (s *Server) updateAccountHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) deleteAccountHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// serverFunction is the server handlers' form.
type serverFunction func(w http.ResponseWriter, r *http.Request) error

// serverError is the error type used across the server.
type serverError struct {
	Error string
}

// renderJSON renders 'v' as JSON and writes it as a response into w.
func renderJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	js, err := json.Marshal(v)

	if err != nil {

		log.Printf("An error occurred while encoding json: %v\n", err.Error())

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

// makeHTTPHandleFunc calls a handler function of type serverFunction
// and returns a function of type http.HandlerFunc
// which is the required type for mux's Router's method HandleFunc().
func makeHTTPHandleFunc(f serverFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			renderJSON(w, http.StatusBadRequest, serverError{Error: err.Error()})

		}
	}
}
