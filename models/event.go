package models

import "time"

type Event struct {
	ID         int `gorm:"primarykey"`
	CreatedAt  time.Time
	Name       string
	Subscriber *[]Subscriber `gorm:"many2many:subscriber_events;joinForeignKey:subscriber_id;joinReferences:event_id";"`
}

const (
	TwoYearsSell string = "twoYearsSell"
	InMomentSell string = "inMomentSell"
)
