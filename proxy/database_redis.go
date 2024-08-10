package proxy

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// DatabaseRedis implements Database interface for Redis database.
type DatabaseRedis struct {
	Client *redis.Client
}

// Get returns data by key.
func (db DatabaseRedis) Get(key string) (string, error) {
	return db.Client.Get(ctx, key).Result()
}

// Set stores value identified by key with expiration timeout.
func (db DatabaseRedis) Set(key string, value string, expiration time.Duration) error {
	return db.Client.Set(ctx, key, value, expiration).Err()
}

// Delete deletes data by key.
func (db DatabaseRedis) Delete(key string) error {
	return db.Client.Del(ctx, key).Err()
}

// Keys returns stored keys filtered by prefix.
func (db DatabaseRedis) Keys(prefix string) ([]string, error) {
	return db.Client.Keys(ctx, prefix + "*").Result()
}

// Health returns an error if database connection cannot be estabilished.
func (db DatabaseRedis) Health() error {
	return db.Client.Ping(ctx).Err()
}
