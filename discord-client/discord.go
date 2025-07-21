package discordclient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"naptalie-api/api/helpers"
	"naptalie-api/api/types"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

func Ready(sesh *discordgo.Session, event *discordgo.Ready) {
	log.Printf("uwu logged in as: %v#%v", sesh.State.User.Username, sesh.State.User.Discriminator)
}

// SendWeatherForecast sends a formatted weather forecast to Discord using your types
func SendWeatherForecast(s *discordgo.Session, channelID string, response *types.Response) error {
	// Parse the weather data from the response
	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return fmt.Errorf("failed to marshal response data: %v", err)
	}

	var weatherData types.WeatherData
	if err := json.Unmarshal(dataBytes, &weatherData); err != nil {
		return fmt.Errorf("failed to parse weather data: %v", err)
	}

	// Create embed fields for each day
	var fields []*discordgo.MessageEmbedField
	maxDays := len(weatherData.Daily.Time)
	if maxDays > 7 {
		maxDays = 7 // Limit to 7 days for Discord embed
	}

	for i := 0; i < maxDays; i++ {
		if i >= len(weatherData.Daily.TempMax) ||
			i >= len(weatherData.Daily.TempMin) ||
			i >= len(weatherData.Daily.WeatherCode) {
			break
		}

		// Parse date
		date, err := time.Parse("2006-01-02", weatherData.Daily.Time[i])
		if err != nil {
			continue
		}

		dayName := date.Format("Mon, Jan 2")
		emoji, condition := helpers.WeatherCodeToEmoji(weatherData.Daily.WeatherCode[i])

		// Format temperature based on units
		tempUnit := weatherData.DailyUnits.TempMax
		precipUnit := weatherData.DailyUnits.Precipitation

		value := fmt.Sprintf("%s %s\nHigh: %.1f%s | Low: %.1f%s",
			emoji, condition,
			weatherData.Daily.TempMax[i], tempUnit,
			weatherData.Daily.TempMin[i], tempUnit)

		// Add precipitation if significant
		if weatherData.Daily.Precipitation[i] > 0.001 {
			value += fmt.Sprintf("\nRain: %.3f%s", weatherData.Daily.Precipitation[i], precipUnit)
		}

		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   dayName,
			Value:  value,
			Inline: true,
		})
	}

	// Create location string
	locationStr := fmt.Sprintf("📍 %.3f°N, %.3f°W",
		weatherData.Latitude, weatherData.Longitude)
	if weatherData.Elevation > 0 {
		locationStr += fmt.Sprintf(" (%.0fm elevation)", weatherData.Elevation)
	}

	// Create the embed
	embed := &discordgo.MessageEmbed{
		Title:       "🌤️ Weather Forecast",
		Description: locationStr,
		Color:       0x87CEEB, // Sky blue
		Fields:      fields,
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Timezone: %s | Updated", weatherData.Timezone),
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	// Send using ChannelMessageSendComplex with MessageSend options
	_, err = s.ChannelMessageSendComplex(channelID, &discordgo.MessageSend{
		Content: response.Message,
		Embeds:  []*discordgo.MessageEmbed{embed},
		// Additional Discord send options:
		// TTS: false,
		// Files: []*discordgo.File{},
		// AllowedMentions: &discordgo.MessageAllowedMentions{
		//     Parse: []discordgo.AllowedMentionType{},
		// },
		// Reference: &discordgo.MessageReference{
		//     MessageID: "some-message-id",
		// },
		// Components: []discordgo.MessageComponent{},
	})

	return err
}

func MessageCreate(sesh *discordgo.Session, msg *discordgo.MessageCreate) {

	baseUrl := "http://localhost:8090/weather"

	if msg.Author.ID == sesh.State.User.ID {
		return
	}

	if msg.Content == "!ping" {
		sesh.ChannelMessageSend(msg.ChannelID, "Pong!")
	}
	if msg.Content == "!weather" {
		resp, err := http.Get(baseUrl)

		if err != nil {
			sesh.ChannelMessageSend(msg.ChannelID, "Error fetching weather!")
		}

		defer resp.Body.Close()
		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return
		}

		// Decode the JSON into your types.Response struct
		var weatherResponse types.Response
		if err := json.Unmarshal(body, &weatherResponse); err != nil {
			return
		}
		SendWeatherForecast(sesh, msg.ChannelID, &weatherResponse)
	}
}
