package types

// Location represents geographic coordinates
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Units represents measurement unit preferences
type Units struct {
	Temperature   string `json:"temperature_unit,omitempty"`
	Precipitation string `json:"precipitation_unit,omitempty"`
	Windspeed     string `json:"windspeed_unit,omitempty"`
}

// DataParameters represents what data to retrieve
type DataParameters struct {
	Daily   []string `json:"daily,omitempty"`
	Hourly  []string `json:"hourly,omitempty"`
	Current []string `json:"current,omitempty"`
}

// TimeParameters represents time-related settings
type TimeParameters struct {
	Timezone     string `json:"timezone,omitempty"`
	ForecastDays int    `json:"forecast_days,omitempty"`
	PastDays     int    `json:"past_days,omitempty"`
}

// ForecastRequest combines all parameter types
type ForecastRequest struct {
	Location Location       `json:"location"`
	Units    Units          `json:"units"`
	Data     DataParameters `json:"data"`
	Time     TimeParameters `json:"time"`
}

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Success bool   `json:"success"`
}

// WeatherData represents the inner weather API response structure
type WeatherData struct {
	Latitude         float64       `json:"latitude"`
	Longitude        float64       `json:"longitude"`
	GenerationTimeMs float64       `json:"generationtime_ms"`
	UTCOffsetSeconds int           `json:"utc_offset_seconds"`
	Timezone         string        `json:"timezone"`
	TimezoneAbbr     string        `json:"timezone_abbreviation"`
	Elevation        float64       `json:"elevation"`
	DailyUnits       DailyUnits    `json:"daily_units"`
	Daily            DailyForecast `json:"daily"`
}

type DailyUnits struct {
	Time          string `json:"time"`
	TempMax       string `json:"temperature_2m_max"`
	TempMin       string `json:"temperature_2m_min"`
	Precipitation string `json:"precipitation_sum"`
	WeatherCode   string `json:"weathercode"`
}

type DailyForecast struct {
	Time          []string  `json:"time"`
	TempMax       []float64 `json:"temperature_2m_max"`
	TempMin       []float64 `json:"temperature_2m_min"`
	Precipitation []float64 `json:"precipitation_sum"`
	WeatherCode   []int     `json:"weathercode"`
}

// Unit constants
const (
	TempCelsius    = "celsius"
	TempFahrenheit = "fahrenheit"

	PrecipMM   = "mm"
	PrecipInch = "inch"

	WindKmh = "kmh"
	WindMph = "mph"
	WindMs  = "ms"
	WindKn  = "kn"
)

// Daily parameter constants
const (
	DailyTempMax               = "temperature_2m_max"
	DailyTempMin               = "temperature_2m_min"
	DailyPrecipitationSum      = "precipitation_sum"
	DailyWeatherCode           = "weathercode"
	DailyWindSpeedMax          = "windspeed_10m_max"
	DailyWindDirectionDominant = "winddirection_10m_dominant"
	DailySunrise               = "sunrise"
	DailySunset                = "sunset"
)
