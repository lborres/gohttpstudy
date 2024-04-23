package api

import (
	"database/sql"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func ServeAPI(addr string, db *sql.DB) error {
	apiserver := &APIServer{
		addr: addr,
		db:   db,
	}

	if err := apiserver.initRoutes(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (server *APIServer) initRoutes() error {
	router := http.NewServeMux()

	// TODO Emulate subrouter available with Gorilla Mux
	// subrouter := router.PathPrefix("/api/v1").Subrouter()

	log.Println("Initializing Server Routes")

	// 

	log.Printf("HTTP Server listening at %s\n", server.addr)

	return http.ListenAndServe(server.addr, router)
}
