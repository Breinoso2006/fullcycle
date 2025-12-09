package limiter

import (
	"testing"
	"time"
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

func TestRateLimiter_AllowIP(t *testing.T) {
	mockStore := newMockStorage()
	limiter := New(mockStore, 5, 100, 5*time.Minute)

	t.Run("should allow IP within limit", func(t *testing.T) {
		ip := "192.168.1.1"

		for i := 0; i < 5; i++ {
			allowed, err := limiter.AllowIP(ip)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !allowed {
				t.Errorf("request %d should be allowed", i+1)
			}
		}
	})

	t.Run("should block IP after limit exceeded", func(t *testing.T) {
		mockStore := newMockStorage()
		limiter := New(mockStore, 3, 100, 5*time.Minute)
		ip := "192.168.1.2"

		for i := 0; i < 3; i++ {
			limiter.AllowIP(ip)
		}

		allowed, _ := limiter.AllowIP(ip)
		if allowed {
			t.Error("request should be blocked after limit exceeded")
		}
	})

	t.Run("should keep IP blocked during block duration", func(t *testing.T) {
		mockStore := newMockStorage()
		limiter := New(mockStore, 2, 100, 200*time.Millisecond)
		ip := "192.168.1.3"

		limiter.AllowIP(ip)
		limiter.AllowIP(ip)
		limiter.AllowIP(ip)

		allowed, _ := limiter.AllowIP(ip)
		if allowed {
			t.Error("IP should remain blocked")
		}
	})
}

func TestRateLimiter_AllowToken(t *testing.T) {
	mockStore := newMockStorage()
	limiter := New(mockStore, 5, 100, 5*time.Minute)

	t.Run("should allow token within limit", func(t *testing.T) {
		token := "abc123"

		for i := 0; i < 100; i++ {
			allowed, err := limiter.AllowToken(token)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !allowed {
				t.Errorf("request %d should be allowed", i+1)
			}
		}
	})

	t.Run("should block token after limit exceeded", func(t *testing.T) {
		mockStore := newMockStorage()
		limiter := New(mockStore, 5, 10, 5*time.Minute)
		token := "xyz789"

		for i := 0; i < 10; i++ {
			limiter.AllowToken(token)
		}

		allowed, _ := limiter.AllowToken(token)
		if allowed {
			t.Error("request should be blocked after limit exceeded")
		}
	})
}

func TestRateLimiter_Allow(t *testing.T) {
	t.Run("should prioritize token over IP", func(t *testing.T) {
		mockStore := newMockStorage()
		limiter := New(mockStore, 5, 100, 5*time.Minute)
		ip := "192.168.1.4"
		token := "token123"

		for i := 0; i < 6; i++ {
			limiter.Allow(ip, "")
		}

		allowed, _ := limiter.Allow(ip, "")
		if allowed {
			t.Error("IP should be blocked (exceeded 5 req/s)")
		}

		allowed, _ = limiter.Allow(ip, token)
		if !allowed {
			t.Error("token should be allowed (has 100 req/s limit)")
		}
	})

	t.Run("should use IP limit when no token provided", func(t *testing.T) {
		mockStore := newMockStorage()
		limiter := New(mockStore, 3, 100, 5*time.Minute)
		ip := "192.168.1.5"

		for i := 0; i < 3; i++ {
			allowed, _ := limiter.Allow(ip, "")
			if !allowed {
				t.Errorf("request %d should be allowed", i+1)
			}
		}

		allowed, _ := limiter.Allow(ip, "")
		if allowed {
			t.Error("request should be blocked (exceeded IP limit)")
		}
	})
}
