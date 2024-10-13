package services

import (
	"errors"
	"time"

	"github.com/antoniofmoraes/rate-limiter/internals/infra/repositories"
)

type RateLimiterService struct {
	r                 repositories.RateLimiterRepositoryInterface
	timeoutDuration   time.Duration
	ipRequestLimit    int32
	tokenRequestLimit int32
}

func NewRateLimiterService(repository repositories.RateLimiterRepositoryInterface, timeoutDuration time.Duration, ipRequestLimit int32, tokenRequestLimit int32) *RateLimiterService {
	return &RateLimiterService{
		repository,
		timeoutDuration,
		ipRequestLimit,
		tokenRequestLimit,
	}
}

// ### TODO ###
// make it transactional
func (s *RateLimiterService) IsAllowed(identifier string, isToken bool) (bool, error) {
	result, err := s.r.Increment(identifier)
	if err != nil {
		return false, errors.New("unexpected error while validating access rate")
	}

	limit := s.ipRequestLimit
	if isToken {
		limit = s.tokenRequestLimit
	}

	if result > limit {
		return false, errors.New("you have reached the maximum number of requests or actions allowed within a certain time frame")
	}

	if result == limit {
		s.r.Expire(identifier, s.timeoutDuration)
		return false, errors.New("you have reached the maximum number of requests or actions allowed within a certain time frame")
	}

	if result == 1 {
		s.r.Expire(identifier, 1)
	}

	return true, nil
}
