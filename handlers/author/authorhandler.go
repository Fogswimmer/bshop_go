package authorhandler

import (
	"api/train/helpers"
	"api/train/models/dto"
	authorservice "api/train/services/author"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

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
		c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to list authors: %v", err)})
		return
	}
	c.JSON(200, books)
}

func (h *AuthorHandler) CreateAuthor(c *gin.Context) {

	var dto dto.AuthorDto
	if !helpers.ValidateJSON(c, &dto) {
		return
	}

	createdAuthor, err := authorservice.Create(dto, h.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Author creation failed: %v", err)})
		return
	}

	c.JSON(http.StatusOK, createdAuthor)
}

func (h *AuthorHandler) UpdateAuthor(c *gin.Context) {
	id := c.Param("id")
	authorId, _ := strconv.Atoi(id)
	var dto dto.AuthorDto
	if !helpers.ValidateJSON(c, &dto) {
		return
	}
	err := authorservice.Update(authorId, dto, h.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Author updating failed: %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": fmt.Sprintf("Author with ID %d successfully edited", authorId)})
}
