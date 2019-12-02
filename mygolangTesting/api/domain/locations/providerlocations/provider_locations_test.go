package providerlocations

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("hello")
	rest.StartMockupServer()
	os.Exit(m.Run())
	rest.FlushMockups()
}

func TestGetCountryRestclientError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/BR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody: `{
			"message": "Country not found",
			"error": "not_found", 
			"status": 404, 
			"cause": []`,
	})
	//Init
	// Execute
	country, err := GetCountry("BR")
	//Validate
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when getting country BR", err.Message)
}

func TestGetCountryTimeout(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: 0,
	})
	//Init
	// Execute
	country, err := GetCountry("BR")
	//Validate
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when getting country BR", err.Message)
}

func TestGetCountryNotFound(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/IR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody: `{
			 "message": "Country not found",
			 "error": "not_found", 
			 "status": 404, 
			 "cause": []}`,
	})
	//Validate
	country, err := GetCountry("IR")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)

}

func TestGetCountryInvalid(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/IR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody: `{
			 "message": "Country not found",
			 "error": "not_found", 
			 "status": "404", 
			 "cause": []}`,
	})

	country, err := GetCountry("BR")
	//Validate
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when getting country BR", err.Message)
}

func TestGetCountryInvalidJson(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/BR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody: `{
			"id": 123,
			"name": "Brasil",
			"time_zone": "GMT-03:00"
		  }`,
	})

	country, err := GetCountry("BR")
	//Validate
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal when getting country data for BR", err.Message)

}

func TestGetCountryNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/BR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":"BR","name":"Brasil","locale":"pt_BR","currency_id":"BRL","decimal_separator":",","thousands_separator":".","time_zone":"GMT-03:00","geo_information":{"location":{"latitude":-23.6821604,"longitude":-46.875494}},"states":[{"id":"BR-AC","name":"Acre"},{"id":"BR-AL","name":"Alagoas"},{"id":"BR-AP","name":"Amapá"},{"id":"BR-AM","name":"Amazonas"},{"id":"BR-BA","name":"Bahia"},{"id":"BR-CE","name":"Ceará"},{"id":"BR-DF","name":"Distrito Federal"},{"id":"BR-ES","name":"Espírito Santo"},{"id":"BR-GO","name":"Goiás"},{"id":"BR-MA","name":"Maranhão"},{"id":"BR-MT","name":"Mato Grosso"},{"id":"BR-MS","name":"Mato Grosso do Sul"},{"id":"BR-MG","name":"Minas Gerais"},{"id":"BR-PR","name":"Paraná"},{"id":"BR-PB","name":"Paraíba"},{"id":"BR-PA","name":"Pará"},{"id":"BR-PE","name":"Pernambuco"},{"id":"BR-PI","name":"Piauí"},{"id":"BR-RN","name":"Rio Grande do Norte"},{"id":"BR-RS","name":"Rio Grande do Sul"},{"id":"BR-RJ","name":"Rio de Janeiro"},{"id":"BR-RO","name":"Rondônia"},{"id":"BR-RR","name":"Roraima"},{"id":"BR-SC","name":"Santa Catarina"},{"id":"BR-SE","name":"Sergipe"},{"id":"BR-SP","name":"São Paulo"},{"id":"BR-TO","name":"Tocantins"}]}`,
	})
	country, err := GetCountry("BR")
	//Validate
	assert.NotNil(t, country)
	assert.Nil(t, err)
	assert.EqualValues(t, "BR", country.ID)
	assert.EqualValues(t, "Brasil", country.Name)
	assert.EqualValues(t, "GMT-03:00", country.TimeZone)
	assert.EqualValues(t, 6, len(country.Name))
	fmt.Println(len(country.Name))

}
