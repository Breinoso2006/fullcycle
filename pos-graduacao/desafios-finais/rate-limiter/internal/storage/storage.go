package storage

import "time"

type Storage interface {
	Allow(key string, limit int, window time.Duration) (bool, error)
	Block(key string, duration time.Duration) error
	IsBlocked(key string) (bool, error)
}
