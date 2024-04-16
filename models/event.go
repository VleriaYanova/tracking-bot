package models

type Event struct {
	ID int `gorm:"primarykey"`
	// CreatedAt  time.Time
	Name       string
	Subscriber *[]Subscriber `gorm:"many2many:subscriber_events;"`
}

const (
	TwoYearsSell string = "twoYearsSell"
	InMomentSell string = "inMomentSell"
)
