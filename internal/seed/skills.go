package seed

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/herojk64/portfolio-backend/internal/database/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type SkillsData struct {
	Name         string `json:"name"`
	Category     string `json:"category"`
	Proficiency  string `json:"proficiency"`
	Icon         string `json:"icon"`
	DisplayOrder int32  `json:"display_order"`
}

func (s *Seeder) SeedSkills() error {
	ctx := context.Background()

	skillsBytes, err := os.ReadFile("./internal/seed/data/skills.json")
	if err != nil {
		log.Println("Error reading user json")
		return err
	}

	var skills []SkillsData
	err = json.Unmarshal(skillsBytes, &skills)
	if err != nil {
		return err
	}

	for i := 0; i < len(skills); i++ {
		_, err := s.q.GetSkillByName(ctx, skills[i].Name)
		if err == nil {
			continue
		}

		id := uuid.New()

		pgid := pgtype.UUID{
			Bytes: id,
			Valid: true,
		}

		category := pgtype.Text{
			String: skills[i].Category,
			Valid:  skills[i].Category != "",
		}

		proficiency := pgtype.Text{
			String: skills[i].Proficiency,
			Valid:  skills[i].Proficiency != "",
		}

		icon := pgtype.Text{
			String: skills[i].Icon,
			Valid:  skills[i].Icon != "",
		}

		_, err = s.q.CreateSkill(ctx, sqlc.CreateSkillParams{
			ID:           pgid,
			Name:         skills[i].Name,
			Category:     category,
			Proficiency:  proficiency,
			Icon:         icon,
			DisplayOrder: skills[i].DisplayOrder,
		})
		if err != nil {
			fmt.Println("Something went wrong while seeding skills data: ", err)
			return err
		}
	}

	return err
}
