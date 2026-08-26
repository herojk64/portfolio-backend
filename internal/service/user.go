package service

import (
	"context"

	"github.com/herojk64/portfolio/internal/database/sqlc"
)

type UserService struct {
	q *sqlc.Queries
}

func NewUserService(q *sqlc.Queries) *UserService {
	return &UserService{q: q}
}

func (s *UserService) ListUsers(ctx context.Context, params sqlc.SelectUsersParams) ([]sqlc.SelectUsersRow, error) {
	return s.q.SelectUsers(ctx, params)
}
