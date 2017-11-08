package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strings"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
)


type AccuCity struct {
  Key         string `json: Key`
  EnglishName string `json: EnglishName`
  WeatherText string `json: WeatherText`
  WeatherIcon int    `json: WeatherIcon`
}

type CityView struct {
  City    string
  Weather string
  Icon    string
}

func GetCitiesWeather(w http.ResponseWriter, req *http.Request) {

  keys, ok := req.URL.Query()["search"]
  searchTerm := ""

  if ok && len(keys) > 0 {
    searchTerm = keys[0]
  }

  // should be extracted in a func
  API_KEY := "OvWIPZTdAD6ltfebkTv3K5Mh4w0UXCIV"
  topcities := 50

  AccuCity_url := fmt.Sprintf("http://dataservice.accuweather.com/currentconditions/v1/topcities/%d?apikey=%s", topcities, API_KEY)

  resp, err := http.Get(AccuCity_url)

  if err != nil {
    log.Fatal("NewRequest: ", err)
    return
  }

  var allCities []AccuCity

  if err := json.NewDecoder(resp.Body).Decode(&allCities); err != nil {
  		log.Println(err)
  }

  matchingCities := findMatchingWeatherCities(allCities, searchTerm)
  cityViews := createCityViews(matchingCities)

  json.NewEncoder(w).Encode(cityViews)

}


func findMatchingWeatherCities(allCities []AccuCity, text string) []AccuCity {
    if len(text) > 0 {
      var matchingCities []AccuCity

      for _, city := range allCities {
        if (strings.EqualFold(city.WeatherText, text)) {
          matchingCities = append(matchingCities, city)
        }
      }

      return matchingCities
    }

    return allCities
}

func createCityViews(accuCities []AccuCity) []CityView {
  var cityViews []CityView
  for _, accuCity := range accuCities {
    cityView := CityView{
          City: accuCity.EnglishName,
          Weather: accuCity.WeatherText,
          Icon: fmt.Sprintf("https://developer.accuweather.com/sites/default/files/%02d-s.png", accuCity.WeatherIcon)}
    cityViews = append(cityViews, cityView)
  }

  return cityViews
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/weather", GetCitiesWeather).Methods("GET")
    handler := cors.Default().Handler(router)
    log.Fatal(http.ListenAndServe(":12345", handler))
}


/*
AccWeather output example:

[
{
    "Key": "28143",
    "LocalizedName": "Dhaka",
    "EnglishName": "Dhaka",
    "Country": {
      "ID": "BD",
      "LocalizedName": "Bangladesh",
      "EnglishName": "Bangladesh"
    },
    "TimeZone": {
      "Code": "BDT",
      "Name": "Asia/Dhaka",
      "GmtOffset": 6,
      "IsDaylightSaving": false,
      "NextOffsetChange": null
    },
    "GeoPosition": {
      "Latitude": 23.7098,
      "Longitude": 90.40711,
      "Elevation": {
        "Metric": {
          "Value": 5,
          "Unit": "m",
          "UnitType": 5
        },
        "Imperial": {
          "Value": 16,
          "Unit": "ft",
          "UnitType": 0
        }
      }
    },
    "LocalObservationDateTime": "2017-11-07T22:10:00+06:00",
    "EpochTime": 1510071000,
    "WeatherText": "Clear",
    "WeatherIcon": 33,
    "IsDayTime": false,
    "Temperature": {
      "Metric": {
        "Value": 24.2,
        "Unit": "C",
        "UnitType": 17
      },
      "Imperial": {
        "Value": 76,
        "Unit": "F",
        "UnitType": 18
      }
    },
    "MobileLink": "http://m.AccuCity.com/en/bd/dhaka/28143/current-weather/28143?lang=en-us",
    "Link": "http://www.AccuCity.com/en/bd/dhaka/28143/current-weather/28143?lang=en-us"
  }
]
*/
