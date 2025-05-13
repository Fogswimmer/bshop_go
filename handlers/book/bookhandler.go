package bookhandler

import (
	"api/train/helpers"
	"api/train/models/dto"
	bookservice "api/train/services/book"
	fileservice "api/train/services/file"
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
		c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to list books: %v", err)})
		return
	}
	c.JSON(200, books)
}

func (h *BookHandler) FindBook(c *gin.Context) {
	id := c.Param("id")
	bookId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No id provided in the request"})
		return
	}

	b, err := bookservice.Find(bookId, h.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Book not found: %v", err)})
		return
	}

	c.JSON(200, b)
}

func (h *BookHandler) CreateBook(c *gin.Context) {

	var dto dto.BookDto
	if !helpers.ValidateJSON(c, &dto) {
		return
	}

	id, err := bookservice.Create(dto, h.DB)
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
	var dto dto.BookDto
	if !helpers.ValidateJSON(c, &dto) {
		return
	}

	err = bookservice.Update(bookId, dto, h.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Book updating failed: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": fmt.Sprintf("Book with ID %d successfully edited", bookId)})
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	bookId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No id provided in the request"})
		return
	}

	err = bookservice.Delete(bookId, h.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Book deleting failed: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": fmt.Sprintf("Book with ID %d successfully deleted", bookId)})
}

func (h *BookHandler) UploadCover(c *gin.Context) {
	id := c.Param("id")
	bookId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No id provided in the request"})
		return
	}

	b, err := bookservice.Find(bookId, h.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Book not found: %v", err)})
		return

	}

	file, err := c.FormFile("cover")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file provided in the request"})
		return
	}

	absDir := fileservice.GetAbsUploadsSubDir("books", b.Title)
	path, err := fileservice.UploadFile(c, file, absDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to upload cover: %v", err)})
		return
	}

	b, err = bookservice.SaveCover(bookId, path, file.Filename, h.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to save to DB: %v", err)})
		return
	}

	c.JSON(200, b)
}
