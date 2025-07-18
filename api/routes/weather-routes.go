package routes

import (
	"naptalie-api/api/helpers"
	"net/http"
)

func GetWeatherRoute(w http.ResponseWriter, req *http.Request) {
	baseUrl := "https://api.open-meteo.com/v1/forecast?"

	helpers.GetWeather(baseUrl)

}
