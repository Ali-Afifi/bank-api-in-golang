package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// serverFunction is the server handlers form.
type serverFunction func(w http.ResponseWriter, r *http.Request) error

type serverError struct {
	Error string
}

// renderJSON renders 'v' as JSON and writes it as a response into w.
func renderJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

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

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Run() {
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/account/", makeHTTPHandleFunc(s.getAccountHandler)).Methods("GET")
	router.HandleFunc("/account/", makeHTTPHandleFunc(s.updateAccountHandler)).Methods("PUT")
	router.HandleFunc("/account/", makeHTTPHandleFunc(s.createAccountHandler)).Methods("POST")
	router.HandleFunc("/account/", makeHTTPHandleFunc(s.deleteAccountHandler)).Methods("DELETE")

	log.Println("Server is currently running on ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *Server) createAccountHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) getAccountHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) updateAccountHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) deleteAccountHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}
