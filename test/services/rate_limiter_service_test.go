package services_test

import (
	"testing"
	"time"

	"github.com/antoniofmoraes/rate-limiter/internals/services"
	repositories_test "github.com/antoniofmoraes/rate-limiter/test/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup() (*services.RateLimiterService, *repositories_test.RateLimiterRepositoryMock) {
	repo := &repositories_test.RateLimiterRepositoryMock{}
	return services.NewRateLimiterService(repo, 10*time.Minute, 5, 10), repo
}

func TestRateLimiterService_IsAllowed(t *testing.T) {
	testCases := []struct {
		id                string
		isToken           bool
		incrementResponse int32
		expected          bool
	}{
		{"229.138.187.150", false, 3, true},
		{"158.84.114.140", false, 5, true},
		{"116.186.186.30", false, 6, false},
		{"abc123", true, 3, true},
		{"defg45", true, 7, true},
		{"hijk6", true, 10, true},
		{"xyz789", true, 11, false},
	}

	service, repo := setup()

	for _, testCase := range testCases {
		repo.On("Increment", testCase.id).Return(testCase.incrementResponse, nil)
		repo.On("Expire", testCase.id, mock.AnythingOfType("time.Duration")).Return(true)
		result, err := service.IsAllowed(testCase.id, testCase.isToken)
		if testCase.expected {
			assert.Nil(t, err)
		} else {
			assert.Error(t, err, "you have reached the maximum number of requests or actions allowed within a certain time frame")
		}
		assert.Equal(t, testCase.expected, result)
	}
}
