package services

import (
	"tracking-bot/models"
	"tracking-bot/repo"
)

type EventService struct {
	repo *repo.GormEventRepo
}

func NewEventService(repo *repo.GormEventRepo) *EventService {
	return &EventService{
		repo: repo,
	}
}

func (s *EventService) GetAll() (*[]models.Event, error) {
	return s.repo.GetAll()
}

func (s *EventService) GetAllBySubscriber(subscriber *models.Event) (*[]models.Event, error) {
	return s.repo.GetAllBySubscriber(subscriber)
}

func (s *EventService) GetByName(name string) (*models.Event, error) {
	return s.repo.GetByName(name)
}

func (s *EventService) Create(event *models.Event) (*models.Event, error) {
	return s.repo.Create(event)
}

func (s *EventService) DeleteByChatID(id int64, eventType string) error {
	return s.repo.Delete(id, eventType)
}
