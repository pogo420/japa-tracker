package routes

import (
	"fmt"
	"japa-tracker/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetJapaCountByDate(c *gin.Context) {
	c.String(http.StatusOK, "Getting japa count for: "+c.Param("date"))
}

func AddJapaCountByDate(c *gin.Context) {
	var request models.DailyJapaCount
	err := c.BindJSON(&request)
	fmt.Println(request)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid request")
	}
	c.String(http.StatusOK, "Adding japa count")
}

func GetJapaCountTillDate(c *gin.Context) {
	c.String(http.StatusOK, "Getting japa counts till: "+c.Param("date"))
}
