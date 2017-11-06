package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)


type CityWeather struct {
  Key         string `json: Key`
  EnglishName string `json: EnglishName`
  WeatherText string `json: WeatherText`
  WeatherIcon int    `json: WeatherIcon`
}

func GetCitiesWeather(w http.ResponseWriter, req *http.Request) {
  //params := mux.Vars(req)

  API_KEY := "ng3J0uqcGnBR4NfraCjD9aVavQqnZQ4p"
  topcities := 50

  accuweather_url := fmt.Sprintf("http://dataservice.accuweather.com/currentconditions/v1/topcities/%d?apikey=%s", topcities, API_KEY)

  resp, err := http.Get(accuweather_url)

  if err != nil {
    log.Fatal("NewRequest: ", err)
    return
  }

  var records [50]CityWeather

  if err := json.NewDecoder(resp.Body).Decode(&records); err != nil {
  		log.Println(err)
  	}

  json.NewEncoder(w).Encode(records)
}


func main() {
    router := mux.NewRouter()
    router.HandleFunc("/weather", GetCitiesWeather).Methods("GET")
    log.Fatal(http.ListenAndServe(":12345", router))
}
