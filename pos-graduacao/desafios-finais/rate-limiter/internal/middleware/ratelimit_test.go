package middleware

import (
	"io"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/breinoso2006/fullcycle/pos-graduacao/desafios-finais/rate-limiter/internal/limiter"
	"github.com/gofiber/fiber/v3"
)

type mockStorage struct {
	counters map[string]int
	blocks   map[string]bool
}

func newMockStorage() *mockStorage {
	return &mockStorage{
		counters: make(map[string]int),
		blocks:   make(map[string]bool),
	}
}

func (m *mockStorage) Allow(key string, limit int, window time.Duration) (bool, error) {
	m.counters[key]++
	return m.counters[key] <= limit, nil
}

func (m *mockStorage) Block(key string, duration time.Duration) error {
	m.blocks[key] = true
	return nil
}

func (m *mockStorage) IsBlocked(key string) (bool, error) {
	return m.blocks[key], nil
}

func TestRateLimit_AllowsValidRequest(t *testing.T) {
	mockStore := newMockStorage()
	rateLimiter := limiter.New(mockStore, 10, 100, 5*time.Minute)

	app := fiber.New()
	app.Use(RateLimit(rateLimiter))
	app.Get("/test", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	req := httptest.NewRequest("GET", "/test", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	if string(body) != "OK" {
		t.Errorf("expected body 'OK', got '%s'", string(body))
	}
}

func TestRateLimit_BlocksExcessiveRequests(t *testing.T) {
	mockStore := newMockStorage()
	rateLimiter := limiter.New(mockStore, 3, 100, 5*time.Minute)

	app := fiber.New()
	app.Use(RateLimit(rateLimiter))
	app.Get("/test", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	for i := 0; i < 3; i++ {
		req := httptest.NewRequest("GET", "/test", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != 200 {
			t.Errorf("request %d should be allowed", i+1)
		}
	}

	req := httptest.NewRequest("GET", "/test", nil)
	resp, _ := app.Test(req)
	if resp.StatusCode != 429 {
		t.Errorf("expected status 429, got %d", resp.StatusCode)
	}
}

func TestRateLimit_Returns429(t *testing.T) {
	mockStore := newMockStorage()
	rateLimiter := limiter.New(mockStore, 1, 100, 5*time.Minute)

	app := fiber.New()
	app.Use(RateLimit(rateLimiter))
	app.Get("/test", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	req := httptest.NewRequest("GET", "/test", nil)
	app.Test(req)

	req = httptest.NewRequest("GET", "/test", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != 429 {
		t.Errorf("expected status 429, got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	if len(body) == 0 {
		t.Error("expected error message in response body")
	}
}

func TestRateLimit_TokenPriority(t *testing.T) {
	mockStore := newMockStorage()
	rateLimiter := limiter.New(mockStore, 2, 10, 5*time.Minute)

	app := fiber.New()
	app.Use(RateLimit(rateLimiter))
	app.Get("/test", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	for i := 0; i < 2; i++ {
		req := httptest.NewRequest("GET", "/test", nil)
		app.Test(req)
	}

	req := httptest.NewRequest("GET", "/test", nil)
	resp, _ := app.Test(req)
	if resp.StatusCode != 429 {
		t.Error("IP should be blocked after 2 requests")
	}

	req = httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("API_KEY", "test-token")
	resp, _ = app.Test(req)
	if resp.StatusCode != 200 {
		t.Error("token should be allowed (different limit)")
	}
}
