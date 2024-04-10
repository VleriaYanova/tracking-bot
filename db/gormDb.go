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

	err = db.AutoMigrate(&models.Apartment{}, &models.Chat{})
	if err != nil {
		panic(fmt.Sprintf("failed to migrate: %s", err.Error()))
	}
	return db
}
