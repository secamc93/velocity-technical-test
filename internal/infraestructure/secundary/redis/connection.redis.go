package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"
	"velocity-technical-test/internal/domain/orders/dtos"
	"velocity-technical-test/pkg/env"
	"velocity-technical-test/pkg/logger"

	"github.com/go-redis/redis/v8"
)

var (
	ctx         = context.Background()
	redisClient *redis.Client
	redisOnce   sync.Once
)

func NewRedisClient() *redis.Client {
	redisOnce.Do(func() {
		log := logger.NewLogger()
		env := env.LoadEnv()

		redisClient = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", env.RedisHost, env.RedisPort),
			Password: "",
			DB:       0,
		})

		_, err := redisClient.Ping(ctx).Result()
		if err != nil {
			log.Fatal("Failed to connect to Redis: %v", err)
		}

		log.Success("Connected to Redis at %s:%s", env.RedisHost, env.RedisPort)
	})

	return redisClient
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

	// Establecer el tiempo de vida a 24 horas (24 * 60 * 60 segundos)
	err = s.rdb.Set(ctx, key, data, 24*time.Hour).Err()
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
	// Iniciar una transacción
	txf := func(tx *redis.Tx) error {
		var order dtos.Order

		// Obtener el valor actual
		data, err := tx.Get(ctx, key).Result()
		if err == redis.Nil {
			return fmt.Errorf("key does not exist")
		} else if err != nil {
			return fmt.Errorf("failed to get JSON from Redis: %v", err)
		}

		err = json.Unmarshal([]byte(data), &order)
		if err != nil {
			return fmt.Errorf("failed to unmarshal JSON: %v", err)
		}

		// Actualizar el estado
		order.Status = newStatus

		updatedData, err := json.Marshal(order)
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %v", err)
		}

		// Establecer el valor actualizado
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, key, updatedData, 0)
			return nil
		})
		return err
	}

	// Ejecutar la transacción
	for retries := 0; retries < 1000; retries++ {
		err := s.rdb.Watch(ctx, txf, key)
		if err == nil {
			// Éxito
			return nil
		}
		if err == redis.TxFailedErr {
			// Reintentar
			continue
		}
		// Error diferente
		return err
	}
	return fmt.Errorf("failed to update order status after multiple retries")
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
