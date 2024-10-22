package config

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/rs/zerolog/log"
)

func Load() *Config {
	var c Config
	if err := env.Parse(&c); err != nil {
		log.Fatal().Err(err).Msg("could not load configuration")
	}

	return &c
}

type Config struct {
	App      App
	Database Database
}

type App struct {
	SecretKey string `env:"APP_SECRET"`
	JwtKey    string `env:"JWT_SECRET_KEY"`
}

type Database struct {
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
}

func (d Database) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.User, d.Password, d.Name)
}
