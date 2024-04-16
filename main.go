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
	// b, _ := subscribersRepo.Create(&models.Subscriber{ChatID: 324234, Events: &[]models.Event{{ID: 2}}})
	b := subscribersRepo.GetAllByEvent(&models.Event{ID: 1})
	for _, subs := range *b {
		fmt.Println(subs)
	}

	// subscribersServ.Create(&models.Subscriber{Events: fmt.Sprintf("%s;%s", models.TwoYearsSell, models.InMomentSell)})

	// trackHandler := handlers.NewTrackingHandler(appServ, chatServ)

	// go trackHandler.StartBot()

	// time.Sleep(time.Second)
	// trackHandler.StartTracking()
}
