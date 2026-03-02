package domain

type Weather struct {
	Location      Location
	Current       *CurrentConditions
	DailyForecast []DailyForecast
}

type Location struct {
	Address   string
	Latitude  float64
	Longitude float64
	Timezone  string
	TzOffset  float64
}

type CurrentConditions struct {
	DateTime    string
	Temperature float64
	Humidity    float64
	WindSpeed   float64
	Pressure    float64
	Conditions  string
}

type DailyForecast struct {
	Date       string
	TempMax    float64
	TempMin    float64
	TempAvg    float64
	Humidity   float64
	Precip     float64
	WindSpeed  float64
	Conditions string
}
