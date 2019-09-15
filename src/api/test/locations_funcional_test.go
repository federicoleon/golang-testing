package test

import (
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/federicoleon/golang-testing/src/api/utils/errors"
	"encoding/json"
)

func TestGetCountriesNotFound(t *testing.T) {
	// Init:
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"status": 404, "error": "not_found", "message": "no country with id AR"}`,
	})

	// Execution:
	response, err := http.Get("http://localhost:8080/locations/countries/AR")

	// Validation:
	assert.Nil(t, err)
	assert.NotNil(t, response)
	bytes, _ := ioutil.ReadAll(response.Body)

	var apiErr errors.ApiError
	err = json.Unmarshal(bytes, &apiErr)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "not_found", apiErr.Error)
	assert.EqualValues(t, "no country with id AR", apiErr.Message)
}

func TestGetCountriesNoError(t *testing.T) {
	//TODO: Implement me!
}
