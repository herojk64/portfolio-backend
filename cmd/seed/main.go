package main

import (
	"log"

	"github.com/herojk64/portfolio/internal/config"
	"github.com/herojk64/portfolio/internal/database"
	"github.com/herojk64/portfolio/internal/database/sqlc"
	"github.com/herojk64/portfolio/internal/seed"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	queries := sqlc.New(db)
	seeder := seed.New(queries)

	if err := seeder.Run(); err != nil {
		log.Fatal(err)
	}
}
