package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Alef-Daniel/goweather/internal/api/dtos"
	"github.com/Alef-Daniel/goweather/internal/application/usecases"
)

type WeatherHandler struct {
	ucGetWeatherUseCase     usecases.GetWeatherUseCase
	ucGetWeatherByRangeDate usecases.GetWeatherByRangeDateUseCase
}

func NewWeatherHandler(ucGetWeatherUseCase usecases.GetWeatherUseCase, ucGetWeatherByRangeDate usecases.GetWeatherByRangeDateUseCase) *WeatherHandler {
	return &WeatherHandler{ucGetWeatherUseCase, ucGetWeatherByRangeDate}
}

func (wh *WeatherHandler) GetForecastWeather(w http.ResponseWriter, r *http.Request) {
	var req dtos.WeatherRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := wh.ucGetWeatherUseCase.Execute(r.Context(), req.Location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)

}

func (wh *WeatherHandler) GetForecastWeatherByRangeDate(w http.ResponseWriter, r *http.Request) {
	var req dtos.WeatherRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = req.Validate()
	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := wh.ucGetWeatherByRangeDate.Execute(r.Context(), req.Location, req.DateInit, req.DateEnd)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(resp)

}
