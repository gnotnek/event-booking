package config

import (
	"fmt"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Warn().Msg("No .env file found")
	}

	var c Config
	if err := env.Parse(&c); err != nil {
		log.Fatal().Msgf("unable to parse env: %s", err.Error())
	}

	return &c
}

type Config struct {
	App      App
	Database Database
}

type App struct {
	JwtSecretKey string `env:"JWT_SECRET_KEY"`
}

type Database struct {
	Host     string `env:"DATABASE_HOST"`
	Port     int    `env:"DATABASE_PORT"`
	User     string `env:"DATABASE_USER"`
	Password string `env:"DATABASE_PASSWORD"`
	Name     string `env:"DATABASE_NAME"`
}

func (d Database) DataSourceName() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		d.User, d.Password, d.Host, d.Port, d.Name)
}
