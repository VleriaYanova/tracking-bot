package repo

import (
	"tracking-bot/models"

	"gorm.io/gorm"
)

type GormChatRepo struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) *GormChatRepo {
	return &GormChatRepo{db: db}
}

func (r *GormChatRepo) GetAll() (*[]models.Chat, error) {
	chats := &[]models.Chat{}
	result := r.db.Limit(-1).Find(chats)
	return chats, result.Error
}

func (r *GormChatRepo) Create(chat *models.Chat) (*models.Chat, error) {
	err := r.db.Create(chat).Error
	return chat, err
}

func (r *GormChatRepo) DeleteByChatID(id int64) error {
	return r.db.Where(&models.Chat{ChatID: id}).Delete(&models.Chat{}).Error
}
