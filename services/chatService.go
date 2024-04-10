package services

import (
	"tracking-bot/models"
	"tracking-bot/repo"
)

type ChatService struct {
	repo *repo.GormChatRepo
}

func NewChatService(repo *repo.GormChatRepo) *ChatService {
	return &ChatService{
		repo: repo,
	}
}

func (s *ChatService) GetAll() (*[]models.Chat, error) {
	return s.repo.GetAll()
}

func (s *ChatService) Create(chat *models.Chat) (*models.Chat, error) {
	return s.repo.Create(chat)
}

func (s *ChatService) DeleteByChatID(id int64) error { return s.repo.DeleteByChatID(id) }
