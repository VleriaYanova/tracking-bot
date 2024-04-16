// package handlers

// import (
// 	"log"
// 	"tracking-bot/services"
// )

// func NewTwoYearsHandler(s *services.ApartmentsService, subscribersService *services.SubscriberService) *TrackingHandler {
// 	c, err := subscribersService.GetAll()
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	for _, chat := range *c {
// 		if _, ok := chats[chat.ChatID]; !ok {
// 			chats[chat.ChatID] = ""
// 		}
// 	}
// 	return &TrackingHandler{
// 		appService:         s,
// 		subscribersService: subscribersService,
// 		bot:                nil,
// 	}
// }
