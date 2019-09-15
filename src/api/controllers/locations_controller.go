package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/federicoleon/golang-testing/src/api/services"
	"net/http"
)

func GetCountry(c *gin.Context) {
	country, err := services.LocationsService.GetCountry(c.Param("country_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, country)
}
