package models

import (
	"time"
)

type JapaCount struct {
	Date  time.Time `gorm:"type:date;unique;not null"`
	Count int64     `gorm:"not null"`
}
