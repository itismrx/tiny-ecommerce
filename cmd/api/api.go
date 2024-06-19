package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/itismrx/tiny-ecommerce/services/user"
)

// the entry point for the api

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	// User servcie
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Lisining on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
