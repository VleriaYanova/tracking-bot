package handlers

import (
	"fmt"
	"log"
	"tracking-bot/models"
	"tracking-bot/services"

	"github.com/go-telegram/bot"
)

type TrackingHandler struct {
	subscribersService *services.SubscriberService
	eventService       *services.EventService
	bot                *bot.Bot
}

func NewTrackingHandler(e *services.EventService) *TrackingHandler {
	return &TrackingHandler{
		bot:          nil,
		eventService: e,
	}
}

var eventsSubscribers map[string][]models.Subscriber = map[string][]models.Subscriber{}

func (h *TrackingHandler) InitEventSubscriber(eventSubscriberMap map[string][]models.Subscriber) {
	allEvents, err := h.eventService.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, evnt := range *allEvents {
		eventsSubscribers[evnt.Name] = *evnt.Subscriber
	}
}

func (h *TrackingHandler) StartTracking() {
	// if h.bot == nil {
	// 	panic("start bot before tracking")
	// }
	track := true
	h.InitEventSubscriber(eventsSubscribers)
	log.Println("START TRACKING")
	for track {
		log.Println("Processing...")
		allEvents, err := h.eventService.GetAll()

		for _, event := range *allEvents {
			if len(eventsSubscribers[event.Name]) > 0 {
				fmt.Println("start event handler")
			}
		}

		if err != nil {
			fmt.Println(err)
		}
		track = false
	}
}
