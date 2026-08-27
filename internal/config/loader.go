package config

import (
	"errors"

	"github.com/spf13/viper"
)

func Load() (*Config, error) {
	v := viper.New()

	// Defaults: used when no local config.yaml exists.
	v.SetDefault("app.name", "backend")
	v.SetDefault("app.env", "development")
	v.SetDefault("app.port", "8000")
	v.SetDefault("app.allowed_hosts", []string{"http://localhost:3000"})

	v.SetDefault("database.host", "localhost")
	v.SetDefault("database.port", 5432)
	v.SetDefault("database.user", "postgres")
	v.SetDefault("database.password", "")
	v.SetDefault("database.name", "backend")
	v.SetDefault("database.sslmode", "disable")

	v.SetDefault("redis.host", "localhost")
	v.SetDefault("redis.port", 6379)

	v.SetDefault("session.secret", "")
	v.SetDefault("session.expiry_minutes", 60)

	v.SetDefault("csrf.secret", "")
	v.SetDefault("log.level", "info")

	// A config file remains supported for local development, but is optional.
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")

	// Production Docker environment overrides config/default values.
	_ = v.BindEnv("app.name", "APP_NAME")
	_ = v.BindEnv("app.env", "APP_ENV")
	_ = v.BindEnv("app.port", "APP_PORT")
	_ = v.BindEnv("app.allowed_hosts", "APP_ALLOWED_HOSTS")

	_ = v.BindEnv("database.host", "DATABASE_HOST")
	_ = v.BindEnv("database.port", "DATABASE_PORT")
	_ = v.BindEnv("database.user", "DATABASE_USER")
	_ = v.BindEnv("database.password", "DATABASE_PASSWORD")
	_ = v.BindEnv("database.name", "DATABASE_NAME")
	_ = v.BindEnv("database.sslmode", "DATABASE_SSLMODE")

	_ = v.BindEnv("redis.host", "REDIS_HOST")
	_ = v.BindEnv("redis.port", "REDIS_PORT")

	_ = v.BindEnv("session.secret", "SESSION_SECRET")
	_ = v.BindEnv("session.expiry_minutes", "SESSION_EXPIRY_MINUTES")

	_ = v.BindEnv("csrf.secret", "CSRF_SECRET")
	_ = v.BindEnv("log.level", "LOG_LEVEL")

	if err := v.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) {
			return nil, err
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
