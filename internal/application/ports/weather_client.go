package ports

import (
	"context"
	"time"

	"github.com/Alef-Daniel/goweather/internal/api/dtos"
)

type WeatherClient interface {
	GetForecastByLocation(ctx context.Context, location string) (*dtos.WeatherResponsAPI, error)
	GetForecastByLocationAndDateRange(ctx context.Context, location string, dateInit, dateEnd *time.Time) (*dtos.WeatherResponsAPI, error)
}
