package postgres

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate()
	if err != nil {
		log.Fatal().Err(err).Msg("could not migrate database")
	}
}
