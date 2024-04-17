package main

import (
	"net/http"
	"time"
	"tracking-bot/db"
	"tracking-bot/handlers"
	"tracking-bot/repo"
	"tracking-bot/services"
)

func main() {
	db := db.NewGormDb()

	appRepo := repo.NewGormApartmentRepo(db)
	eventRepo := repo.NewEventRepo(db)
	subRepo := repo.NewSubscriberRepo(db)

	appServ := services.NewApartmentsService(appRepo, http.DefaultClient)
	subscribersServ := services.NewSubscriberService(subRepo)
	eventServise := services.NewEventService(eventRepo)

	trackHandler := handlers.NewTrackingHandler(eventServise, subscribersServ, appServ)

	go trackHandler.StartBot()

	time.Sleep(time.Second * 2)
	trackHandler.StartTracking()
}
