package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getWeather(w http.ResponseWriter, req *http.Request) {

	url := "https://api.open-meteo.com/v1/forecast?latitude=39.7684&longitude=-86.1581&daily=temperature_2m_max,temperature_2m_min,precipitation_sum,weathercode&timezone=America/New_York&temperature_unit=fahrenheit&precipitation_unit=inch&windspeed_unit=mph&forecast_days=1"
	resp, err := http.Get(url)
	if err != nil {
		log.Println("error making req: %d %s", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("error reading ")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading response body: %v", err)
	}
	fmt.Printf("response from site!", string(body))
}

func main() {
	http.HandleFunc("/", getWeather)

	fmt.Println("server running on port 8090")
	err := http.ListenAndServe(":8090", nil)

	if err != nil {
		fmt.Printf("error starting server %s\n", err)
	}
}
