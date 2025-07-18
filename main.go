package main

import (
	"fmt"
	"naptalie-api/api/routes"
	"net/http"
)

func main() {

	fmt.Println("server running on port 8090")
	http.HandleFunc("/weather", routes.GetWeatherRoute)
	err := http.ListenAndServe(":8090", nil)

	if err != nil {
		fmt.Printf("error starting server %s\n", err)
	}
}
