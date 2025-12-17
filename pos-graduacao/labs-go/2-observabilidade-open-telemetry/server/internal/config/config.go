package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port            string `mapstructure:"PORT"`
	OrchestratorURL string `mapstructure:"ORCHESTRATOR_URL"`
	ServiceName     string `mapstructure:"SERVICE_NAME"`
	ZipkinURL       string `mapstructure:"ZIPKIN_URL"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
