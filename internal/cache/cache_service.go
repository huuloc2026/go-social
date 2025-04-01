package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisCache là struct triển khai Redis
type RedisCache struct {
	client *redis.Client
}

// NewRedisCache khởi tạo RedisCache
func NewRedisCache(redisURL string) *RedisCache {
	options, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err) // Hoặc log.Fatal(err)
	}

	client := redis.NewClient(options)

	return &RedisCache{
		client: client,
	}
}

// Set lưu dữ liệu vào Redis
func (r *RedisCache) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

// Get lấy dữ liệu từ Redis
func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

// Delete xóa một key trong Redis
func (r *RedisCache) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
