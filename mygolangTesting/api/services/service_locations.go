package services

import (
	"fmt"
	"golang-testing/mygolangTestinggolangTesting/api/domain/locations"
	"golang-testing/mygolangTestinggolangTesting/api/domain/locations/providerlocations"
	"golang-testing/mygolangTestinggolangTesting/api/utils/errors"
)

// initialized = false set to true when the package is called
type locationsService struct{}

type locationServiceInterface interface {
	GetCountry(countryID string) (*locations.Country, *errors.APIerror)
}

// LocationsService is the locationServiceInterface
var LocationsService locationServiceInterface

func init() {
	fmt.Println("Init Service")
	LocationsService = &locationsService{}
}

// GetCountry takes a countryID
// and returns a pointer to locations.Country and a pointer to an APIerror
func (s *locationsService) GetCountry(countryID string) (*locations.Country, *errors.APIerror) {
	fmt.Println("Inside Service")
	return providerlocations.GetCountry(countryID)

}
