package repo

import (
	"tracking-bot/models"

	"gorm.io/gorm"
)

type GormEventRepo struct {
	db *gorm.DB
}

func NewEventRepo(db *gorm.DB) *GormEventRepo {
	return &GormEventRepo{db: db}
}

func (r *GormEventRepo) GetAll() (*[]models.Event, error) {
	Event := &[]models.Event{}
	result := r.db.Limit(-1).Preload("Subscriber").Find(Event)
	return Event, result.Error
}

func (r *GormEventRepo) GetAllBySubscriber(subscriber *models.Event) (*[]models.Event, error) {
	Event := &[]models.Event{}

	result := r.db.Preload("Subscriber", "Subscriber.ID = ?", subscriber.ID).Find(Event)
	return Event, result.Error
}

func (r *GormEventRepo) GetByName(name string) (*models.Event, error) {
	Event := &models.Event{}

	result := r.db.Where(&models.Event{Name: name}).Find(Event)
	return Event, result.Error
}

func (r *GormEventRepo) Create(event *models.Event) (*models.Event, error) {
	err := r.db.Create(event).Error
	return event, err
}

func (r *GormEventRepo) Delete(id int) error {
	return r.db.Where(&models.Event{ID: id}).Delete(&models.Event{}).Error
}
