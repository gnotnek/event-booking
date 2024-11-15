package health

import (
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repo {
	return &repo{
		db: db,
	}
}

func (r *repo) CheckDatabase() error {
	return r.db.Exec("SELECT 1").Error
}
