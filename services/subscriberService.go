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

func (s *SubscriberService) Create(Subscriber *models.Subscriber) (*models.Subscriber, error) {
	return s.repo.Create(Subscriber)
}

func (s *SubscriberService) DeleteByChatID(id int64) error { return s.repo.DeleteByChatID(id) }

func (s *SubscriberService) GetAllByEvent(event string) (*[]models.Subscriber, error) {
	return s.repo.GetAllByEvent(event)
}
