package dbconfig

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	ServerPort int
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadDBConfig() *DBConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &DBConfig{
		ServerPort: 8080,
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     5432,
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
}

func (c *DBConfig) PostgresDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}
