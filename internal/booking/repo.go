package booking

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

func (r *repo) Create(booking *entity.Booking) (*entity.Booking, error) {
	if err := r.db.Create(booking).Error; err != nil {
		return nil, err
	}

	return booking, nil
}

func (r *repo) Save(booking *entity.Booking) (*entity.Booking, error) {
	if err := r.db.Save(booking).Error; err != nil {
		return nil, err
	}

	return booking, nil
}

func (r *repo) FindAll() ([]entity.Booking, error) {
	var bookings []entity.Booking
	if err := r.db.Model(&entity.Booking{}).Preload("Event").Preload("User").Find(&bookings).Error; err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *repo) FindByUserID(userID string) ([]entity.Booking, error) {
	var bookings []entity.Booking
	if err := r.db.Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *repo) FindByEventID(eventID string) ([]entity.Booking, error) {
	var bookings []entity.Booking
	if err := r.db.Model(&entity.Booking{}).Preload("Event").Preload("User").Where("event_id = ?", eventID).Find(&bookings).Error; err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *repo) Find(id string) (*entity.Booking, error) {
	booking := new(entity.Booking)
	if err := r.db.Where("id = ?", id).Preload("Event").Preload("User").First(booking).Error; err != nil {
		return nil, err
	}

	return booking, nil
}

func (r *repo) Delete(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&entity.Booking{}).Error; err != nil {
		return err
	}

	return nil
}
