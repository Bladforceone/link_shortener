package db

import (
	"go_pro_api/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDB(conf *configs.Config) *DB {
	db, err := gorm.Open(postgres.Open(conf.DB.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &DB{db}
}
