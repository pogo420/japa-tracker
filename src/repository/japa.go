package repository

import (
	"time"

	"japa-tracker/src/models"
	"japa-tracker/src/schema"

	"gorm.io/gorm"
)

func GetJapaCount(date time.Time, db *gorm.DB) int64 {

	var japaCount models.JapaCount

	db.Where("date = ?", date).First(&japaCount)

	return japaCount.Count
}

func GetJapaCountTill(date time.Time, db *gorm.DB) int64 {

	var japaCount int64

	db.Model(&models.JapaCount{}).Where("date <= ?", date).Select("sum(count)").Scan(&japaCount)

	return japaCount
}

func AddJapaCount(data schema.DailyJapaCount, db *gorm.DB) (int64, error) {

	layout := "2006-01-02"
	t, _ := time.Parse(layout, data.Date)

	japaCount := models.JapaCount{
		Date:  t,
		Count: data.Count,
	}

	result := db.Create(&japaCount)

	return result.RowsAffected, result.Error
}
