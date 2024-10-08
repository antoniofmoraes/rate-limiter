package repositories

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RateLimiterRedisRepository struct {
	client *redis.Client
}

func NewRateLimiterRedisRepository(client *redis.Client) *RateLimiterRedisRepository {
	return &RateLimiterRedisRepository{
		client,
	}
}

func (r *RateLimiterRedisRepository) Increment(key string) (int32, error) {
	ctx := context.Background()

	pipe := r.client.Pipeline()

	count := pipe.Incr(ctx, key)
	if count.Err() != nil {
		return 0, count.Err()
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return 0, err
	}

	return int32(count.Val()), nil
}
