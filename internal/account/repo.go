package account

import (
	"event-booking/internal/entity"

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

func (r *repo) CreateAccount(user *entity.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *repo) FindByEmail(email string) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repo) DeleteByEmail(email string) error {
	if err := r.db.Where("email = ?", email).Delete(&entity.User{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *repo) FindByID(id string) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
