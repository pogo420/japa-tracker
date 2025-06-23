package main

import (
	"fmt"
	"japa-tracker/src/models"
	"japa-tracker/src/routes"
	"japa-tracker/src/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func rootEndpoint(c *gin.Context) {
	c.String(http.StatusOK, utils.WELCOME_MESSAGE)
}

func init_db() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	fmt.Print(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Issue in db init")
	}

	// Auto migrate all models
	err = db.AutoMigrate(&models.JapaCount{})
	if err != nil {
		panic("Failed to migrate:" + err.Error())
	}
	return db
}

func main() {
	router := gin.Default()
	_ = godotenv.Load() // loads .env file
	db_connection := init_db()

	// Pass DB to routes via middleware
	router.Use(func(c *gin.Context) {
		c.Set("db", db_connection)
		c.Next()
	})
	router.GET("/", rootEndpoint)

	router.GET("/japa/:date", routes.GetJapaCountByDate)
	router.POST("/japa", routes.AddJapaCountByDate)

	router.GET("/japa-count/:date", routes.GetJapaCountTillDate)

	router.Run()
}
