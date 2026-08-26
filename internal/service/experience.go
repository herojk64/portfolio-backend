package service

import (
	"context"

	"github.com/herojk64/portfolio/internal/database/sqlc"
	"github.com/herojk64/portfolio/internal/pkg/convert"
	"github.com/jackc/pgx/v5/pgtype"
)

type ExperienceService struct {
	q *sqlc.Queries
}

func NewExperienceService(q *sqlc.Queries) *ExperienceService {
	return &ExperienceService{q: q}
}

type ExperienceParams struct {
	Company     string
	Role        string
	Description *string
	StartDate   string
	EndDate     *string
	Location    *string
}

func (s *ExperienceService) List(ctx context.Context, company, role, search string, limit, offset int32) ([]sqlc.Experience, int64, error) {
	// search filters against both company and role columns
	companyFilter := pgtype.Text{}
	roleFilter := pgtype.Text{}
	if search != "" {
		companyFilter = pgtype.Text{String: search, Valid: true}
		roleFilter = pgtype.Text{String: search, Valid: true}
	}
	if company != "" {
		companyFilter = pgtype.Text{String: company, Valid: true}
	}
	if role != "" {
		roleFilter = pgtype.Text{String: role, Valid: true}
	}

	items, err := s.q.ListExperience(ctx, sqlc.ListExperienceParams{
		Column1: companyFilter,
		Column2: roleFilter,
		Limit:   limit,
		Offset:  offset,
	})
	if err != nil {
		return nil, 0, err
	}

	total, err := s.q.CountExperience(ctx, sqlc.CountExperienceParams{
		Column1: companyFilter,
		Column2: roleFilter,
	})
	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (s *ExperienceService) Get(ctx context.Context, id pgtype.UUID) (sqlc.Experience, []sqlc.Skill, error) {
	exp, err := s.q.GetExperienceByID(ctx, id)
	if err != nil {
		return sqlc.Experience{}, nil, err
	}

	skills, err := s.q.ListExperienceSkills(ctx, id)
	if err != nil {
		return sqlc.Experience{}, nil, err
	}

	return exp, skills, nil
}

func (s *ExperienceService) Create(ctx context.Context, p ExperienceParams) (sqlc.Experience, error) {
	return s.q.CreateExperience(ctx, sqlc.CreateExperienceParams{
		ID:          convert.NewUUID(),
		Company:     p.Company,
		Role:        p.Role,
		Description: convert.ToText(p.Description),
		StartDate:   p.StartDate,
		EndDate:     convert.ToText(p.EndDate),
		Location:    convert.ToText(p.Location),
	})
}

func (s *ExperienceService) Update(ctx context.Context, id pgtype.UUID, p ExperienceParams) (sqlc.Experience, error) {
	return s.q.UpdateExperience(ctx, sqlc.UpdateExperienceParams{
		ID:          id,
		Company:     p.Company,
		Role:        p.Role,
		Description: convert.ToText(p.Description),
		StartDate:   p.StartDate,
		EndDate:     convert.ToText(p.EndDate),
		Location:    convert.ToText(p.Location),
	})
}

func (s *ExperienceService) Delete(ctx context.Context, id pgtype.UUID) error {
	return s.q.DeleteExperience(ctx, id)
}

func (s *ExperienceService) AddSkill(ctx context.Context, experienceID, skillID pgtype.UUID) error {
	return s.q.AddExperienceSkill(ctx, sqlc.AddExperienceSkillParams{
		ExperienceID: experienceID,
		SkillID:      skillID,
	})
}

func (s *ExperienceService) RemoveSkill(ctx context.Context, experienceID, skillID pgtype.UUID) error {
	return s.q.RemoveExperienceSkill(ctx, sqlc.RemoveExperienceSkillParams{
		ExperienceID: experienceID,
		SkillID:      skillID,
	})
}
