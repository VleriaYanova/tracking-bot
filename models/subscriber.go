package models

import "time"

type Subscriber struct {
	ID        int `gorm:"primarykey"`
	CreatedAt time.Time
	ChatID    int64
	Events    string
}
