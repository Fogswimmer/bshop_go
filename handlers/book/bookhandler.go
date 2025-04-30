// Хендлер - для обработки внешних HTTP запросов и соединения с базой данных
package bookhandler

import (
	"api/train/models"
	bookservice "api/train/services/books"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Объявляем структуру
type BookHandler struct {
	DB *sql.DB
}

// Создание экземпляра
func NewHandler(db *sql.DB) *BookHandler {
	return &BookHandler{DB: db}
}

// ==== МЕТОДЫ СТРУКТУРЫ ======
// После слова func - ресивер с указателем
func (h *BookHandler) GetBooks(c *gin.Context) {
	books, err := bookservice.List(h.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get books"})
		return
	}
	c.JSON(200, books)
}

func (h *BookHandler) CreateBook(c *gin.Context) {

	var book models.BookRequest
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	createdBook, err := bookservice.Create(book, h.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Book creation failed: %v", err)})
		return
	}

	c.JSON(http.StatusOK, createdBook)
}
