package database

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var pool = sync.Pool{}

var (
	ErrUnexpected = errors.New("unexpected error")
)

func GetDbConnection() (*gorm.DB, error) {
	v := pool.Get()
	if v == nil {
		db, err := Connect(os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DBNAME"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_SSLMODE"),
			os.Getenv("POSTGRES_TIMEZONE"))

		if err == nil {
			pool.Put(db)
		}

		return db, err
	}

	db, ok := v.(*gorm.DB)
	if !ok {
		return db, ErrUnexpected
	}

	return db, nil
}

func Connect(host, user, password, dbname, port, sslmode, timeZone string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timeZone)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
}
