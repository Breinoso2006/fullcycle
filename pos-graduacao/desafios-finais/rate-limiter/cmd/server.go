package main

import (
	"log"

	"github.com/breinoso2006/fullcycle/pos-graduacao/desafios-finais/rate-limiter/internal/config"
	"github.com/breinoso2006/fullcycle/pos-graduacao/desafios-finais/rate-limiter/internal/limiter"
	"github.com/breinoso2006/fullcycle/pos-graduacao/desafios-finais/rate-limiter/internal/middleware"
	"github.com/breinoso2006/fullcycle/pos-graduacao/desafios-finais/rate-limiter/internal/storage"
	"github.com/gofiber/fiber/v3"
)

func main() {
	cfg := config.Load()

	redisStorage, err := storage.NewRedis(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	rateLimiter := limiter.New(redisStorage, cfg.RateLimitIP, cfg.RateLimitToken, cfg.BlockDuration)

	app := fiber.New(fiber.Config{
		AppName: "Rate Limiter API v1.0",
	})

	app.Use(middleware.RateLimit(rateLimiter))

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Rate limiter is working!",
			"status":  "ok",
		})
	})

	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := app.Listen(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
