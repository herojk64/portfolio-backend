// Package seed provides seeding for database settings
package seed

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/herojk64/portfolio/internal/database/sqlc"
)

type SettingsData struct {
	KEY   string          `json:"key"`
	VALUE json.RawMessage `json:"value"`
}

func (s *Seeder) SeedSettings() error {
	ctx := context.Background()

	settingsBytes, err := os.ReadFile("./internal/seed/data/settings.json")
	if err != nil {
		fmt.Println("Error reading the settings json data: ", err)
		return err
	}

	var settingsData []SettingsData

	err = json.Unmarshal(settingsBytes, &settingsData)
	if err != nil {
		fmt.Println("Error unmarshaling json data:", err)
		return err
	}

	for i := 0; i < len(settingsData); i++ {
		_, err := s.q.GetSettingsByKey(ctx, settingsData[i].KEY)
		if err == nil {
			continue
		}

		_, err = s.q.CreateSettings(ctx, sqlc.CreateSettingsParams{
			Key:   settingsData[i].KEY,
			Value: settingsData[i].VALUE,
		})
		if err != nil {
			fmt.Println("Something went wrong: ", err)
			return err
		}

	}

	return err
}
