package main

import (
	"time"

	"github.com/antoniofmoraes/rate-limiter/internals/infra/repositories"
	"github.com/antoniofmoraes/rate-limiter/internals/infra/webserver"
	"github.com/antoniofmoraes/rate-limiter/internals/services"
	"github.com/redis/go-redis/v9"
)

func main() {
	// ### TODO ###
	// Make it .env configured
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, // use default DB
	})

	rateLimiterRepository := repositories.NewRateLimiterRedisRepository(redisClient)

	// ### TODO ###
	// Make it .env configured
	rateLimiterService := services.NewRateLimiterService(rateLimiterRepository, time.Second*10, 5, 7)

	httpServer := webserver.NewHttpServer(rateLimiterService)
	httpServer.Start()
}
