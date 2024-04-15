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

var twoYearsLink = "https://fr.mos.ru/pokupka-nedvizhimosti-dlya-vseh/ajax.php?category[]=PARTICIPANTS&y2_sell=1&status[]=PROCESSING&status[]=FINISHED&price_min=0&price_max=100000000000000&price_m_min=0&price_m_max=100000000000&area_min=0&area_max=100000000&floor_min=-1&floor_max=100000000&open_sale=0&pagesize=100000"
var inMomentLink = "https://fr.mos.ru/pokupka-nedvizhimosti-dlya-vseh/ajax.php?category[]=PARTICIPANTS&for_sell=1&status[]=PROCESSING&status[]=FINISHED&price_min=8000000&price_max=22000000&price_m_min=184000&price_m_max=351000&area_min=39&area_max=86&floor_min=2&floor_max=24&open_sale=0&pagesize=100000"

func (s *ApartmentsService) GetApartments(command string) (*[]models.Apartment, error) {
	fmt.Println(command)
	switch command {
	case "/twoyears":
		fmt.Println("123")
	case "/in-moment":
		fmt.Println("lala")
	}
	response, err := s.httpClient.Get(twoYearsLink)
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

func (s *ApartmentsService) CreateIfNotExist(app *models.Apartment) (bool, error) {
	foundApp, err := s.repo.Find(&models.Apartment{ID: app.ID})
	found := foundApp.Name != ""

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	if found {
		return false, nil
	}
	_, err = s.repo.Create(app)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *ApartmentsService) RemoveDeletedApps(outerApps *[]models.Apartment) (*[]models.Apartment, error) {
	innerApps, err := s.repo.GetAll()
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
		err = s.repo.DeleteByID(inApp.ID)
		if err != nil {
			return nil, err
		}
		*deletedApps = append(*deletedApps, inApp)
	}

	return deletedApps, nil
}

func (s *ApartmentsService) GetById(id string) (*models.Apartment, error) {
	return s.repo.Get(id)
}

func (s *ApartmentsService) Update(app *models.Apartment, id string) (*models.Apartment, error) {
	return s.repo.Update(app, id)
}
