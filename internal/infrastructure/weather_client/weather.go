package weather_client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/Alef-Daniel/goweather/internal/api/dtos"
	"github.com/Alef-Daniel/goweather/internal/domain"
)

type Weather struct {
	client  http.Client
	baseURL string
	apiKey  string
}

func (w *Weather) GetForecastByLocation(ctx context.Context, location string) (*dtos.WeatherResponsAPI, error) {
	buildURL, err := w.buildURL("VisualCrossingWebServices/rest/services/timeline/", location, nil, nil)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, buildURL, nil)

	if err != nil {
		return nil, err
	}

	resp, err := w.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("weather request failed: %w", err)
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var apiResp dtos.WeatherResponsAPI
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response: %w", err)
		}
		err = json.Unmarshal(body, &apiResp)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal response: %w", err)
		}
		return &apiResp, nil
	case http.StatusBadRequest:
		return nil, domain.ErrInvalidLocation

	case http.StatusUnauthorized:
		return nil, domain.ErrUnauthorized

	case http.StatusNotFound:
		return nil, domain.ErrLocationNotFound

	case http.StatusTooManyRequests:
		return nil, domain.ErrRateLimited

	case http.StatusInternalServerError:
		return nil, domain.ErrExternalService

	default:
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

}

func (w *Weather) GetForecastByLocationAndDateRange(ctx context.Context, location string, dateInit, dateEnd *time.Time) (*dtos.WeatherResponsAPI, error) {
	buildURL, err := w.buildURL("VisualCrossingWebServices/rest/services/timeline/", location, dateInit, dateEnd)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, buildURL, nil)
	fmt.Println(req.URL.String())
	if err != nil {
		return nil, err
	}

	resp, err := w.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("weather request failed: %w", err)
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var apiResp dtos.WeatherResponsAPI
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response: %w", err)
		}
		err = json.Unmarshal(body, &apiResp)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal response: %w", err)
		}
		return &apiResp, nil
	case http.StatusBadRequest:
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
		return nil, domain.ErrInvalidLocation

	case http.StatusUnauthorized:
		return nil, domain.ErrUnauthorized

	case http.StatusNotFound:
		return nil, domain.ErrLocationNotFound

	case http.StatusTooManyRequests:
		return nil, domain.ErrRateLimited

	case http.StatusInternalServerError:
		return nil, domain.ErrExternalService

	default:
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

}

func (w *Weather) buildURL(pathComplete, location string, date_first, date_end *time.Time) (string, error) {
	var dateInitStr, dateEndStr string
	url, err := url.Parse(w.baseURL)
	if err != nil {
		return "", err
	}
	url.Path = path.Join(url.Path, pathComplete, location)
	if date_first != nil && date_end != nil {

		dateInitStr = date_first.Format("2006-01-02")
		dateEndStr = date_end.Format("2006-01-02")
		url.Path = path.Join(url.Path, dateEndStr, dateInitStr)
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
