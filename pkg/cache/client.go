package cache

import (
	"context"
	"time"
)

type RedisClient interface {
	HashSet(ctx context.Context, key string, values map[string]interface{}) error
	Set(ctx context.Context, key string, value interface{}) error
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	Get(ctx context.Context, key string) (string, error)
	HGet(ctx context.Context, key string, field string) (string, error)
	Expire(ctx context.Context, key string, expiration time.Duration) error
	Ping(ctx context.Context) error
	Close() error
}
