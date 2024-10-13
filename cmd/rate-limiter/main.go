package main

import (
	"fmt"
	"time"

	"github.com/antoniofmoraes/rate-limiter/configs"
	"github.com/antoniofmoraes/rate-limiter/internals/infra/repositories"
	"github.com/antoniofmoraes/rate-limiter/internals/infra/webserver"
	"github.com/antoniofmoraes/rate-limiter/internals/services"
	"github.com/redis/go-redis/v9"
)

func main() {
	configs, err := configs.LoadConfig(".", ".env")
	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", configs.DBHost, configs.DBPort),
		Password: configs.DBPassword,
		DB:       0, // use default DB
	})

	rateLimiterRepository := repositories.NewRateLimiterRedisRepository(redisClient)

	rateLimiterService := services.NewRateLimiterService(
		rateLimiterRepository,
		time.Duration(configs.RateLimiterTimeout)*time.Second,
		configs.RateLimiterIpLimit,
		configs.RateLimiterTokenLimit,
	)

	httpServer := webserver.NewHttpServer(rateLimiterService)
	httpServer.Start()
}
