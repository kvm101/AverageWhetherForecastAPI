package awf

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type Param struct {
	Api  string `json:"api"`
	City string `json:"city"`
}

var rd = Redis{redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:9736",
	Password: "",
	DB:       0,
})}

var client = http.DefaultClient

func HandlerWCurrent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		req := r.Body
		defer req.Close()

		var param Param

		data, err := io.ReadAll(req)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(data, &param)

		if err != nil {
			log.Fatal(err)
		}

		url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", param.Api, param.City)

		resp, err := client.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		rd.setValue("current", string(body), time.Hour*12)
		w.WriteHeader(http.StatusOK)

		fmt.Fprintln(w, string(body))
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
