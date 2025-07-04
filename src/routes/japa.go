// Package contains all route logics.
// Mostly processing logic pre and post query
package routes

import (
	"fmt"
	"japa-tracker/src/repository"
	"japa-tracker/src/schema"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetJapaCountByDate(c *gin.Context) {
	fmt.Println("Getting japa count for: " + c.Param("date"))
	db := c.MustGet("db").(*gorm.DB)
	layout := "2006-01-02"
	t, _ := time.Parse(layout, c.Param("date"))
	dbResponse := repository.GetJapaCount(t, db)
	c.String(
		http.StatusOK, "japa count: "+fmt.Sprint(dbResponse))

}

func AddJapaCountByDate(c *gin.Context) {
	var request schema.DailyJapaCount
	err := c.BindJSON(&request)
	fmt.Println(request)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid request")
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	_, err = repository.AddJapaCount(request, db)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid request")
		return
	}
	c.String(http.StatusOK, "Adding japa count")
}

func GetJapaCountTillDate(c *gin.Context) {
	fmt.Println("Getting japa count till: " + c.Param("date"))
	db := c.MustGet("db").(*gorm.DB)
	layout := "2006-01-02"
	t, _ := time.Parse(layout, c.Param("date"))
	dbResponse := repository.GetJapaCountTill(t, db)
	c.String(
		http.StatusOK, "japa count: "+fmt.Sprint(dbResponse))
}
