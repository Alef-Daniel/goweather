package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Alef-Daniel/goweather/config"
	"github.com/Alef-Daniel/goweather/internal/api/handlers"
	"github.com/Alef-Daniel/goweather/internal/api/route"
	"github.com/Alef-Daniel/goweather/internal/application/usecases"
	"github.com/Alef-Daniel/goweather/internal/infrastructure/cache"
	"github.com/Alef-Daniel/goweather/internal/infrastructure/weather_client"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	weatherClient := weather_client.New(cfg.WeatherAPI.BaseURL, cfg.WeatherAPI.APIKey)
	cacheClient := cache.NewRedis(cfg.Cache.Endpoint, "")

	getWeatherUC := &usecases.GetWeatherUseCaseImpl{
		WeatherClient: weatherClient,
		Cache:         cacheClient,
	}

	getWeatherByRangeDateUC := &usecases.GetWeatherByRangeDateUseCaseImpl{
		WeatherClient: weatherClient,
		Cache:         cacheClient,
	}

	weatherHandler := handlers.NewWeatherHandler(getWeatherUC, getWeatherByRangeDateUC)

	rt := route.NewRouter(weatherHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), rt); err != nil {
		log.Fatal(err)
	}

}
