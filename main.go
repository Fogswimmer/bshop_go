package main

import (
	"api/train/config"
	"api/train/routes"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig()

	db, err := sql.Open("postgres", cfg.PostgresDSN())
	if err != nil {
		log.Fatal("Ошибка подключения к базе:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("База данных недоступна:", err)
	}

	router := gin.Default()
	routes.SetupRoutes(router, db)

	router.Run(fmt.Sprintf(":%d", cfg.ServerPort))
}
