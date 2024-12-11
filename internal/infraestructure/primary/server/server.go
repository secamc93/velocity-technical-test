package server

import (
	"fmt"
	"net/http"
	"time"
	"velocity-technical-test/internal/infraestructure/primary/gingonic/routers"
	"velocity-technical-test/internal/infraestructure/secundary/mysql"
	"velocity-technical-test/internal/infraestructure/secundary/mysql/migrate"
	"velocity-technical-test/internal/infraestructure/secundary/redis"
	"velocity-technical-test/pkg/env"
	"velocity-technical-test/pkg/logger"
)

func RunServer() {
	log := logger.NewLogger()
	dbConn := mysql.NewDBConnection()
	defer func() {
		if err := dbConn.CloseDB(); err != nil {
			log.Error("Error closing DB: %v", err)
		}
	}()
	migrate.Migrate()

	redisClient := redis.NewRedisClient()
	defer func() {
		if err := redisClient.Close(); err != nil {
			log.Error("Error closing Redis: %v", err)
		}
	}()

	port := env.LoadEnv().ServerPort
	if port == "" {
		port = "8080"
	}

	address := fmt.Sprintf(":%s", port)
	log.Success("Server running on port %s", port)

	router := routers.SetupRouter()

	server := &http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Error("Failed to start server: %v", err)
	}
}
