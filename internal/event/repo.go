package event

import (
	"event-booking/internal/entity"

	"github.com/google/uuid"
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

func (r *repo) Create(event *entity.Event) (*entity.Event, error) {
	if err := r.db.Create(event).Error; err != nil {
		return nil, err
	}

	return event, nil
}

func (r *repo) Save(event *entity.Event) (*entity.Event, error) {
	if err := r.db.Save(event).Error; err != nil {
		return nil, err
	}

	return event, nil
}

func (r *repo) FindAll() ([]entity.Event, error) {
	var events []entity.Event
	if err := r.db.Find(&events).Error; err != nil {
		return nil, err
	}

	return events, nil
}

func (r *repo) Find(id uuid.UUID) (*entity.Event, error) {
	var event entity.Event
	if err := r.db.Where("id = ?", id).First(&event).Error; err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *repo) Delete(id uuid.UUID) error {
	if err := r.db.Where("id = ?", id).Delete(&entity.Event{}).Error; err != nil {
		return err
	}

	return nil
}
