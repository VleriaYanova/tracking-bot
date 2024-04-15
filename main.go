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

	courseRepo := repo.NewGormCoursesRepo(db)
	chatRepo := repo.NewChatRepo(db)

	appServ := services.NewApartmentsService(courseRepo, http.DefaultClient)
	chatServ := services.NewChatService(chatRepo)

	trackHandler := handlers.NewTrackingHandler(appServ, chatServ)

	go trackHandler.StartBot()

	time.Sleep(time.Second)
	// trackHandler.StartTracking()
}
