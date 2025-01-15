package main

import (
	"awf/awf"
	"log"
	"net/http"
	"time"
)

func startInfoUpdate() {
	// client := http.DefaultClient

	for {
		// client.Get("https://api.weatherapi.com/v1/current.json?key=h7454705f74g44b429e231834251970&q=London&aqi=no")
		time.Sleep(time.Hour * 24)
	}
}

func main() {
	wheaterAPI := http.NewServeMux()
	userAPI := http.NewServeMux()
	authAPI := http.NewServeMux()

	wheaterAPI.HandleFunc("/weather/current", awf.HandlerWCurrent)
	wheaterAPI.HandleFunc("/weather/forecast", awf.HandlerWForecast)
	wheaterAPI.HandleFunc("/weather/forecast-10", awf.HandlerWForecastT)
	wheaterAPI.HandleFunc("/weather/alerts", awf.HandlerWAlert)
	wheaterAPI.HandleFunc("/weather/history", awf.HandlerWHistory)
	userAPI.HandleFunc("/user/preferences", awf.HandlerUPreferences)
	userAPI.HandleFunc("/user/history", awf.HandlerUHistory)
	authAPI.HandleFunc("/auth/login", awf.HandlerAuthLogin)
	authAPI.HandleFunc("/auth/registration", awf.HandlerAuthRegistration)

	log.Println("Server is running. api:port - (wheaterAPI:9898, userAPI:9797, authAPI:9696).")
	go startInfoUpdate()

	if err := http.ListenAndServe(":9898", wheaterAPI); err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":9797", userAPI); err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":9696", authAPI); err != nil {
		log.Fatal(err)
	}
}
