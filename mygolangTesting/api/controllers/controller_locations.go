package controllers

import (
	"fmt"
	"github.com/golang-testing/mygolangTesting/api/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCountry takes a gin.Context and serves a response as Json
func GetCountry(c *gin.Context) {
	fmt.Println(c)
	fmt.Println("Inside Controller")
	country, err := services.LocationsService.GetCountry(c.Param("country_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, country)
}
