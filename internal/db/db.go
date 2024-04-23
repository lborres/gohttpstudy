package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/lborres/gohttpstudy/internal/config"
)

func InitDB(cfg config.PGConfig) (*sql.DB, error) {
	log.Println("Initializing Database Connection...")

	db, err := sql.Open("postgres", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Printf("...Database Connection Established")

	return db, nil
}
