package main

import (
	"log"

	"github.com/herojk64/portfolio-backend/internal/config"
	"github.com/herojk64/portfolio-backend/internal/database"
	"github.com/herojk64/portfolio-backend/internal/database/sqlc"
	"github.com/herojk64/portfolio-backend/internal/seed"
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
