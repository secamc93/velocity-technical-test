package redis

import (
	"context"
	"fmt"
	"velocity-technical-test/pkg/env"
	"velocity-technical-test/pkg/logger"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func NewRedisClient() *redis.Client {
	log := logger.NewLogger()
	env := env.LoadEnv()

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", env.RedisHost, env.RedisPort),
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis: %v", err)
	}

	log.Success("Connected to Redis at %s:%s", env.RedisHost, env.RedisPort)
	return rdb
}
