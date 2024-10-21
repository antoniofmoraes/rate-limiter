package repositories_test

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type RateLimiterRepositoryMock struct {
	mock.Mock
}

func (m *RateLimiterRepositoryMock) Increment(key string) (int32, error) {
	args := m.Called(key)
	return args.Get(0).(int32), args.Error(1)
}
func (m *RateLimiterRepositoryMock) Expire(key string, duration time.Duration) bool {
	args := m.Called(key, duration)
	return args.Bool(0)
}
