package authorhandler

import (
	"api/train/models"
	authorservice "api/train/services/author"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	DB *sql.DB
}

func NewHandler(db *sql.DB) *AuthorHandler {
	return &AuthorHandler{DB: db}
}

func (h *AuthorHandler) GetAuthors(c *gin.Context) {
	books, err := authorservice.List(h.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get authors"})
		return
	}
	c.JSON(200, books)
}

func (h *AuthorHandler) CreateAuthor(c *gin.Context) {

	var book models.AuthorRequest
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	createdBook, err := authorservice.Create(book, h.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Author creation failed: %v", err)})
		return
	}

	c.JSON(http.StatusOK, createdBook)
}
