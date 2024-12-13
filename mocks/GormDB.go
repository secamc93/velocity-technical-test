package mocks

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GormDB struct {
	mock.Mock
	DB *gorm.DB
}

func NewGormDB() *GormDB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{}) // Use an in-memory SQLite database for testing
	return &GormDB{
		DB: db,
	}
}

func (m *GormDB) Model(value interface{}) *gorm.DB {
	args := m.Called(value)
	if args.Get(0) != nil {
		return args.Get(0).(*gorm.DB)
	}
	return m.DB
}

func (m *GormDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	if args.Get(0) != nil {
		return args.Get(0).(*gorm.DB)
	}
	return m.DB
}

func (m *GormDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	callArgs := m.Called(query, args)
	if callArgs.Get(0) != nil {
		return callArgs.Get(0).(*gorm.DB)
	}
	return m.DB
}

func (m *GormDB) Update(column string, value interface{}) *gorm.DB {
	args := m.Called(column, value)
	if args.Get(0) != nil {
		return args.Get(0).(*gorm.DB)
	}
	return m.DB
}

func (m *GormDB) Count(count *int64) *gorm.DB {
	args := m.Called(count)
	if args.Get(0) != nil {
		return args.Get(0).(*gorm.DB)
	}
	return m.DB
}

func (m *GormDB) Pluck(column string, dest interface{}) *gorm.DB {
	args := m.Called(column, dest)
	if args.Get(0) != nil {
		return args.Get(0).(*gorm.DB)
	}
	return m.DB
}
