package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	APPName  string `mapstructure:"app_name"`
	Port     int    `mapstructure:"port"`
	LogLevel string `mapstructure:"log_level"`

	WeatherAPI WeatherAPIConfig `mapstructure:"weather_api"`
	Cache      CacheConfig      `mapstructure:"cache"`
}

type WeatherAPIConfig struct {
	BaseURL string `mapstructure:"base_url"`
	APIKey  string `mapstructure:"api_key"`
}

type CacheConfig struct {
	Endpoint string `mapstructure:"endpoint"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile("config/config.json")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil

}
