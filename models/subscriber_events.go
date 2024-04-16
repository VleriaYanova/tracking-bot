package models

type Subscribers_events struct {
	SubscriberId int `gorm:"primaryKey;column:subscriber_id" json:"subscriberId"`
	EventId      int `gorm:"primaryKey;column:event_id" json:"eventId"`
}
