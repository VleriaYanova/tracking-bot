package repo

import (
	"tracking-bot/models"

	"gorm.io/gorm"
)

type GormApartmentRepo struct {
	db *gorm.DB
}

func NewGormCoursesRepo(db *gorm.DB) *GormApartmentRepo {
	return &GormApartmentRepo{db: db}
}

func (r *GormApartmentRepo) Create(app *models.Apartment) (*models.Apartment, error) {
	err := r.db.Create(app).Error
	return app, err
}

func (r *GormApartmentRepo) Get(id string) (*models.Apartment, error) {
	app := &models.Apartment{}
	err := r.db.First(app, id).Error
	return app, err
}

func (r *GormApartmentRepo) GetAll() (*[]models.Apartment, error) {
	apps := &[]models.Apartment{}
	result := r.db.Limit(-1).Find(apps)
	return apps, result.Error
}

func (r *GormApartmentRepo) Find(app *models.Apartment) (*models.Apartment, error) {
	result := r.db.Where(app).First(app)
	return app, result.Error
}

func (r *GormApartmentRepo) DeleteByID(id string) error {
	app := &models.Apartment{}
	return r.db.Delete(app, id).Error
}

func (r *GormApartmentRepo) Update(app *models.Apartment, id string) (*models.Apartment, error) {
	err := r.db.Model(&models.Apartment{ID: id}).Updates(app).Error
	return app, err
}
