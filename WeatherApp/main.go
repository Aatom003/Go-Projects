package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Struct to hold API key data
type apiConfigData struct {
	OpenWeatherApiKey string `json:"OpenWeatherApiKey"`
}

// Configuration with your OpenWeather API key
var config = apiConfigData{
	OpenWeatherApiKey: "03b5533d80e05e933cf39deb345d6381", // Add your OpenWeather API key here
}

// Struct to represent the JSON response from OpenWeatherMap
type WeatherResponse struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure float64 `json:"pressure"`
		Humidity float64 `json:"humidity"`
	} `json:"main"`
}

// Middleware to log incoming requests
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}

// Handler for weather information
func weatherServer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Extracting the city from the form
		city := r.FormValue("city")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<p><strong>Name:</strong> %s</p>", city)

		// Creating the URL for the OpenWeatherMap API request
		url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, config.OpenWeatherApiKey)

		// Making a GET request to the OpenWeatherMap API
		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Reading the response body
		Body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read weather data", http.StatusInternalServerError)
			return
		}

		// Unmarshaling the JSON response into WeatherResponse struct
		var weatherResp WeatherResponse
		err = json.Unmarshal(Body, &weatherResp)
		if err != nil {
			http.Error(w, "Failed to parse weather data", http.StatusInternalServerError)
			return
		}

		// Displaying the temperature from the parsed response
		fmt.Fprintf(w, "<p><strong>Temperature:</strong> %.2f</p>", weatherResp.Main.Temp)
	}
}

func main() {
	// Create a file server to serve static files from the "./static" directory
	fileServer := http.FileServer(http.Dir("./static"))

	// Handle requests to the root path ("/") by serving files from the file server
	http.Handle("/", fileServer)

	// Handle requests to "/getWeather" by invoking weatherServer
	http.HandleFunc("/getWeather", weatherServer)

	fmt.Printf("Starting at port 8083.......")

	// Start the server with the logging middleware
	err := http.ListenAndServe(":8083", logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
