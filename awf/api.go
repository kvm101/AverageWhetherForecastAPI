package awf

import (
	"net/http"
)

func HandlerWCurrent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getValue("")

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
