package helpers

import (
	"fmt"
	"io"
	"log"
	"naptalie-api/api/types"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ToURLValues converts ForecastRequest to url.Values
func ConvertToUrlValues(r *types.ForecastRequest) url.Values {
	values := url.Values{}

	// Location
	values.Set("latitude", strconv.FormatFloat(r.Location.Latitude, 'f', -1, 64))
	values.Set("longitude", strconv.FormatFloat(r.Location.Longitude, 'f', -1, 64))

	// Units
	if r.Units.Temperature != "" {
		values.Set("temperature_unit", r.Units.Temperature)
	}
	if r.Units.Precipitation != "" {
		values.Set("precipitation_unit", r.Units.Precipitation)
	}
	if r.Units.Windspeed != "" {
		values.Set("windspeed_unit", r.Units.Windspeed)
	}

	// Data parameters
	if len(r.Data.Daily) > 0 {
		values.Set("daily", strings.Join(r.Data.Daily, ","))
	}
	if len(r.Data.Hourly) > 0 {
		values.Set("hourly", strings.Join(r.Data.Hourly, ","))
	}
	if len(r.Data.Current) > 0 {
		values.Set("current", strings.Join(r.Data.Current, ","))
	}

	// Time parameters
	if r.Time.Timezone != "" {
		values.Set("timezone", r.Time.Timezone)
	}
	if r.Time.ForecastDays > 0 {
		values.Set("forecast_days", strconv.Itoa(r.Time.ForecastDays))
	}
	if r.Time.PastDays > 0 {
		values.Set("past_days", strconv.Itoa(r.Time.PastDays))
	}

	return values
}

// Helper functions for common locations
func IndianapolisLocation() types.Location {
	return types.Location{
		Latitude:  39.7684,
		Longitude: -86.1581,
	}
}

// Helper functions for common unit sets
func USUnits() types.Units {
	return types.Units{
		Temperature:   types.TempFahrenheit,
		Precipitation: types.PrecipInch,
		Windspeed:     types.WindMph,
	}
}

func MetricUnits() types.Units {
	return types.Units{
		Temperature:   types.TempCelsius,
		Precipitation: types.PrecipMM,
		Windspeed:     types.WindKmh,
	}
}

// Helper for common daily data
func BasicDailyData() types.DataParameters {
	return types.DataParameters{
		Daily: []string{
			types.DailyTempMax,
			types.DailyTempMin,
			types.DailyPrecipitationSum,
			types.DailyWeatherCode,
		},
	}
}

// Helper for common time settings
func WeekForecast() types.TimeParameters {
	return types.TimeParameters{
		Timezone:     "America/New_York",
		ForecastDays: 7,
	}
}

func BuildWeatherUrl(baseUrl string) string {
	fq := &types.ForecastRequest{
		Location: IndianapolisLocation(),
		Units:    USUnits(),
		Data:     BasicDailyData(),
	}
	urlValues := ConvertToUrlValues(fq)
	return baseUrl + urlValues.Encode()
}

func GetWeather(baseUrl string) {

	formattedUrl := BuildWeatherUrl(baseUrl)
	resp, err := http.Get(formattedUrl)
	if err != nil {
		log.Printf("error making req: %d %s", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("error reading ")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading response body: %v", err)
	}
	fmt.Printf("response from site! %s", string(body))
}
