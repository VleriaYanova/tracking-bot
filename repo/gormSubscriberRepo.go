package repo

import (
	"tracking-bot/models"

	"gorm.io/gorm"
)

type GormSubscriberRepo struct {
	db *gorm.DB
}

func NewSubscriberRepo(db *gorm.DB) *GormSubscriberRepo {
	return &GormSubscriberRepo{db: db}
}

func (r *GormSubscriberRepo) GetAll() (*[]models.Subscriber, error) {
	Subscribers := &[]models.Subscriber{}
	result := r.db.Limit(-1).Find(Subscribers)
	return Subscribers, result.Error
}

func (r *GormSubscriberRepo) GetAllByEvent(event string) (*[]models.Subscriber, error) {
	Subscribers := &[]models.Subscriber{}
	result := r.db.Limit(-1).Where("events LIKE ?", "%"+event+"%").Find(Subscribers)
	return Subscribers, result.Error
}

func (r *GormSubscriberRepo) Create(Subscriber *models.Subscriber) (*models.Subscriber, error) {
	err := r.db.Create(Subscriber).Error
	return Subscriber, err
}

func (r *GormSubscriberRepo) DeleteByChatID(id int64) error {
	return r.db.Where(&models.Subscriber{ChatID: id}).Delete(&models.Subscriber{}).Error
}
