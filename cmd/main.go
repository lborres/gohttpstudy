package main

import (
	"context"
	"fmt"
	"log"

	"github.com/lborres/gohttpstudy/api"
	"github.com/lborres/gohttpstudy/internal/config"
	"github.com/lborres/gohttpstudy/internal/db"
)

func run(_ context.Context, cfg config.Config) error {
	// TODO implement graceful shutdown,etc
	// ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	// defer cancel()

	addr := fmt.Sprintf("%s:%s", cfg.PublicHost, cfg.ServerPort)

	db, err := db.InitDB(*cfg.PGConfig)
	if err != nil {
		log.Fatal(err)
	}

	if err := api.StartAPIServer(addr, db); err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	log.Println("Starting HTTP Server")

	ctx := context.Background()
	if err := run(ctx, config.InitConfig()); err != nil {
		log.Fatal(err)
	}
}
