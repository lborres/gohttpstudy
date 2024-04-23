package main

import (
	"fmt"
	"log"

	"github.com/lborres/gohttpstudy/api"
	"github.com/lborres/gohttpstudy/internal/config"
	"github.com/lborres/gohttpstudy/internal/db"
)

func run(cfg config.Config) error {
	addr := fmt.Sprintf("%s:%s", cfg.PublicHost, cfg.ServerPort)

	db, err := db.InitDB(*cfg.PGConfig)
	if err != nil {
		log.Fatal(err)
	}

	if err := api.ServeAPI(addr, db); err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	log.Println("Starting HTTP Server")
	if err := run(config.InitConfig()); err != nil {
		log.Fatal(err)
	}
}
