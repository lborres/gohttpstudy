package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/lborres/gohttpstudy/service/cart"
	"github.com/lborres/gohttpstudy/service/misc"
	"github.com/lborres/gohttpstudy/service/product"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func StartAPIServer(addr string, db *sql.DB) error {
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
	// Routers via go 1.22 http.NewServeMux()
	mux := http.NewServeMux()

	log.Println("Initializing Server Routes")

	miscHandler := misc.NewHandler()
	miscHandler.RegisterRoutes(mux)

	productStorage := product.NewStorage(server.db)
	productHandler := product.NewHandler(productStorage)
	productHandler.RegisterRoutes(mux)

	cartStorage := cart.NewStorage(server.db)
	cartHandler := cart.NewHandler(cartStorage)
	cartHandler.RegisterRoutes(mux)

	log.Printf("HTTP Server listening at %s\n", server.addr)

	return http.ListenAndServe(server.addr, mux)
}
