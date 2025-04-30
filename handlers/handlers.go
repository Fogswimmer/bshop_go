// Хендлер - для обработки внешних HTTP запросов и соединения с базой данных
package handlers

import (
	"api/train/services"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Объявляем структуру
type Handler struct {
	DB *sql.DB
}

// Создание экземпляра
func NewHandler(db *sql.DB) *Handler {
	return &Handler{DB: db}
}

// ==== МЕТОДЫ СТРУКТУРЫ ======
// После слова func - ресивер с указателем
func (h *Handler) GetBooks(c *gin.Context) {
	books, err := services.FetchBooks(h.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get books"})
		return
	}
	c.JSON(200, books)
}

func (h *Handler) CreateBook(c *gin.Context) {

	var book services.BookRequest
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	createdBook, err := services.CreateBook(book, h.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Book creation failed: %v", err)})
		return
	}

	c.JSON(http.StatusOK, createdBook)
}
