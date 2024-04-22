package repo

import (
	"tracking-bot/models"

	"gorm.io/gorm"
)

type GormApartmentRepo struct {
	db *gorm.DB
}

func NewGormApartmentRepo(db *gorm.DB) *GormApartmentRepo {
	return &GormApartmentRepo{db: db}
}

func (r *GormApartmentRepo) Create(app *models.Apartment, eventType string) (*models.Apartment, error) {
	event := eventType[1:]
	err := r.db.Table(event).Create(app).Error
	return app, err
}

func (r *GormApartmentRepo) Get(id string, eventType string) (*models.Apartment, error) {
	event := eventType[1:]
	app := &models.Apartment{}
	err := r.db.Table(event).First(app, id).Error
	return app, err
}

func (r *GormApartmentRepo) GetAll(eventType string) (*[]models.Apartment, error) {
	event := eventType[1:]
	apps := &[]models.Apartment{}
	result := r.db.Limit(-1).Table(event).Find(apps)
	return apps, result.Error
}

func (r *GormApartmentRepo) Find(app *models.Apartment, eventType string) (*models.Apartment, error) {
	event := eventType[1:]
	result := r.db.Table(event).Where(app).First(app)
	return app, result.Error
}

func (r *GormApartmentRepo) DeleteByID(id string, eventType string) error {
	event := eventType[1:]
	app := &models.Apartment{}
	return r.db.Table(event).Where("id=?", id).Delete(app, id).Error
}

func (r *GormApartmentRepo) Update(app *models.Apartment, id string, eventType string) (*models.Apartment, error) {
	event := eventType[1:]
	err := r.db.Table(event).Where("id=?", id).Updates(app).Error
	return app, err
}
