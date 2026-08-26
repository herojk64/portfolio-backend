package config

import (
	"github.com/spf13/viper"
)

func Load() (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")

	_ = v.BindEnv("database.host", "DATABASE_HOST")
	_ = v.BindEnv("database.port", "DATABASE_PORT")
	_ = v.BindEnv("database.user", "DATABASE_USER")
	_ = v.BindEnv("database.password", "DATABASE_PASSWORD")
	_ = v.BindEnv("database.name", "DATABASE_NAME")
	_ = v.BindEnv("redis.host", "REDIS_HOST")
	_ = v.BindEnv("redis.port", "REDIS_PORT")
	_ = v.BindEnv("app.env", "APP_ENV")
	_ = v.BindEnv("session.secret", "SESSION_SECRET")
	_ = v.BindEnv("csrf.secret", "CSRF_SECRET")

	var cfg Config

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
