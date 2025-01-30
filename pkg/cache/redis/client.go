package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tokenoff03/lib_ad1lek/pkg/cache"
)

type client struct {
	pool *redis.Client
}

func NewClient(dsn string) cache.RedisClient {
	redisClient := redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	return &client{
		pool: redisClient,
	}
}

func (c *client) HashSet(ctx context.Context, key string, values map[string]interface{}) error {
	return c.pool.HSet(ctx, key, values).Err()
}

func (c *client) Set(ctx context.Context, key string, value interface{}) error {
	return c.pool.Set(ctx, key, value, 0).Err()
}

func (c *client) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return c.pool.HGetAll(ctx, key).Result()
}

func (c *client) HGet(ctx context.Context, key string, field string) (string, error) {
	return c.pool.HGet(ctx, key, field).Result()
}

func (c *client) Get(ctx context.Context, key string) (string, error) {
	return c.pool.Get(ctx, key).Result()
}

func (c *client) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return c.pool.Expire(ctx, key, expiration).Err()
}

func (c *client) Ping(ctx context.Context) error {
	return c.pool.Ping(ctx).Err()
}

func (c *client) Close() error {
	return c.pool.Close()
}
