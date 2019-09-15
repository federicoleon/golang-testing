package services

import (
	"github.com/federicoleon/golang-testing/src/api/domain/locations"
	"github.com/federicoleon/golang-testing/src/api/utils/errors"
	"github.com/federicoleon/golang-testing/src/api/providers/locations_provider"
)

type locationsService struct{}

type locationsServiceInterface interface {
	GetCountry(countryId string) (*locations.Country, *errors.ApiError)
}

var (
	LocationsService locationsServiceInterface
)

func init() {
	LocationsService = &locationsService{}
}

func (s *locationsService) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return locations_provider.GetCountry(countryId)
}
