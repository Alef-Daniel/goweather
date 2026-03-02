package weather_client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/Alef-Daniel/goweather/internal/domain"
)

type Weather struct {
	client  http.Client
	baseURL string
	apiKey  string
}

func (w *Weather) GetForecastByLocation(ctx context.Context, location string) (string, error) {
	if location == "" {
		return "", domain.ErrInvalidLocation
	}

	buildURL, err := w.buildURL("VisualCrossingWebServices/rest/services/timeline/", location, "", "")
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, buildURL, nil)
	if err != nil {
		return "", err
	}

	resp, err := w.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("weather request failed: %w", err)
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("failed to read response: %w", err)
		}
		return string(body), nil
	case http.StatusBadRequest:
		return "", domain.ErrInvalidLocation

	case http.StatusUnauthorized:
		return "", domain.ErrUnauthorized

	case http.StatusNotFound:
		return "", domain.ErrLocationNotFound

	case http.StatusTooManyRequests:
		return "", domain.ErrRateLimited

	case http.StatusInternalServerError:
		return "", domain.ErrExternalService

	default:
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

}

func (w *Weather) GetForecastByLocationAndDateRange(location string, from string, to string) (string, error) {
	return "", nil
}

func (w *Weather) buildURL(pathComplete, location, date_first, date_end string) (string, error) {
	url, err := url.Parse(w.baseURL)
	if err != nil {
		return "", err
	}
	url.Path = path.Join(url.Path, pathComplete, location)
	if date_first != "" && date_end != "" {
		url.Path = path.Join(url.Path, date_first, date_end)
	}

	q := url.Query()
	q.Set("key", w.apiKey)

	url.RawQuery = q.Encode()
	return url.String(), nil

}

func New(baseURL, APIkey string) *Weather {
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	return &Weather{client: *client, baseURL: baseURL, apiKey: APIkey}
}
