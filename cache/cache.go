package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	client *redis.Client
}

func NewCache(address string) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr: address,
	})
	return &Cache{client: client}
}

func (c *Cache) Get(key string) (string, error) {
	return c.client.Get(context.Background(), key).Result()
}

func (c *Cache) Set(key string, value string, ttl time.Duration) error {
	return c.client.Set(context.Background(), key, value, ttl).Err()
}
