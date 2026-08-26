package service

import (
	"context"

	"github.com/herojk64/portfolio-backend/internal/database/sqlc"
	"github.com/herojk64/portfolio-backend/internal/pkg/convert"
	"github.com/jackc/pgx/v5/pgtype"
)

type SkillService struct {
	q *sqlc.Queries
}

func NewSkillService(q *sqlc.Queries) *SkillService {
	return &SkillService{q: q}
}

type SkillParams struct {
	Name         string
	Category     *string
	Proficiency  *string
	Icon         *string
	DisplayOrder int32
}

func (s *SkillService) List(ctx context.Context, category *string, search string, limit, offset int32) ([]sqlc.Skill, int64, error) {
	filter := pgtype.Text{}
	if category != nil {
		filter = pgtype.Text{String: *category, Valid: true}
	}

	skills, err := s.q.ListSkills(ctx, sqlc.ListSkillsParams{
		Column1: filter,
		Column2: search,
		Limit:   limit,
		Offset:  offset,
	})
	if err != nil {
		return nil, 0, err
	}

	total, err := s.q.CountSkills(ctx, sqlc.CountSkillsParams{
		Column1: filter,
		Column2: search,
	})
	if err != nil {
		return nil, 0, err
	}

	return skills, total, nil
}

func (s *SkillService) Get(ctx context.Context, id pgtype.UUID) (sqlc.Skill, error) {
	return s.q.GetSkillByID(ctx, id)
}

func (s *SkillService) Create(ctx context.Context, p SkillParams) (sqlc.Skill, error) {
	return s.q.CreateSkill(ctx, sqlc.CreateSkillParams{
		ID:           convert.NewUUID(),
		Name:         p.Name,
		Category:     convert.ToText(p.Category),
		Proficiency:  convert.ToText(p.Proficiency),
		Icon:         convert.ToText(p.Icon),
		DisplayOrder: p.DisplayOrder,
	})
}

func (s *SkillService) Update(ctx context.Context, id pgtype.UUID, p SkillParams) (sqlc.Skill, error) {
	return s.q.UpdateSkill(ctx, sqlc.UpdateSkillParams{
		ID:           id,
		Name:         p.Name,
		Category:     convert.ToText(p.Category),
		Proficiency:  convert.ToText(p.Proficiency),
		Icon:         convert.ToText(p.Icon),
		DisplayOrder: p.DisplayOrder,
	})
}

func (s *SkillService) Delete(ctx context.Context, id pgtype.UUID) error {
	return s.q.DeleteSkill(ctx, id)
}
