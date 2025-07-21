package routes

import (
	"encoding/json"
	"naptalie-api/api/helpers"
	"naptalie-api/api/types"
	"net/http"
)

func HandleDiscordWebhookWeather(w http.ResponseWriter, r *http.Request) {
	// Set headers for JSON response
	w.Header().Set("Content-Type", "application/json")

	baseUrl := "https://api.open-meteo.com/v1/forecast?"

	getWeatherData := helpers.GetWeather(baseUrl)
	// Your logic here
	response := types.Response{
		Message: "Command executed successfully",
		Data:    getWeatherData.Data,
		Success: true,
	}

	// Encode and send JSON response
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
