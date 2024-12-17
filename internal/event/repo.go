package event

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

func (r *repo) Find(id string) (*entity.Event, error) {
	var event entity.Event
	if err := r.db.Where("id = ?", id).First(&event).Error; err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *repo) FindByName(name string) (*entity.Event, error) {
	var event entity.Event
	if err := r.db.Where("name = ?", name).First(&event).Error; err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *repo) Delete(id string) error {
	err := r.db.Where("id = ?", id).First(&entity.Event{}).Delete(&entity.Event{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) FilterByCriteria(criteria map[string]interface{}) ([]entity.Event, error) {
	var events []entity.Event
	if err := r.db.Where(criteria).Find(&events).Error; err != nil {
		return nil, err
	}

	return events, nil
}
