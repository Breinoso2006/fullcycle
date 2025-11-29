package limiter

import (
	"time"

	"github.com/breinoso2006/fullcycle/pos-graduacao/desafios-finais/rate-limiter/internal/storage"
)

type RateLimiter struct {
	storage       storage.Storage
	ipLimit       int
	tokenLimit    int
	blockDuration time.Duration
}

func New(storage storage.Storage, ipLimit, tokenLimit int, blockDuration time.Duration) *RateLimiter {
	return &RateLimiter{
		storage:       storage,
		ipLimit:       ipLimit,
		tokenLimit:    tokenLimit,
		blockDuration: blockDuration,
	}
}

func (rl *RateLimiter) AllowIP(ip string) (bool, error) {
	key := "ip:" + ip

	blocked, err := rl.storage.IsBlocked(key)
	if err != nil {
		return false, err
	}
	if blocked {
		return false, nil
	}

	allowed, err := rl.storage.Allow(key, rl.ipLimit, 1*time.Second)
	if err != nil {
		return false, err
	}

	if !allowed {
		rl.storage.Block(key, rl.blockDuration)
		return false, nil
	}

	return true, nil
}

func (rl *RateLimiter) AllowToken(token string) (bool, error) {
	key := "token:" + token

	blocked, err := rl.storage.IsBlocked(key)
	if err != nil {
		return false, err
	}
	if blocked {
		return false, nil
	}

	allowed, err := rl.storage.Allow(key, rl.tokenLimit, 1*time.Second)
	if err != nil {
		return false, err
	}

	if !allowed {
		rl.storage.Block(key, rl.blockDuration)
		return false, nil
	}

	return true, nil
}

func (rl *RateLimiter) Allow(ip, token string) (bool, error) {
	if token != "" {
		return rl.AllowToken(token)
	}

	return rl.AllowIP(ip)
}
