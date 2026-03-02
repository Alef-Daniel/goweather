package weather_client

import (
	"net/http"
	"time"
)

type Weather struct {
	client http.Client
}

func (w *Weather) GetForecastByLocation(location string) (string, error) {
	return "", nil
}

func (w *Weather) GetForecastByLocationAndDateRange(location string, from string, to string) (string, error) {
	return "", nil
}

func New() *Weather {
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	return &Weather{client: *client}
}
