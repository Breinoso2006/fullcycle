package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	WeatherAPIKey string
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	_ = viper.ReadInConfig()

	weatherKey := os.Getenv("weatherKey")
	if weatherKey == "" {
		weatherKey = os.Getenv("WEATHER_API_KEY")
	}
	if weatherKey == "" {
		weatherKey = viper.GetString("weatherKey")
	}

	config := &Config{
		WeatherAPIKey: weatherKey,
	}

	if weatherKey == "" {
		return config, fmt.Errorf("weatherKey não encontrada (use weatherKey ou WEATHER_API_KEY)")
	}

	return config, nil
}
