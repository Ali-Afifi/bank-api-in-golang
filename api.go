package main

import (
	"encoding/json"
	"fmt"
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
	router.HandleFunc("/account/", makeHTTPHandleFunc(s.createAccountHandler)).Methods("POST")
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.getAccountHandler)).Methods("GET")
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.updateAccountHandler)).Methods("PUT")
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.deleteAccountHandler)).Methods("DELETE")

	log.Println("Server is currently running on ", s.listenAddr)

	log.Fatal(http.ListenAndServe(s.listenAddr, router))

}

func (s *Server) createAccountHandler(w http.ResponseWriter, r *http.Request) error {
	createAccountRequestBody := CreateAccountRequestBody{}

	if err := json.NewDecoder(r.Body).Decode(&createAccountRequestBody); err != nil {
		return err
	}

	account := NewAccount(createAccountRequestBody.FirstName, createAccountRequestBody.LastName)

	if err := s.dataStore.CreateAccount(account); err != nil {
		return err
	}

	renderJSON(w, http.StatusOK, account)

	return nil
}

func (s *Server) getAllAccountsHandler(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.dataStore.GetAllAccounts()

	if err != nil {
		return nil
	}

	renderJSON(w, http.StatusOK, accounts)

	return nil
}

func (s *Server) getAccountHandler(w http.ResponseWriter, r *http.Request) error {

	id, err := getID(r)

	if err != nil {
		return err
	}

	account, dbErr := s.dataStore.GetAccountByID(id)

	if dbErr != nil {
		return dbErr
	}

	renderJSON(w, http.StatusOK, account)

	return nil
}

func (s *Server) updateAccountHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) deleteAccountHandler(w http.ResponseWriter, r *http.Request) error {

	id, err := getID(r)

	if err != nil {
		return err
	}

	if err := s.dataStore.DeleteAccount(id); err != nil {
		return err
	}

	renderJSON(w, http.StatusOK, map[string]int{"id": id})

	return nil
}

// serverFunction is the server handlers' form.
type serverFunction func(w http.ResponseWriter, r *http.Request) error

// serverError is the error type used across the server.
type serverError struct {
	Error string `json:"error"`
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

func getID(r *http.Request) (int, error) {
	strID := mux.Vars(r)["id"]

	intID, err := strconv.Atoi(strID)

	if err != nil {
		return intID, fmt.Errorf("id: %s is invalid", strID)
	}

	return intID, nil
}
