package seed

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/herojk64/portfolio-backend/internal/database/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type ExperienceData struct {
	Company     string `json:"company"`
	Role        string `json:"role"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Location    string `json:"location"`
}

func (s *Seeder) SeedExperience() error {
	ctx := context.Background()

	data, err := os.ReadFile("./internal/seed/data/experience.json")
	if err != nil {
		fmt.Println("Error reading experience json:", err)
		return err
	}

	var experiences []ExperienceData
	if err = json.Unmarshal(data, &experiences); err != nil {
		return err
	}

	for _, exp := range experiences {
		count, err := s.q.CountExperience(ctx, sqlc.CountExperienceParams{
			Column1: pgtype.Text{String: exp.Company, Valid: true},
			Column2: pgtype.Text{String: exp.Role, Valid: true},
		})
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}

		endDate := pgtype.Text{Valid: false}
		if exp.EndDate != "" {
			endDate = pgtype.Text{String: exp.EndDate, Valid: true}
		}

		_, err = s.q.CreateExperience(ctx, sqlc.CreateExperienceParams{
			ID:          pgtype.UUID{Bytes: uuid.New(), Valid: true},
			Company:     exp.Company,
			Role:        exp.Role,
			Description: pgtype.Text{String: exp.Description, Valid: exp.Description != ""},
			StartDate:   exp.StartDate,
			EndDate:     endDate,
			Location:    pgtype.Text{String: exp.Location, Valid: exp.Location != ""},
		})
		if err != nil {
			fmt.Println("Error seeding experience:", err)
			return err
		}
	}

	return nil
}
