package models

import "time"

type Chat struct {
	ID        int `gorm:"primarykey"`
	CreatedAt time.Time
	ChatID    int64
}
