package repo

import (
	"fmt"
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
	subscribers := &[]models.Subscriber{}
	result := r.db.Limit(-1).Find(subscribers)
	return subscribers, result.Error
}

func (r *GormSubscriberRepo) GetAllByEvent(event *models.Event) (*[]models.Subscriber, error) {
	subscribers := &[]models.Subscriber{}

	result := r.db.Preload("Events", "Events.ID = ?", event.ID).Find(subscribers)
	return subscribers, result.Error
}

func (r *GormSubscriberRepo) GetByChatID(chatID int64) (*models.Subscriber, error) {
	subscriber := &models.Subscriber{}

	result := r.db.Where(&models.Subscriber{ChatID: chatID}).Preload("Events").Find(subscriber)
	return subscriber, result.Error
}

func (r *GormSubscriberRepo) Create(subscriber *models.Subscriber) (*models.Subscriber, error) {
	err := r.db.Create(subscriber).Error
	return subscriber, err
}

func (r *GormSubscriberRepo) Update(subscriber *models.Subscriber, id int) (*models.Subscriber, error) {
	err := r.db.Model(subscriber).Updates(subscriber).Update("Events", subscriber.Events).Error
	return subscriber, err
}

func (r *GormSubscriberRepo) DeleteByChatID(id int64, eventType string) error {
	fmt.Println("deleted")
	return r.db.Where(&models.Subscriber{ChatID: id}).Preload("Events").Delete(&models.Subscriber{Events: &[]models.Event{{Name: eventType}}}).Error
}

func (r *GormSubscriberRepo) Unsubscribe(subID int, eventID int) error {
	return r.db.Table("subscriber_events").Where("event_id = ?", eventID).Where("subscriber_id = ?", subID).Delete("").Error
}
