package service

import (
	cities "github.com/kuzminprog/cities_information_service"
	"github.com/kuzminprog/cities_information_service/internal/repository"
)

type City interface {
	Create(city cities.CityRequest) (string, error)
	Delete(id int) error
	SetPopulation(id, population int) error
	GetFromRegion(region string) ([]string, error)
	GetFromDistrict(district string) ([]string, error)
	GetFromPopulation(population string) ([]string, error)
	GetFromFoundation(foundation string) ([]string, error)
	GetFull(id int) (*cities.City, error)
}

type Service struct {
	City
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		City: NewCityService(repos.CityList),
	}
}
