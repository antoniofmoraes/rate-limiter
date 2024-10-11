package services

import (
	"time"

	"github.com/antoniofmoraes/rate-limiter/internals/infra/repositories"
)

type rateLimiterService struct {
	r                 repositories.RateLimiterRepositoryInterface
	timeoutDuration   time.Duration
	ipRequestLimit    int
	tokenRequestLimit int
}

func NewRateLimiterService(repository repositories.RateLimiterRepositoryInterface, timeoutDuration time.Duration, ipRequestLimit int, tokenRequestLimit int) *rateLimiterService {
	return &rateLimiterService{
		repository,
		timeoutDuration,
		ipRequestLimit,
		tokenRequestLimit,
	}
}

// ### TODO ###
// make it transactional
func (s *rateLimiterService) IsAllowed(identifier string, isToken bool) (bool, error) {
	result, err := s.r.Increment(identifier)
	if err != nil {
		return false, nil
	}

	if result == 1 {
		s.r.Expire(identifier, time.Second)
	} else if result == 6 {
		s.r.Expire(identifier, s.timeoutDuration)
	}

	limit := s.ipRequestLimit
	if isToken {
		limit = s.tokenRequestLimit
	}

	return result > int32(limit), nil
}
