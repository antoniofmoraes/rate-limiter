package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type rateLimiterRedisRepository struct {
	client *redis.Client
}

const prefix = "rate_limiter"

func NewRateLimiterRedisRepository(client *redis.Client) *rateLimiterRedisRepository {
	return &rateLimiterRedisRepository{
		client,
	}
}

func (r *rateLimiterRedisRepository) Increment(key string) (int32, error) {
	ctx := context.Background()
	defer ctx.Done()

	pipe := r.client.Pipeline()

	finalKey := formatKey(key)

	count := pipe.Incr(ctx, finalKey)
	if count.Err() != nil {
		return 0, count.Err()
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return 0, err
	}

	return int32(count.Val()), nil
}

func (r *rateLimiterRedisRepository) Expire(key string, duration time.Duration) bool {
	ctx := context.Background()
	defer ctx.Done()

	res := r.client.Expire(ctx, formatKey(key), duration)

	return res.Val()
}

func formatKey(key string) string {
	return fmt.Sprintf("%s:%s", prefix, key)
}
