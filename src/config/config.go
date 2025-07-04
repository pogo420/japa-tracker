// Package has all central logic for creating config object
// Function LoadConfig create a instance of a struct
package config

import (
	"os"

	"github.com/joho/godotenv"
)

// similating enum for Env values
type EnvName string

const (
	Prod EnvName = "Prod"
	Dev  EnvName = "Dev"
)

// Config schema/structure
// In case of new addition, following structure needs to be modified
type Config struct {
	Env        EnvName
	DbUser     string
	DbPassword string
	DbName     string
	DbHost     string
	DbPort     string
}

// Global variable for config
var AppConfig Config

// Generates config object
func LoadConfig() {
	_ = godotenv.Load() // loads .env file

	// Creating object
	AppConfig = Config{
		Env:        EnvName(os.Getenv("ENV")),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbHost:     os.Getenv("DB_HOST"),
		DbName:     os.Getenv("DB_NAME"),
		DbPort:     os.Getenv("DB_PORT"),
	}

}
