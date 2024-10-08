package main

import (
	"github.com/antoniofmoraes/rate-limiter/internals/infra/repositories"
	"github.com/antoniofmoraes/rate-limiter/internals/infra/webserver"
	"github.com/redis/go-redis/v9"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, // use default DB
	})

	repositories.NewRateLimiterRedisRepository(redisClient)

	webserver.Start()
}
