package middlewares

// import (
// 	"context"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/limiter"
// 	"github.com/huuloc2026/go-social/config"
// 	"github.com/redis/go-redis/v9"
// )

// var ctx = context.Background()

// // RedisStorage implements fiber.Storage interface for RateLimiter
// type RedisStorage struct {
// 	client *redis.Client
// }

// // NewRedisStorage khởi tạo Redis client
// func NewRedisStorage(cfg *config.Config) *RedisStorage {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
// 		Password: cfg.RedisPassword,
// 		DB:       0,
// 	})

// 	return &RedisStorage{client: rdb}
// }

// // Get lấy giá trị từ Redis
// func (s *RedisStorage) Get(key string) (int, error) {
// 	val, err := s.client.Get(ctx, key).Int()
// 	if err == redis.Nil {
// 		return 0, nil
// 	}
// 	return val, err
// }

// // Set đặt giá trị vào Redis với thời gian hết hạn
// func (s *RedisStorage) Set(key string, val int, exp time.Duration) error {
// 	return s.client.Set(ctx, key, val, exp).Err()
// }

// // Delete xoá key khỏi Redis
// func (s *RedisStorage) Delete(key string) error {
// 	return s.client.Del(ctx, key).Err()
// }

// // Reset xoá toàn bộ dữ liệu trong Redis
// func (s *RedisStorage) Reset() error {
// 	return s.client.FlushDB(ctx).Err()
// }

// // RateLimiter middleware sử dụng Redis làm storage
// func RateLimiter(cfg *config.Config) fiber.Handler {
// 	// redisStorage := NewRedisStorage(cfg)

// 	return limiter.New(limiter.Config{
// 		Max:        cfg.RateLimit,
// 		Expiration: cfg.RateLimitWindow,
// 		KeyGenerator: func(c *fiber.Ctx) string {
// 			return c.IP() // Giới hạn theo IP
// 		},
// 		LimitReached: func(c *fiber.Ctx) error {
// 			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
// 				"error":   "too many requests",
// 				"message": "Please try again later",
// 			})
// 		},
// 		// Storage: redisStorage, // Dùng Redis làm backend lưu trữ
// 	})
// }
