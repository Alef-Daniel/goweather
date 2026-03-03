package route

import (
	"github.com/Alef-Daniel/goweather/internal/api/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"net/http"
)

func NewRouter(weatherHandler *handlers.WeatherHandler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(15 * 1000 * 1000 * 1000)) // 15s

	r.Route("/api", func(r chi.Router) {
		r.Post("/weather", weatherHandler.GetForecastWeather)
		r.Post("/weather/date", weatherHandler.GetForecastWeatherByRangeDate)
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	return r
}
