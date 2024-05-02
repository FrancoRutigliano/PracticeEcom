package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/FrancoRutigliano/pkg/middleware"
)

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
	// enrutador principal
	router := http.NewServeMux()
	//enrutador secundario - v1
	v1 := http.NewServeMux()
	v1.Handle("api/v1/", http.StripPrefix("/api/v1", router))

	//middleware chain
	middleware := middleware.MiddlewareChain()

	// LOGICA

	log.Println("Listening on port: ", s.addr)

	server := http.Server{
		Addr:    s.addr,
		Handler: middleware(v1),
	}

	return server.ListenAndServe()
}
