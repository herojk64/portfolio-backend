package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/herojk64/portfolio-backend/internal/database/sqlc"
	"github.com/jackc/pgx/v5"
)

type SettingService struct {
	q *sqlc.Queries
}

func NewSettingService(q *sqlc.Queries) *SettingService {
	return &SettingService{q: q}
}

func (s *SettingService) List(ctx context.Context) ([]sqlc.Setting, error) {
	return s.q.GetSettings(ctx, sqlc.GetSettingsParams{
		Limit:  100,
		Offset: 0,
	})
}

func (s *SettingService) Get(ctx context.Context, key string) (sqlc.Setting, error) {
	return s.q.GetSettingsByKey(ctx, key)
}

func (s *SettingService) Upsert(ctx context.Context, key string, value json.RawMessage) (sqlc.Setting, error) {
	existing, err := s.q.GetSettingsByKey(ctx, key)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			row, createErr := s.q.CreateSettings(ctx, sqlc.CreateSettingsParams{
				Key:   key,
				Value: value,
			})
			if createErr != nil {
				return sqlc.Setting{}, createErr
			}
			return sqlc.Setting{Key: row.Key, Value: row.Value}, nil
		}
		return sqlc.Setting{}, err
	}

	_ = existing
	return s.q.UpdateSettings(ctx, sqlc.UpdateSettingsParams{
		Key:   key,
		Value: value,
	})
}

func (s *SettingService) Update(ctx context.Context, key string, value json.RawMessage) (sqlc.Setting, error) {
	return s.q.UpdateSettings(ctx, sqlc.UpdateSettingsParams{
		Key:   key,
		Value: value,
	})
}

func (s *SettingService) Delete(ctx context.Context, key string) error {
	return s.q.DeleteSettings(ctx, key)
}
