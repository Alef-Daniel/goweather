package usecases

import (
	"context"
	"encoding/json"
	"fmt"

	"time"

	"github.com/Alef-Daniel/goweather/internal/api/dtos"
	"github.com/Alef-Daniel/goweather/internal/application/ports"
	"github.com/Alef-Daniel/goweather/internal/domain"
)

type GetWeatherByRangeDateUseCase interface {
	Execute(ctx context.Context, location string, dateInit, DateEnd *time.Time) (*domain.Weather, error)
}

type GetWeatherByRangeDateUseCaseImpl struct {
	WeatherClient ports.WeatherClient
	Cache         ports.Cache
}

func (g *GetWeatherByRangeDateUseCaseImpl) Execute(ctx context.Context, location string, dateInit, dateEnd *time.Time) (*domain.Weather, error) {
	if location == "" {
		return nil, domain.ErrInvalidLocation
	}

	var cacheKey string
	if dateInit != nil && dateEnd != nil {
		cacheKey = fmt.Sprintf("%s:%s:%s", location, dateInit.Format("2006-01-02"), dateEnd.Format("2006-01-02"))
	} else {
		cacheKey = location
	}

	var apiResp *dtos.WeatherResponsAPI

	if cached, err := g.Cache.Get(ctx, cacheKey); err == nil && cached != "" {
		if err := json.Unmarshal([]byte(cached), &apiResp); err != nil {
			apiResp = nil
		}
	}

	if apiResp == nil {
		var err error
		apiResp, err = g.WeatherClient.GetForecastByLocationAndDateRange(ctx, location, dateInit, dateEnd)
		if err != nil {
			return nil, err
		}

		data, _ := json.Marshal(apiResp)
		_ = g.Cache.Set(ctx, cacheKey, string(data), 12*time.Hour)
	}

	weather := mapWeatherDTOToDomain(apiResp)

	return weather, nil
}
