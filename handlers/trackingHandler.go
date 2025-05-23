package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
	"tracking-bot/models"
	"tracking-bot/services"

	"github.com/go-telegram/bot"
	botModels "github.com/go-telegram/bot/models"
	"gorm.io/gorm"
)

type TrackingHandler struct {
	subscribersService *services.SubscriberService
	eventService       *services.EventService
	appService         *services.ApartmentsService
	bot                *bot.Bot
}

type NotifyType string

const (
	AppAdded                  NotifyType = "add"
	AppRemoved                NotifyType = "removed"
	AppStatusAuction          NotifyType = "statusAuction"
	AppStatusFirstDeclaration NotifyType = "statusFirstDeclaration"
)

func NewTrackingHandler(eventService *services.EventService, subService *services.SubscriberService, appService *services.ApartmentsService) *TrackingHandler {
	return &TrackingHandler{
		bot:                nil,
		eventService:       eventService,
		appService:         appService,
		subscribersService: subService,
	}
}

var eventsSubscribers = map[string][]models.Subscriber{
	models.TwoYearsSell: {},
	models.InMomentSell: {},
}

func (h *TrackingHandler) InitEventSubscriber() {
	allEvents, err := h.eventService.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, evnt := range *allEvents {
		eventsSubscribers[evnt.Name] = *evnt.Subscriber
	}
}

func (h *TrackingHandler) StartTracking() {
	if h.bot == nil {
		panic("start bot before tracking")
	}
	track := true
	h.InitEventSubscriber()
	log.Println("START TRACKING")
	for track {
		log.Println("Run main loop...")
		allEvents, err := h.eventService.GetAll()

		for _, event := range *allEvents {
			if len(eventsSubscribers[event.Name]) > 0 {
				go h.Handle(event.Name)
			}
		}

		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Minute * 5)
	}
}

var Link string
var prefLink string

func (h *TrackingHandler) Handle(eventName string) {
	switch eventName {
	case models.TwoYearsSell:
		Link = twoYearsSellLink
	case models.InMomentSell:
		Link = inMomentSellLink
	}

	h.ApartmentsHandler(Link, eventName)
}

func (h *TrackingHandler) readableEventType(eventType string) string {
	switch eventType {
	case models.TwoYearsSell:
		return "ПОКУПКА В ТЕЧЕНИИ 2-Х ЛЕТ"
	case models.InMomentSell:
		return "ДОКУПКА В МОМЕНТ ПЕРЕЕЗДА"
	}
	return ""
}

func (h *TrackingHandler) NotifyAllSubscribers(app *models.Apartment, notifyType NotifyType, eventType string) {
	if app == nil || app.ID == "" {
		fmt.Println("Got malformed app in notify method")
		return
	}
	fmt.Printf("Notifying all chats. Notify type: %s\n", notifyType)
	text := h.readableEventType(eventType) + "\n\n"
	switch notifyType {
	case AppAdded:
		text += "✅ Добавлена новая квартира в список: "
	case AppRemoved:
		text += "❌ Квартира удалена из списка: "
	case AppStatusAuction:
		text += "⚖️ Назначен аукцион на квартиру: "
	case AppStatusFirstDeclaration:
		text += "📝 Подано первое заявление на квартиру: "
	}
	if app.Y2_sell == 1 {
		prefLink = fmt.Sprintf("https://fr.mos.ru/uchastnikam-programmy/karta-renovatsii/%s/?ft=1&object=%s&object_type=TWO_YEARS_SELL&flat_id=%s", app.Object_code, app.Object_id, app.ID)
	} else {
		prefLink = fmt.Sprintf("https://fr.mos.ru/uchastnikam-programmy/karta-renovatsii/%s/?ft=1&object=%s&object_type=AVAILABLE_FOR_SELL&&flat_id=%s", app.Object_code, app.Object_id, app.ID)
	}
	text += "\n\n" + prefLink

	media := &botModels.InputMediaPhoto{
		Media:   fmt.Sprintf("https://fr.mos.ru" + app.Plan),
		Caption: text,
	}
	for _, sub := range eventsSubscribers[eventType] {
		h.bot.SendMediaGroup(context.Background(), &bot.SendMediaGroupParams{
			ChatID: sub.ChatID,
			Media: []botModels.InputMedia{
				media,
			},
		})
	}
}

