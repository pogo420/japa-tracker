// Package contains all logics for db connection setup
package db

import "gorm.io/gorm"

// Defining protocol for db connection
type DbConnection interface {
	InitDb() *gorm.DB
}
