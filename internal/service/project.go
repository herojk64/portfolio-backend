package service

import (
	"context"

	"github.com/herojk64/portfolio-backend/internal/database/sqlc"
	"github.com/herojk64/portfolio-backend/internal/pkg/convert"
	"github.com/jackc/pgx/v5/pgtype"
)

type ProjectService struct {
	q *sqlc.Queries
}

func NewProjectService(q *sqlc.Queries) *ProjectService {
	return &ProjectService{q: q}
}

type ProjectParams struct {
	Title        string
	Description  *string
	RepoUrl      *string
	LiveUrl      *string
	ImageUrl     *string
	IsFeatured   bool
	DisplayOrder int32
}

func (s *ProjectService) List(ctx context.Context, featured *bool, search string, limit, offset int32) ([]sqlc.Project, int64, error) {
	filter := pgtype.Bool{}
	if featured != nil {
		filter = pgtype.Bool{Bool: *featured, Valid: true}
	}

	projects, err := s.q.ListProjects(ctx, sqlc.ListProjectsParams{
		Column1: filter,
		Column2: search,
		Limit:   limit,
		Offset:  offset,
	})
	if err != nil {
		return nil, 0, err
	}

	total, err := s.q.CountProjects(ctx, sqlc.CountProjectsParams{
		Column1: filter,
		Column2: search,
	})
	if err != nil {
		return nil, 0, err
	}

	return projects, total, nil
}

func (s *ProjectService) Get(ctx context.Context, id pgtype.UUID) (sqlc.Project, []sqlc.Skill, error) {
	project, err := s.q.GetProjectByID(ctx, id)
	if err != nil {
		return sqlc.Project{}, nil, err
	}

	skills, err := s.q.ListProjectSkills(ctx, id)
	if err != nil {
		return sqlc.Project{}, nil, err
	}

	return project, skills, nil
}

func (s *ProjectService) Create(ctx context.Context, p ProjectParams) (sqlc.Project, error) {
	return s.q.CreateProject(ctx, sqlc.CreateProjectParams{
		ID:           convert.NewUUID(),
		Title:        p.Title,
		Description:  convert.ToText(p.Description),
		RepoUrl:      convert.ToText(p.RepoUrl),
		LiveUrl:      convert.ToText(p.LiveUrl),
		ImageUrl:     convert.ToText(p.ImageUrl),
		IsFeatured:   p.IsFeatured,
		DisplayOrder: p.DisplayOrder,
	})
}

func (s *ProjectService) Update(ctx context.Context, id pgtype.UUID, p ProjectParams) (sqlc.Project, error) {
	return s.q.UpdateProject(ctx, sqlc.UpdateProjectParams{
		ID:           id,
		Title:        p.Title,
		Description:  convert.ToText(p.Description),
		RepoUrl:      convert.ToText(p.RepoUrl),
		LiveUrl:      convert.ToText(p.LiveUrl),
		ImageUrl:     convert.ToText(p.ImageUrl),
		IsFeatured:   p.IsFeatured,
		DisplayOrder: p.DisplayOrder,
	})
}

func (s *ProjectService) Delete(ctx context.Context, id pgtype.UUID) error {
	return s.q.DeleteProject(ctx, id)
}

func (s *ProjectService) AddSkill(ctx context.Context, projectID, skillID pgtype.UUID) error {
	return s.q.AddProjectSkill(ctx, sqlc.AddProjectSkillParams{
		ProjectID: projectID,
		SkillID:   skillID,
	})
}

func (s *ProjectService) RemoveSkill(ctx context.Context, projectID, skillID pgtype.UUID) error {
	return s.q.RemoveProjectSkill(ctx, sqlc.RemoveProjectSkillParams{
		ProjectID: projectID,
		SkillID:   skillID,
	})
}
