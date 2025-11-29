package middleware

import (
	"github.com/breinoso2006/fullcycle/pos-graduacao/desafios-finais/rate-limiter/internal/limiter"
	"github.com/gofiber/fiber/v3"
)

func RateLimit(rateLimiter *limiter.RateLimiter) fiber.Handler {
	return func(c fiber.Ctx) error {

		ip := c.IP()

		token := c.Get("API_KEY")

		allowed, err := rateLimiter.Allow(ip, token)
		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "internal server error",
			})
		}

		if !allowed {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "you have reached the maximum number of requests or actions allowed within a certain time frame",
			})
		}

		return c.Next()
	}
}
