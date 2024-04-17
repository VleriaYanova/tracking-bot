package models

import "time"

type Subscriber struct {
	ID        int `gorm:"primarykey"`
	CreatedAt time.Time
	ChatID    int64
	Events    *[]Event `gorm:"many2many:subscriber_events;constraint:OnUpdate:CASCADE"`
}
