package main

import (
	"fmt"
	"tracking-bot/db"
	"tracking-bot/models"
	"tracking-bot/repo"
)

func main() {
	db := db.NewGormDb()

	// courseRepo := repo.NewGormCoursesRepo(db)
	subscribersRepo := repo.NewSubscriberRepo(db)

	// appServ := services.NewApartmentsService(courseRepo, http.DefaultClient)
	// subscribersServ := services.NewSubscriberService(subscribersRepo)

	b := subscribersRepo.GetAllByEvent(&models.Event{ID: 3})

	fmt.Println(b)

	// subscribersServ.Create(&models.Subscriber{Events: fmt.Sprintf("%s;%s", models.TwoYearsSell, models.InMomentSell)})

	// trackHandler := handlers.NewTrackingHandler(appServ, chatServ)

	// go trackHandler.StartBot()

	// time.Sleep(time.Second)
	// trackHandler.StartTracking()
}
