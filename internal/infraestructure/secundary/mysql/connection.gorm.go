package mysql

import (
	"fmt"
	"sync"
	"velocity-technical-test/pkg/env"
	"velocity-technical-test/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConnection interface {
	GetDB() *gorm.DB
	CloseDB() error
	Reconnect() error
	PingDB() error
}

type dbConnection struct {
	db *gorm.DB
}

var (
	instance *dbConnection
	once     sync.Once
	log      = logger.NewLogger()
)

func NewDBConnection() DBConnection {
	once.Do(func() {
		instance = &dbConnection{}
	})
	return instance
}

func (conn *dbConnection) GetDB() *gorm.DB {
	if conn.db == nil {
		envVars := env.LoadEnv()
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			envVars.DBUser, envVars.DBPassword, envVars.DBHost, envVars.DBPort, envVars.DBName)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect database: %v", err)
			panic("failed to connect database")
		}
		log.Success("Database connection established")
		conn.db = db
	}

	if err := conn.PingDB(); err != nil {
		log.Warn("Database connection lost, attempting to reconnect")
		if err := conn.Reconnect(); err != nil {
			log.Fatal("Failed to reconnect to the database: %v", err)
			panic("Failed to reconnect to the database")
		}
	}
	return conn.db
}

func (conn *dbConnection) CloseDB() error {
	sqlDB, err := conn.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (conn *dbConnection) Reconnect() error {
	envVars := env.LoadEnv()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		envVars.DBUser, envVars.DBPassword, envVars.DBHost, envVars.DBPort, envVars.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to reconnect database: %v", err)
		return err
	}
	conn.db = db
	log.Info("Database reconnected successfully")
	return nil
}

func (conn *dbConnection) PingDB() error {
	sqlDB, err := conn.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}
