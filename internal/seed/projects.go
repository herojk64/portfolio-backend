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

type ProjectData struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	RepoURL      string `json:"repo_url"`
	LiveURL      string `json:"live_url"`
	ImageURL     string `json:"image_url"`
	IsFeatured   bool   `json:"is_featured"`
	DisplayOrder int32  `json:"display_order"`
}

func (s *Seeder) SeedProjects() error {
	ctx := context.Background()

	data, err := os.ReadFile("./internal/seed/data/projects.json")
	if err != nil {
		fmt.Println("Error reading projects json:", err)
		return err
	}

	var projects []ProjectData
	if err = json.Unmarshal(data, &projects); err != nil {
		return err
	}

	for _, proj := range projects {
		count, err := s.q.CountProjects(ctx, sqlc.CountProjectsParams{
			Column1: pgtype.Bool{Valid: false},
			Column2: proj.Title,
		})
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}

		_, err = s.q.CreateProject(ctx, sqlc.CreateProjectParams{
			ID:           pgtype.UUID{Bytes: uuid.New(), Valid: true},
			Title:        proj.Title,
			Description:  pgtype.Text{String: proj.Description, Valid: proj.Description != ""},
			RepoUrl:      pgtype.Text{String: proj.RepoURL, Valid: proj.RepoURL != ""},
			LiveUrl:      pgtype.Text{String: proj.LiveURL, Valid: proj.LiveURL != ""},
			ImageUrl:     pgtype.Text{String: proj.ImageURL, Valid: proj.ImageURL != ""},
			IsFeatured:   proj.IsFeatured,
			DisplayOrder: proj.DisplayOrder,
		})
		if err != nil {
			fmt.Println("Error seeding project:", err)
			return err
		}
	}

	return nil
}
