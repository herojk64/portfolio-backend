package seed

import "github.com/herojk64/portfolio-backend/internal/database/sqlc"

type Seeder struct {
	q *sqlc.Queries
}

func New(q *sqlc.Queries) *Seeder {
	return &Seeder{
		q: q,
	}
}

func (s *Seeder) Run() error {
	s.SeedAdmin()
	s.SeedSettings()
	s.SeedSkills()
	s.SeedExperience()
	s.SeedProjects()
	return nil
}
