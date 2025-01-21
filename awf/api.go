package awf

import (
	"net/http"

	"github.com/redis/go-redis/v9"
)

var red = Redis{redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:9736",
	Password: "",
	DB:       0,
})}

func HandlerWCurrent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		w.WriteHeader(http.StatusOK)
	}
}

func HandlerWForecast(w http.ResponseWriter, r *http.Request) {

}

func HandlerWForecastT(w http.ResponseWriter, r *http.Request) {

}

func HandlerWAlert(w http.ResponseWriter, r *http.Request) {

}

func HandlerWHistory(w http.ResponseWriter, r *http.Request) {

}
