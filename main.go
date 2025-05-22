package main

import (
	dbconfig "api/train/infra/db/config"
	"api/train/infra/db/seeder"
	"api/train/middleware"
	"api/train/routes"
	fileservice "api/train/services/file"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg := dbconfig.LoadDBConfig()
	db := conn(*cfg)
	defer db.Close()

	setupRouter(cfg, db)
	err := handleArgs(db)
	if err != nil {
		log.Fatal(err)
	}
}

func conn(cfg dbconfig.DBConfig) *sql.DB {
	db, err := sql.Open("postgres", cfg.PostgresDSN())
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("DB unreachable:", err)
	}

	return db
}

func setupRouter(cfg *dbconfig.DBConfig, db *sql.DB) {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8100"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(middleware.RequestLogger())

	rDir := fileservice.GetUploadRootDir()
	r.Static("/uploads", rDir)

	routes.SetupRoutes(r, db)
	r.Run(fmt.Sprintf(":%d", cfg.ServerPort))
}

func handleArgs(db *sql.DB) error {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return nil
	}
	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			err := seeder.Seed(db)
			if err != nil {
				os.Exit(1)
				return err
			}
			os.Exit(0)
		default:
			fmt.Printf("Unknown command: %s\n", args[0])
			os.Exit(1)
		}
	}
	return nil
}
