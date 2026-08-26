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
	"golang.org/x/crypto/bcrypt"
)

type UserData struct {
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	IsSuper  bool    `json:"is_super"`
	Avatar   *string `json:"avatar"`
}

func (s *Seeder) SeedAdmin() error {
	ctx := context.Background()

	userBytes, err := os.ReadFile("./internal/seed/data/users.json")
	if err != nil {
		log.Println("Error reading user json")
		return err
	}

	var users []UserData
	err = json.Unmarshal(userBytes, &users)
	if err != nil {
		return err
	}

	avatar := pgtype.Text{
		Valid: false,
	}

	for i := 0; i < len(users); i++ {
		_, err := s.q.GetUserByUsername(ctx, users[i].Username)
		if err == nil {
			continue
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(users[i].Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		id := uuid.New()

		pgid := pgtype.UUID{
			Bytes: id,
			Valid: true,
		}

		_, err = s.q.CreateUser(ctx, sqlc.CreateUserParams{
			ID:           pgid,
			Name:         users[i].Name,
			Username:     users[i].Username,
			Email:        users[i].Email,
			Avatar:       avatar,
			PasswordHash: string(hash),
			IsSuper:      users[i].IsSuper,
		})
		if err != nil {
			fmt.Println("Something went wrong while seeding data: ", err)
			return err
		}
	}

	return err
}
