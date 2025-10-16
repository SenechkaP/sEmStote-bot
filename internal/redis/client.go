package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Incr(ctx context.Context, key string) *redis.IntCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
}

type Client struct {
	redis *redis.Client
}

func New(addr, password string) *Client {
	return &Client{
		redis: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
		}),
	}
}

func (c *Client) Get(ctx context.Context, key string) *redis.StringCmd {
	return c.redis.Get(ctx, key)
}

func (c *Client) Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd {
	return c.redis.Set(ctx, key, value, expiration)
}

func (c *Client) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	return c.redis.Del(ctx, keys...)
}

func (c *Client) Incr(ctx context.Context, key string) *redis.IntCmd {
	return c.redis.Incr(ctx, key)
}

func (c *Client) Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	return c.redis.Expire(ctx, key, expiration)
}
