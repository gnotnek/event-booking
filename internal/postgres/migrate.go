package postgres

import (
	"event-booking/internal/entity"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&entity.User{}, &entity.Event{}, &entity.Booking{}, &entity.HealthComponent{}, &entity.Review{})
	if err != nil {
		log.Fatal().Err(err).Msg("could not migrate database")
	}
}
