package main

import (
	"awf/auth"
	"awf/awf"
	"awf/user"
	"log"
	"net/http"
)

func main() {
	wheaterAPI := http.NewServeMux()
	userAPI := http.NewServeMux()
	authAPI := http.NewServeMux()

	wheaterAPI.HandleFunc("/weather/current", awf.HandlerWCurrent)
	wheaterAPI.HandleFunc("/weather/forecast", awf.HandlerWForecast)
	wheaterAPI.HandleFunc("/weather/forecast-10", awf.HandlerWForecastT)
	wheaterAPI.HandleFunc("/weather/alerts", awf.HandlerWAlert)
	wheaterAPI.HandleFunc("/weather/history", awf.HandlerWHistory)
	userAPI.HandleFunc("/user/preferences", user.HandlerUPreferences)
	userAPI.HandleFunc("/user/history", user.HandlerUHistory)
	authAPI.HandleFunc("/auth/login", auth.HandlerAuthLogin)
	authAPI.HandleFunc("/auth/registration", auth.HandlerAuthRegistration)

	log.Println("Server is running. api:port - (wheaterAPI:9898, userAPI:9797, authAPI:9696).")

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
