package storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	client *redis.Client
}

func NewRedis(addr, password string, db int) (*RedisStorage, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisStorage{client: rdb}, nil
}

func (r *RedisStorage) Allow(key string, limit int, window time.Duration) (bool, error) {
	ctx := context.Background()
	counterKey := "counter:" + key

	count, err := r.client.Incr(ctx, counterKey).Result()
	if err != nil {
		return false, err
	}

	if count == 1 {
		r.client.Expire(ctx, counterKey, window)
	}

	return count <= int64(limit), nil
}

func (r *RedisStorage) Block(key string, duration time.Duration) error {
	ctx := context.Background()
	blockedKey := "blocked:" + key

	err := r.client.Set(ctx, blockedKey, "1", duration).Err()
	return err
}

func (r *RedisStorage) IsBlocked(key string) (bool, error) {
	ctx := context.Background()
	blockedKey := "blocked:" + key

	exists, err := r.client.Exists(ctx, blockedKey).Result()
	if err != nil {
		return false, err
	}

	return exists > 0, nil
}
