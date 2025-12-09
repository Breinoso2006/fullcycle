package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	RateLimitIP    int
	RateLimitToken int
	BlockDuration  time.Duration
	RedisAddr      string
	RedisPassword  string
	RedisDB        int
	ServerPort     string
}

func Load() *Config {

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.SetDefault("RATE_LIMIT_IP", 10)
	viper.SetDefault("RATE_LIMIT_TOKEN", 100)
	viper.SetDefault("BLOCK_DURATION", 300)
	viper.SetDefault("REDIS_ADDR", "localhost:6379")
	viper.SetDefault("REDIS_PASSWORD", "")
	viper.SetDefault("REDIS_DB", 0)
	viper.SetDefault("SERVER_PORT", "8080")

	if err := viper.ReadInConfig(); err != nil {

	}

	return &Config{
		RateLimitIP:    viper.GetInt("RATE_LIMIT_IP"),
		RateLimitToken: viper.GetInt("RATE_LIMIT_TOKEN"),
		BlockDuration:  time.Duration(viper.GetInt("BLOCK_DURATION")) * time.Second,
		RedisAddr:      viper.GetString("REDIS_ADDR"),
		RedisPassword:  viper.GetString("REDIS_PASSWORD"),
		RedisDB:        viper.GetInt("REDIS_DB"),
		ServerPort:     viper.GetString("SERVER_PORT"),
	}
}
