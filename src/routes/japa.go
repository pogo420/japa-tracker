// Package contains all route logics.
// Mostly processing logic pre and post query
package routes

import (
	"fmt"
	"japa-tracker/src/repository"
	"japa-tracker/src/schema"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetJapaCountByDate(c *gin.Context) {
	logger := c.MustGet("logger").(*slog.Logger)
	logger.Info("Received request", "Query date", c.Param("date"))
	db := c.MustGet("db").(*gorm.DB)
	layout := "2006-01-02"
	t, _ := time.Parse(layout, c.Param("date"))
	dbResponse := repository.GetJapaCount(t, db)
	logger.Debug("Db response", "Japa count", dbResponse)
	c.String(
		http.StatusOK, "japa count: "+fmt.Sprint(dbResponse))

}

func AddJapaCountByDate(c *gin.Context) {
	var request schema.DailyJapaCount
	logger := c.MustGet("logger").(*slog.Logger)
	err := c.BindJSON(&request)
	fmt.Println(request)
	logger.Info("Received request", "Payload", request)
	if err != nil {
		logger.Error("Error - Invalid payload request")
		c.String(http.StatusBadRequest, "Invalid request")
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	_, err = repository.AddJapaCount(request, db)
	if err != nil {
		logger.Error("Error - Issue in processing japa count query")
		c.String(http.StatusBadRequest, "Invalid request")
		return
	}
	logger.Debug("Adding japa count", "data", request)
	c.String(http.StatusOK, "Adding japa count")
}

func GetJapaCountTillDate(c *gin.Context) {
	logger := c.MustGet("logger").(*slog.Logger)
	logger.Info("Received request for aggregating japa counts", "Query date", c.Param("date"))
	db := c.MustGet("db").(*gorm.DB)
	layout := "2006-01-02"
	t, _ := time.Parse(layout, c.Param("date"))
	dbResponse := repository.GetJapaCountTill(t, db)
	logger.Info("Db resposen", "Aggregated count", dbResponse)
	c.String(
		http.StatusOK, "japa count: "+fmt.Sprint(dbResponse))
}
