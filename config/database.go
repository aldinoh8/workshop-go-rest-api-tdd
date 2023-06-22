package config

import (
	"log"
	"workshoptdd/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database!")
	}

	db.AutoMigrate(&entity.Task{})

	return db
}
