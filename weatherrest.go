package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strings"
    "github.com/gorilla/mux"
)


type CityWeather struct {
  Key         string `json: Key`
  EnglishName string `json: EnglishName`
  WeatherText string `json: WeatherText`
  WeatherIcon int    `json: WeatherIcon`
}

func GetCitiesWeather(w http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)

  // should be extracted in a func
  API_KEY := "ng3J0uqcGnBR4NfraCjD9aVavQqnZQ4p"
  topcities := 50

  accuweather_url := fmt.Sprintf("http://dataservice.accuweather.com/currentconditions/v1/topcities/%d?apikey=%s", topcities, API_KEY)

  resp, err := http.Get(accuweather_url)

  if err != nil {
    log.Fatal("NewRequest: ", err)
    return
  }

  var allCities []CityWeather

  if err := json.NewDecoder(resp.Body).Decode(&allCities); err != nil {
  		log.Println(err)
  }

  matchingCities := findMatchingWeatherCities(allCities, params["searchText"])


  json.NewEncoder(w).Encode(matchingCities)
}


func findMatchingWeatherCities(allCities []CityWeather, text string) []CityWeather {
    if len(text) > 0 {
      var matchingCities []CityWeather

      for _, city := range allCities {
        if (strings.EqualFold(city.WeatherText, text)) {
          matchingCities = append(matchingCities, city)
        }
      }

      return matchingCities
    }

    return allCities
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/weather/{searchText}", GetCitiesWeather).Methods("GET")
    log.Fatal(http.ListenAndServe(":12345", router))
}
