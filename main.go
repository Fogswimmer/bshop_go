package main

import (
	dbconfig "api/train/db/config"
	"api/train/routes"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg := dbconfig.LoadDBConfig()

	db, err := sql.Open("postgres", cfg.PostgresDSN())
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("DB unreachable:", err)
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8100"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routes.SetupBookRoutes(r, db)
	routes.SetupAuthorRoutes(r, db)

	r.Run(fmt.Sprintf(":%d", cfg.ServerPort))
}
