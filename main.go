package main

import (
	"net/http"
	"os"
	"time"
	"tracking-bot/db"
	"tracking-bot/handlers"
	"tracking-bot/repo"
	"tracking-bot/services"
)

func main() {
	if os.Args[1] != "-bot-key" && os.Args[2] != "" {
		panic("set '-bot-key {key}' as first app argument")
	}

	db := db.NewGormDb()

	appRepo := repo.NewGormApartmentRepo(db)
	eventRepo := repo.NewEventRepo(db)
	subRepo := repo.NewSubscriberRepo(db)

	appServ := services.NewApartmentsService(appRepo, http.DefaultClient)
	subscribersServ := services.NewSubscriberService(subRepo)
	eventServise := services.NewEventService(eventRepo)

	trackHandler := handlers.NewTrackingHandler(eventServise, subscribersServ, appServ)

	go trackHandler.StartBot(os.Args[2])

	time.Sleep(time.Second * 2)
	trackHandler.StartTracking()
}
