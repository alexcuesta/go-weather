## Go Weather

This is a very basic REST service written in Go which pulls information from AccuWeather

### Installation

go get -u github.com/gorilla/mux
go get github.com/rs/cors
go build

### Run

chmod 744 go-weather
./go-weather

### Usage

From PostMan try:

GET localhost:12345/weather
