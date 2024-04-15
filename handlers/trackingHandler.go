package handlers

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
	"tracking-bot/models"
	"tracking-bot/services"

	"github.com/go-telegram/bot"
	botModels "github.com/go-telegram/bot/models"
)

var chats map[int64]string = map[int64]string{}

type NotifyType string

const (
	AppAdded                  NotifyType = "add"
	AppRemoved                NotifyType = "removed"
	AppStatusAuction          NotifyType = "statusAuction"
	AppStatusFirstDeclaration NotifyType = "statusFirstDeclaration"
)

type TrackingHandler struct {
	chatService *services.ChatService
	appService  *services.ApartmentsService
	bot         *bot.Bot
}

func NewTrackingHandler(s *services.ApartmentsService, chatService *services.ChatService) *TrackingHandler {
	c, err := chatService.GetAll()
	if err != nil {
		log.Println(err.Error())
	}
	for _, chat := range *c {
		if _, ok := chats[chat.ChatID]; !ok {
			chats[chat.ChatID] = ""
		}
	}
	return &TrackingHandler{
		appService:  s,
		chatService: chatService,
		bot:         nil,
	}
}

func (h *TrackingHandler) StartTracking(update *botModels.Update) {
	if h.bot == nil {
		panic("start bot before tracking")
	}
	track := true
	log.Println("START TRACKING")
	for track {
		log.Println("Processing...")
		// Get all appartments from site
		outerApps, err := h.appService.GetApartments(update.Message.Text)
		if err != nil {
			fmt.Println(err.Error())
		}

		// Saves appartment in database if it not exists
		// and if appartment was created - notify all bot subscribers
		for _, outApp := range *outerApps {
			created, err := h.appService.CreateIfNotExist(&outApp)
			if err != nil {
				fmt.Println(err.Error())
			}
			if created {
				h.NotifyAllChats(&outApp, AppAdded)
				continue
			}
			inApp, err := h.appService.GetById(outApp.ID)
			if err != nil {
				fmt.Println(err.Error())
			}
			if inApp.Requested == outApp.Requested {
				continue
			}
			_, err = h.appService.Update(&outApp, inApp.ID)
			if err != nil {
				fmt.Println(err.Error())
			}
			if outApp.Requested == 1 {
				h.NotifyAllChats(&outApp, AppStatusFirstDeclaration)
			} else if inApp.Requested < 2 && outApp.Requested > 1 {
				h.NotifyAllChats(&outApp, AppStatusAuction)
			}
		}

		// Remove appartment from database if it was deleted from site
		// and notify all bot subscribers
		removedApps, err := h.appService.RemoveDeletedApps(outerApps)
		if err != nil {
			fmt.Println(err.Error())
		}
		for _, rmApp := range *removedApps {
			h.NotifyAllChats(&rmApp, AppRemoved)
		}

		log.Println("Sleep for 10 minutes...")
		time.Sleep(time.Minute * 10)
	}
}

func (h *TrackingHandler) NotifyAllChats(app *models.Apartment, notifyType NotifyType) {
	if app == nil || app.ID == "" {
		fmt.Println("Got malformed app in notify method")
		return
	}
	fmt.Printf("Notifying all chats. Notify type: %s\n", notifyType)
	text := ""
	switch notifyType {
	case AppAdded:
		text += "Добавлена новая квартира в список: "
	case AppRemoved:
		text += "Квартира удалена из списка: "
	case AppStatusAuction:
		text += "Назначен аукцион на квартиру: "
	case AppStatusFirstDeclaration:
		text += "Подано первое заявление на квартиру: "
	}

	link := fmt.Sprintf("https://fr.mos.ru/uchastnikam-programmy/karta-renovatsii/%s/?ft=1&object=%s&object_type=TWO_YEARS_SELL&flat_id=%s", app.Object_code, app.Object_id, app.ID)
	text += "\n\n" + link

	media := &botModels.InputMediaPhoto{
		Media:   fmt.Sprintf("https://fr.mos.ru" + app.Plan),
		Caption: text,
	}
	for chatID := range chats {
		h.bot.SendMediaGroup(context.Background(), &bot.SendMediaGroupParams{
			ChatID: chatID,
			Media: []botModels.InputMedia{
				media,
			},
		})
	}
}

func (h *TrackingHandler) StartBot() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(h.ChatHook()),
	}

	b, err := bot.New("6939638145:AAEUhk_3pqdDwLWPfpUC6Jdm_YAUu6eOskQ", opts...)
	if err != nil {
		panic("bot start failed: " + err.Error())
	}
	h.bot = b

	ok, err := b.SetMyCommands(ctx, &bot.SetMyCommandsParams{
		Commands: []botModels.BotCommand{{
			Command:     "/start",
			Description: "Запустить бота",
		}, {
			Command:     "/stop",
			Description: "Остановить бота",
		}},
		Scope:        nil,
		LanguageCode: "",
	})
	if !ok {
		panic("failed to add bot menu: " + err.Error())
	}

	b.Start(ctx)
}

func (h *TrackingHandler) ChatHook() func(ctx context.Context, b *bot.Bot, update *botModels.Update) {
	return func(ctx context.Context, b *bot.Bot, update *botModels.Update) {
		if _, ok := chats[update.Message.Chat.ID]; !ok && update.Message.Text == "/start" {
			_, err := h.chatService.Create(&models.Chat{
				ChatID: update.Message.Chat.ID,
			})
			if err != nil {
				log.Println("failed to add chat to db: " + err.Error())
				return
			}
			chats[update.Message.Chat.ID] = ""
			h.StartTracking(update)
		} else if ok && update.Message.Text == "/stop" {
			delete(chats, update.Message.Chat.ID)
			err := h.chatService.DeleteByChatID(update.Message.Chat.ID)
			if err != nil {
				log.Println("failed to delete chat from db: " + err.Error())
				return
			}
		}
	}
}
