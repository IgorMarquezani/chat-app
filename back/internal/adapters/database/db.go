package database

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	ErrUnexpected = errors.New("unexpected error")
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

func GetDbConnection() (*gorm.DB, error) {
	var err error

	dbOnce.Do(func() {
		dbInstance, err = Connect(
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DBNAME"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_SSLMODE"),
			os.Getenv("POSTGRES_TIMEZONE"),
		)
	})

	return dbInstance, err
}

func Connect(host, user, password, dbname, port, sslmode, timeZone string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timeZone)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
}
