package env

import (
	"os"
	"velocity-technical-test/pkg/logger"

	"github.com/joho/godotenv"
)

type Env struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	ServerPort string
	RedisHost  string
	RedisPort  string
}

func LoadEnv() *Env {
	log := logger.NewLogger()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := &Env{
		DBHost:     getEnv("MYSQL_DB_HOST", log),
		DBPort:     getEnv("MYSQL_DB_PORT", log),
		DBName:     getEnv("MYSQL_DB_NAME", log),
		DBUser:     getEnv("MYSQL_DB_USER", log),
		DBPassword: getEnv("MYSQL_DB_PASSWORD", log),
		ServerPort: getEnv("SERVER_PORT", log),
		RedisHost:  getEnv("REDIS_HOST", log),
		RedisPort:  getEnv("REDIS_PORT", log),
	}

	return env
}

func getEnv(key string, log logger.ILogger) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal("Environment variable %s not set", key)
	}
	return value
}
