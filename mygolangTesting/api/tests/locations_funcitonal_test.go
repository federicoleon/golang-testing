package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/golang-testing/mygolangTesting/api/utils/errors"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestGetCountriesnotFound(t *testing.T) {
	// Mock
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/BI",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusInternalServerError,
		RespBody:     `{"status":500,"message":"invalid error interface when getting country BR","error":""}`,
	})
	// Execute
	fmt.Println("Function test get countries")
	response, err := http.Get("http://localhost:8000/locations/countries/BR")

	// Validate
	assert.Nil(t, err)
	assert.NotNil(t, response)
	var apiErrors errors.APIerror
	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &apiErrors)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, response.StatusCode)
	assert.EqualValues(t, "invalid error interface when getting country BR", apiErrors.Message)
}

func TestGetCountriesNoError(t *testing.T) {
	// Mock
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/BR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":"BR","name":"Brasil","locale":"pt_BR","currency_id":"BRL","decimal_separator":",","thousands_separator":".","time_zone":"GMT-03:00","geo_information":{"location":{"latitude":-23.6821604,"longitude":-46.875494}},"states":[{"id":"BR-AC","name":"Acre"},{"id":"BR-AL","name":"Alagoas"},{"id":"BR-AP","name":"Amapá"},{"id":"BR-AM","name":"Amazonas"},{"id":"BR-BA","name":"Bahia"},{"id":"BR-CE","name":"Ceará"},{"id":"BR-DF","name":"Distrito Federal"},{"id":"BR-ES","name":"Espírito Santo"},{"id":"BR-GO","name":"Goiás"},{"id":"BR-MA","name":"Maranhão"},{"id":"BR-MT","name":"Mato Grosso"},{"id":"BR-MS","name":"Mato Grosso do Sul"},{"id":"BR-MG","name":"Minas Gerais"},{"id":"BR-PR","name":"Paraná"},{"id":"BR-PB","name":"Paraíba"},{"id":"BR-PA","name":"Pará"},{"id":"BR-PE","name":"Pernambuco"},{"id":"BR-PI","name":"Piauí"},{"id":"BR-RN","name":"Rio Grande do Norte"},{"id":"BR-RS","name":"Rio Grande do Sul"},{"id":"BR-RJ","name":"Rio de Janeiro"},{"id":"BR-RO","name":"Rondônia"},{"id":"BR-RR","name":"Roraima"},{"id":"BR-SC","name":"Santa Catarina"},{"id":"BR-SE","name":"Sergipe"},{"id":"BR-SP","name":"São Paulo"},{"id":"BR-TO","name":"Tocantins"}]}`,
	})
	// Execute
	fmt.Println("Function test get countries")
	response, err := http.Get("http://localhost:8000/locations/countries/BR")

	// Validate
	assert.Nil(t, err)
	assert.NotNil(t, response)
	var apiErrors errors.APIerror
	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &apiErrors)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, response.StatusCode)
	assert.EqualValues(t, "", apiErrors.Message)
}
