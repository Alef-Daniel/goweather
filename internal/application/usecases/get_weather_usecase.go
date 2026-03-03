package usecases

import (
	"context"
	"encoding/json"

	"time"

	"github.com/Alef-Daniel/goweather/internal/api/dtos"
	"github.com/Alef-Daniel/goweather/internal/application/ports"
	"github.com/Alef-Daniel/goweather/internal/domain"
)

type GetWeatherUseCase interface {
	Execute(ctx context.Context, location string) (*domain.Weather, error)
}

type GetWeatherUseCaseImpl struct {
	WeatherClient ports.WeatherClient
	Cache         ports.Cache
}

func (g *GetWeatherUseCaseImpl) Execute(ctx context.Context, location string) (*domain.Weather, error) {
	if location == "" {
		return nil, domain.ErrInvalidLocation
	}

	var apiResp *dtos.WeatherResponsAPI

	if cached, err := g.Cache.Get(ctx, location); err == nil && cached != "" {
		if err := json.Unmarshal([]byte(cached), &apiResp); err != nil {

			apiResp = nil
		}
	}

	if apiResp == nil {
		var err error
		apiResp, err = g.WeatherClient.GetForecastByLocation(ctx, location)
		if err != nil {
			return nil, err
		}

		data, _ := json.Marshal(apiResp)
		_ = g.Cache.Set(ctx, location, string(data), 12*time.Hour)
	}

	weather := mapWeatherDTOToDomain(apiResp)

	return weather, nil
}

func mapWeatherDTOToDomain(dto *dtos.WeatherResponsAPI) *domain.Weather {
	w := &domain.Weather{
		Location: domain.Location{
			Address:   dto.ResolvedAddress,
			Latitude:  dto.Latitude,
			Longitude: dto.Longitude,
			Timezone:  dto.Timezone,
			TzOffset:  dto.Tzoffset,
		},
		Current: &domain.CurrentConditions{
			DateTime:    dto.Days[0].Datetime,
			Temperature: dto.Days[0].Temp,
			Humidity:    dto.Days[0].Humidity,
			WindSpeed:   dto.Days[0].Windspeed,
			Pressure:    dto.Days[0].Pressure,
			Conditions:  dto.Days[0].Conditions,
		},
		DailyForecast: make([]domain.DailyForecast, len(dto.Days)),
	}

	for i, d := range dto.Days {
		w.DailyForecast[i] = domain.DailyForecast{
			Date:       d.Datetime,
			TempMax:    d.Tempmax,
			TempMin:    d.Tempmin,
			TempAvg:    d.Temp,
			Humidity:   d.Humidity,
			Precip:     d.Precip,
			WindSpeed:  d.Windspeed,
			Conditions: d.Conditions,
		}
	}

	return w
}
