package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"velocity-technical-test/internal/domain/orders/dtos"
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

type RedisService struct {
	rdb *redis.Client
}

func NewRedisService(rdb *redis.Client) *RedisService {
	return &RedisService{rdb: rdb}
}

func (s *RedisService) SetJSON(key string, status string, value interface{}) error {
	type Order struct {
		Status   string      `json:"status"`
		Response interface{} `json:"response"`
	}

	order := Order{
		Status:   status,
		Response: value,
	}

	data, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	err = s.rdb.Set(ctx, key, data, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set JSON in Redis: %v", err)
	}

	return nil
}

func (s *RedisService) GetOrder(key string) (dtos.OrderDTORedis, error) {
	var order dtos.OrderDTORedis

	data, err := s.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return order, fmt.Errorf("key does not exist")
	} else if err != nil {
		return order, fmt.Errorf("failed to get JSON from Redis: %v", err)
	}

	err = json.Unmarshal([]byte(data), &order)
	if err != nil {
		return order, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return order, nil
}

func (s *RedisService) GetOrderStatus(key string) (string, error) {
	var order dtos.Order

	data, err := s.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key does not exist")
	} else if err != nil {
		return "", fmt.Errorf("failed to get JSON from Redis: %v", err)
	}

	err = json.Unmarshal([]byte(data), &order)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return order.Status, nil
}

func (s *RedisService) UpdateOrderStatus(key string, newStatus string) error {
	var order dtos.Order

	data, err := s.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return fmt.Errorf("key does not exist")
	} else if err != nil {
		return fmt.Errorf("failed to get JSON from Redis: %v", err)
	}

	err = json.Unmarshal([]byte(data), &order)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	order.Status = newStatus

	updatedData, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	err = s.rdb.Set(ctx, key, updatedData, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to update JSON in Redis: %v", err)
	}

	return nil
}

func (s *RedisService) KeyExists(key string) (bool, error) {
	exists, err := s.rdb.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("failed to check if key exists in Redis: %v", err)
	}

	return exists > 0, nil
}

func (s *RedisService) GetJSON(key string) (string, error) {
	data, err := s.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key does not exist")
	} else if err != nil {
		return "", fmt.Errorf("failed to get JSON from Redis: %v", err)
	}

	return data, nil
}
