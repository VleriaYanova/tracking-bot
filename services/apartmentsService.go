package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"tracking-bot/models"
	"tracking-bot/repo"

	"gorm.io/gorm"
)

type ApartmentsService struct {
	repo       *repo.GormApartmentRepo
	httpClient *http.Client
}

func NewApartmentsService(r *repo.GormApartmentRepo, c *http.Client) *ApartmentsService {
	return &ApartmentsService{
		repo:       r,
		httpClient: c,
	}
}

func (s *ApartmentsService) GetApartments(link string) (*[]models.Apartment, error) {
	response, err := s.httpClient.Get(link)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("failed to get. status code: %d\n", response.StatusCode))
	}

	housing := &models.HousingResponse{}
	err = json.NewDecoder(response.Body).Decode(housing)
	if err != nil {
		return nil, err
	}

	return &housing.Housings.Items, nil
}

func (s *ApartmentsService) CreateIfNotExist(app *models.Apartment, eventType string) (bool, error) {
	foundApp, err := s.repo.Find(&models.Apartment{ID: app.ID}, eventType)
	found := foundApp.Name != ""

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	if found {
		return false, nil
	}
	_, err = s.repo.Create(app, eventType)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *ApartmentsService) RemoveDeletedApps(outerApps *[]models.Apartment, eventType string) (*[]models.Apartment, error) {
	innerApps, err := s.repo.GetAll(eventType)
	if err != nil {
		return nil, err
	}

	deletedApps := &[]models.Apartment{}
	for _, inApp := range *innerApps {
		found := false
		for _, outApp := range *outerApps {
			if inApp.ID == outApp.ID {
				found = true
				break
			}
		}
		if found {
			continue
		}
		err = s.repo.DeleteByID(inApp.ID, eventType)
		if err != nil {
			return nil, err
		}
		*deletedApps = append(*deletedApps, inApp)
	}

	return deletedApps, nil
}

func (s *ApartmentsService) GetById(id string, eventType string) (*models.Apartment, error) {
	return s.repo.Get(id, eventType)
}

func (s *ApartmentsService) Update(app *models.Apartment, id string, eventType string) (*models.Apartment, error) {
	return s.repo.Update(app, id, eventType)
}
