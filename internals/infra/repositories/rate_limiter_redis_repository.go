package repositories

import (
	"context"
	"fmt"

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

	pipe := r.client.Pipeline()

	finalKey := fmt.Sprintf("%s:%s", prefix, key)

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
