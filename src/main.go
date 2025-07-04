package main

import (
	"japa-tracker/src/config"
	"japa-tracker/src/db"
	"japa-tracker/src/routes"
	"japa-tracker/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func rootEndpoint(c *gin.Context) {
	c.String(http.StatusOK, utils.WELCOME_MESSAGE)
}

func main() {
	// Initializes config object
	config.LoadConfig()

	// Initializing gin
	router := gin.Default()

	// Initializing db
	var db_connection db.DbConnection = db.PostgresConnection{}
	db := db_connection.InitDb()

	// Passing DB to routes via middleware
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// All routes are defined below
	router.GET("/", rootEndpoint)
	router.GET("/japa/:date", routes.GetJapaCountByDate)
	router.POST("/japa", routes.AddJapaCountByDate)
	router.GET("/japa-count/:date", routes.GetJapaCountTillDate)

	// Starting the gin server
	router.Run()
}