func (h *TrackingHandler) SubscriberHook() func(ctx context.Context, b *bot.Bot, update *botModels.Update) {
	return func(ctx context.Context, b *bot.Bot, update *botModels.Update) {
		if update.Message == nil {
			return
		}
		eventName := update.Message.Text
		if eventName == models.ActiveEvents {
			h.SendActiveSubscriberEvents(update.Message.Chat.ID)
			return
		}
		if _, ok := eventsSubscribers[eventName]; !ok {
			return
		}
		event, err := h.eventService.GetByName(eventName)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("failed get event by name: " + err.Error())
		}
		if errors.Is(err, gorm.ErrRecordNotFound) || event.ID == 0 {
			event, err = h.eventService.Create(&models.Event{Name: eventName})
			if err != nil {
				panic("cannot create event: " + err.Error())
			}
		}

		sub, err := h.subscribersService.GetByChatID(update.Message.Chat.ID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			panic("failed get sub by chat ID: " + err.Error())
		}

		// User first time subcribes on any of events
		if errors.Is(err, gorm.ErrRecordNotFound) || sub.ID == 0 {
			sub, err = h.subscribersService.Create(&models.Subscriber{ChatID: update.Message.Chat.ID, Events: &[]models.Event{*event}})
			if err != nil {
				panic("cannot create subscriber: " + err.Error())
			}
			eventsSubscribers[eventName] = append(eventsSubscribers[eventName], *sub)
			return
		}

		if len(*sub.Events) > 0 {
			found := false
			for _, event := range *sub.Events {
				if event.Name == eventName {
					err = h.subscribersService.Unsubscribe(sub.ID, event.ID)
					if err != nil {
						panic("cannot update subscriber: " + err.Error())
					}
					found = true
				}
			}
			if found {
				for i, foundSub := range eventsSubscribers[eventName] {
					if foundSub.ChatID == sub.ChatID {
						eventsSubscribers[eventName] = append(eventsSubscribers[eventName][:i], eventsSubscribers[eventName][i+1:]...)
					}
				}
				return
			}
		}
		*sub.Events = append(*sub.Events, *event)
		_, err = h.subscribersService.Update(sub, sub.ID)
		if err != nil {
			panic("cannot update subscriber: " + err.Error())
		}
		eventsSubscribers[eventName] = append(eventsSubscribers[eventName], *sub)
	}
}

func (h *TrackingHandler) SendActiveSubscriberEvents(chatID int64) {
	subscriber, err := h.subscribersService.GetByChatID(chatID)
	if err != nil {
		fmt.Println("Ошибка при получении подписчика:", err)
		return
	}

	var events string
	if subscriber == nil || len(*subscriber.Events) == 0 {
		events = "Вы ни на что не подписаны"
	} else {
		events = "Вы подписаны на: \n"
		for _, s := range *subscriber.Events {
			events += "\n" + h.readableEventType(s.Name)
		}
	}

	h.bot.SendMessage(context.Background(), &bot.SendMessageParams{
		ChatID: subscriber.ChatID,
		Text:   events,
	})
}

func (h *TrackingHandler) GetEventTypeByApp(app *models.Apartment) string {
	if app.Y2_sell == 1 {
		return models.TwoYearsSell
	} else {
		return models.InMomentSell
	}
}

func (h *TrackingHandler) StartBot(botKey string) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(h.SubscriberHook()),
	}

	b, err := bot.New(botKey, opts...)
	if err != nil {
		panic("bot start failed: " + err.Error())
	}
	h.bot = b

	ok, err := b.SetMyCommands(ctx, &bot.SetMyCommandsParams{
		Commands: []botModels.BotCommand{{
			Command:     models.TwoYearsSell,
			Description: "Подписаться/Отписаться на покупку квартир в течении 2х лет",
		}, {
			Command:     models.InMomentSell,
			Description: "Подписаться/Отписаться на покупку квартир в момент переезда",
		}, {
			Command:     models.ActiveEvents,
			Description: "Активные подписки",
		}},
		Scope:        nil,
		LanguageCode: "",
	})
	if !ok {
		panic("failed to add bot menu: " + err.Error())
	}

	b.Start(ctx)
}
