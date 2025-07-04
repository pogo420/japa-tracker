package db

import (
	"fmt"
	"japa-tracker/src/config"
	"japa-tracker/src/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Low level connection logic for postgres
type PostgresConnection struct {
}

// Low level connection logic
func (pg PostgresConnection) InitDb() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.AppConfig.DbHost,
		config.AppConfig.DbUser,
		config.AppConfig.DbPassword,
		config.AppConfig.DbName,
		config.AppConfig.DbPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Issue in db init")
	}

	// TODO: shall migration tools
	// Auto migrate all models
	err = db.AutoMigrate(&models.JapaCount{})
	if err != nil {
		panic("Failed to migrate:" + err.Error())
	}

	// Managing connections stats
	sqlDB, err := db.DB()
	if err != nil {
		panic("Issue in db init")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	fmt.Println("Db stats:", sqlDB.Stats()) // prints pool stats
	return db
}
