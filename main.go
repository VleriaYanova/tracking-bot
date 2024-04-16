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

	// a, _ := subscribersServ.GetAllByEvent(models.InMomentSell)
	// fmt.Println(a)

	b, _ := subscribersRepo.Create(&models.Subscriber{ChatID: 54654456, Events: &[]models.Event{{ID: 2}, {ID: 1}}})

	fmt.Println(b.Events)

	// subscribersServ.Create(&models.Subscriber{Events: fmt.Sprintf("%s;%s", models.TwoYearsSell, models.InMomentSell)})

	// trackHandler := handlers.NewTrackingHandler(appServ, chatServ)

	// go trackHandler.StartBot()

	// time.Sleep(time.Second)
	// trackHandler.StartTracking()
}
