package handlers

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"os/signal"
// 	"time"
// 	"tracking-bot/models"
// 	"tracking-bot/services"

// 	"github.com/go-telegram/bot"
// botModels "github.com/go-telegram/bot/models"
// )

// var eventsSubscribers map[models.Event][]models.Subscriber = map[models.Event][]models.Subscriber{
// 	models.InMomentSell: []models.Subscriber{},
// 	models.TwoYearsSell: []models.Subscriber{},
// }

// type TrackingHandler struct {
// 	subscribersService *services.SubscriberService
// 	appService         *services.ApartmentsService
// 	bot                *bot.Bot
// }

// func (h *TrackingHandler) StartTracking(update *botModels.Update) {
// 	if h.bot == nil {
// 		panic("start bot before tracking")
// 	}
// 	track := true
// 	log.Println("START TRACKING")
// 	for track {

// 	}
// }
