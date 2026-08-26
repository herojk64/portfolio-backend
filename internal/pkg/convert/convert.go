package convert

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ToText(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{}
	}
	return pgtype.Text{String: *s, Valid: true}
}

func NewUUID() pgtype.UUID {
	id := uuid.New()
	return pgtype.UUID{Bytes: id, Valid: true}
}

func ParseUUID(s string) (pgtype.UUID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return pgtype.UUID{}, err
	}
	return pgtype.UUID{Bytes: id, Valid: true}, nil
}
