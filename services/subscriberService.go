package services

import (
	"tracking-bot/models"
	"tracking-bot/repo"
)

type SubscriberService struct {
	repo *repo.GormSubscriberRepo
}

func NewSubscriberService(repo *repo.GormSubscriberRepo) *SubscriberService {
	return &SubscriberService{
		repo: repo,
	}
}

func (s *SubscriberService) GetAll() (*[]models.Subscriber, error) {
	return s.repo.GetAll()
}

func (s *SubscriberService) GetAllByEvent(event *models.Event) (*[]models.Subscriber, error) {
	return s.repo.GetAllByEvent(event)
}

func (s *SubscriberService) GetByChatID(chatID int64) (*models.Subscriber, error) {
	return s.repo.GetByChatID(chatID)
}

func (s *SubscriberService) Create(subscriber *models.Subscriber) (*models.Subscriber, error) {
	return s.repo.Create(subscriber)
}

func (s *SubscriberService) Update(subscriber *models.Subscriber, id int) (*models.Subscriber, error) {
	return s.repo.Update(subscriber, id)
}

func (s *SubscriberService) DeleteByChatID(id int64) error { return s.repo.DeleteByChatID(id) }
