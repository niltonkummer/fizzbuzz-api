package repository

import (
	"errors"

	"github.com/go-redis/redis/v8"
)

type CacheRedis struct {
	client *redis.Client
}

// NewCacheRedis creates a new instance of CacheRedis
func NewCacheRedis(client *redis.Client) *CacheRedis {
	return &CacheRedis{
		client: client,
	}
}

// Get retrieves a value from the Redis cache by key
func (c *CacheRedis) Get(key string) (string, error) {
	ctx := c.client.Context()
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		}
		return "", err
	}
	return val, nil
}

// Set stores a value in the Redis cache with a key
func (c *CacheRedis) Set(key string, value string) error {
	ctx := c.client.Context()
	err := c.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
