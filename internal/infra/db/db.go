package db

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database Добавить другие базы данных по желанию
type Database struct {
	DbPostgres *gorm.DB
}

func NewDatabase(dsn string, logger *zap.Logger) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
		return nil, err
	}
	return &Database{DbPostgres: db}, nil
}
