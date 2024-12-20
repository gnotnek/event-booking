package review

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

func (r *repo) Create(review *entity.Review) (*entity.Review, error) {
	if err := r.db.Create(review).Error; err != nil {
		return nil, err
	}

	return review, nil
}

func (r *repo) Save(review *entity.Review) (*entity.Review, error) {
	if err := r.db.Save(review).Error; err != nil {
		return nil, err
	}

	return review, nil
}

func (r *repo) FindAll() ([]entity.Review, error) {
	var reviews []entity.Review
	if err := r.db.Find(&reviews).Error; err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *repo) Find(id string) (*entity.Review, error) {
	var review entity.Review
	if err := r.db.Where("id = ?", id).First(&review).Error; err != nil {
		return nil, err
	}

	return &review, nil
}

func (r *repo) FindByEventID(eventID string) ([]entity.Review, error) {
	var reviews []entity.Review
	if err := r.db.Preload("Event").Where("event_id = ?", eventID).Find(&reviews).Error; err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *repo) FindByUserID(userID string) ([]entity.Review, error) {
	var reviews []entity.Review
	if err := r.db.Preload("User").Where("user_id = ?", userID).Find(&reviews).Error; err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *repo) Delete(id string) error {
	err := r.db.Where("id = ?", id).First(&entity.Review{}).Delete(&entity.Review{}).Error
	if err != nil {
		return err
	}

	return nil
}
