package repositories

import "time"

type RateLimiterRepositoryInterface interface {
	Increment(key string) (int32, error)
	Expire(key string, duration time.Duration) bool
}
