package database

import (
	"fmt"
	"os"

	"github.com/omerfruk/Go-generics-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type config struct {
	host     string
	port     string
	user     string
	password string
	name     string
	sslMode  string
}

func configFromEnv() config {
	return config{
		host:     envOrDefault("DB_HOST", "localhost"),
		port:     envOrDefault("DB_PORT", "5432"),
		user:     envOrDefault("DB_USER", "go_generics"),
		password: envOrDefault("DB_PASSWORD", "go_generics"),
		name:     envOrDefault("DB_NAME", "go_generics"),
		sslMode:  envOrDefault("DB_SSLMODE", "disable"),
	}
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

func DBConnect() error {
	cfg := configFromEnv()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.host,
		cfg.user,
		cfg.password,
		cfg.name,
		cfg.port,
		cfg.sslMode,
	)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return err
}

func DB() *gorm.DB {
	return db
}

func AutoMigrate() error {
	return db.AutoMigrate(model.User{}, model.Book{})
}
