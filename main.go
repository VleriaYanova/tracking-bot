package main

import (
	"tracking-bot/db"
	"tracking-bot/handlers"
	"tracking-bot/repo"
	"tracking-bot/services"
)

func main() {
	db := db.NewGormDb()

	// courseRepo := repo.NewGormCoursesRepo(db)
	eventRepo := repo.NewEventRepo(db)

	// appServ := services.NewApartmentsService(courseRepo, http.DefaultClient)
	// subscribersServ := services.NewSubscriberService(subscribersRepo)
	eventServise := services.NewEventService(eventRepo)
	trackHandler := handlers.NewTrackingHandler(eventServise)

	// go trackHandler.StartBot()

	// time.Sleep(time.Second)
	trackHandler.StartTracking()
}
