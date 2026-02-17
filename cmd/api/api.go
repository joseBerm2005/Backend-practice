package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joseBerm2005/Backend-practice/service/user"
)

// Server struct that has address with type string
// db with a pointer to an sql database
type APIServer struct {
	addr string
	db	*sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

// Run function that creates three variables:
// router is responsible for handling all HTTP requests
// subrouter groups added routes that are related

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}