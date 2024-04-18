package db

import (
	"fmt"
	"tracking-bot/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewGormDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("houses.db"), &gorm.Config{})

	if err != nil {

	}

	err = db.AutoMigrate(&models.Subscriber{}, &models.Event{})
	db.Table("twoyears").AutoMigrate(&models.Apartment{})
	db.Table("momentsell").AutoMigrate(&models.Apartment{})
	if err != nil {
		panic(fmt.Sprintf("failed to migrate: %s", err.Error()))
	}
	return db
}
