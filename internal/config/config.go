package config

import "fmt"

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Session  SessionConfig  `mapstructure:"session"`
	CSRF     CsrfConfig     `mapstructure:"csrf"`
	Log      LogConfig      `mapstructure:"log"`
}

type AppConfig struct {
	Name        string   `mapstructure:"name"`
	Env         string   `mapstructure:"env"`
	Port        string   `mapstructure:"port"`
	AllowedHost []string `mapstructure:"allowed_hosts"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"sslmode"`
}

type CsrfConfig struct {
	Secret string `mapstructure:"secret"`
}

func (d DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Name,
		d.SSLMode,
	)
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type SessionConfig struct {
	Secret        string `mapstructure:"secret"`
	ExpiryMinutes int    `mapstructure:"expiry_minutes"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
}
