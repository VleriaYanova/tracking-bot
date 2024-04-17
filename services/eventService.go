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

func (s *EventService) Create(event *models.Event) (*models.Event, error) {
	return s.repo.Create(event)
}

func (s *EventService) DeleteByChatID(id int) error { return s.repo.Delete(id) }

func (s *EventService) GetAllBySubscriber(subscriber *models.Event) (*[]models.Event, error) {
	return s.repo.GetAllBySubscriber(subscriber)
}
