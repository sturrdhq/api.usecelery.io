package database

import (
	"fmt"
	"log/slog"

	"github.com/sturrdhq/celery-server/config"
	"github.com/sturrdhq/celery-server/internal/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClient struct {
	*gorm.DB
}

func NewDBClient() (*DBClient, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		config.GetEnv("POSTGRES_HOST"),
		config.GetEnv("POSTGRES_USER"),
		config.GetEnv("POSTGRES_PASSWORD"),
		config.GetEnv("POSTGRES_DB"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	slog.Info("Connection to PostgreSQL database established")

	return &DBClient{db}, nil
}

func (db *DBClient) Setup() error {
	err := db.AutoMigrate(
		models.WaitList{},
		models.Subscription{},
	)

	if err != nil {
		return err
	}

	return nil
}
