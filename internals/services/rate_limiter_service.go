package services

import (
	"errors"
	"time"

	"github.com/antoniofmoraes/rate-limiter/internals/infra/repositories"
)

const (
	unexpectedError      = "unexpected error while validating access rate"
	maximumRequestsError = "you have reached the maximum number of requests or actions allowed within a certain time frame"
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

func (s *RateLimiterService) IsAllowed(identifier string, isToken bool) (bool, error) {
	result, err := s.r.Increment(identifier)
	if err != nil {
		return false, errors.New(unexpectedError)
	}

	limit := s.ipRequestLimit
	if isToken {
		limit = s.tokenRequestLimit
	}

	if result > limit+1 {
		return false, errors.New(maximumRequestsError)
	}

	if result == limit+1 {
		s.r.Expire(identifier, s.timeoutDuration)
		return false, errors.New(maximumRequestsError)
	}

	if result == 1 {
		s.r.Expire(identifier, 1)
	}

	return true, nil
}
