package main

import (

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	dsn string
}

func connect(config Config) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(config.dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}
