package app

import "golang-testing/mygolangTestinggolangTesting/api/controllers"

func mapURLs() {
	router.GET("/locations/countries/:country_id", controllers.GetCountry)
}
