package repositories

type RateLimiterRepositoryInterface interface {
	Increment(key string) (int32, error)
}
