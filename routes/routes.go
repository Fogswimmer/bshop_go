package routes

import (
	"api/train/handlers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	h := handlers.NewHandler(db)

	router.GET("/api/books", h.GetBooks)
	router.POST("/api/book", h.CreateBook)
}
