package bookhandler

import (
	"api/train/models"
	bookservice "api/train/services/book"
	"strconv"

	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	DB *sql.DB
}

func NewHandler(db *sql.DB) *BookHandler {
	return &BookHandler{DB: db}
}

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

	id, err := bookservice.Create(book, h.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Book creation failed: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": fmt.Sprintf("Book with ID %d successfully created", id)})
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	bookId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No id provided in the request"})
		return
	}
	var book models.BookRequest
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	err = bookservice.Update(bookId, book, h.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Book updating failed: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": fmt.Sprintf("Book with ID %d successfully edited", bookId)})
}
