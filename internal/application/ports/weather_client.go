package ports

type WeatherClient interface {
	GetForecastByLocation(location string) (string, error)
	GetForecastByLocationAndDateRange(location string, from string, to string) (string, error)
}
