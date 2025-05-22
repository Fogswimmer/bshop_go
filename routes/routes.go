package routes

import (
	authorhandler "api/train/handlers/author"
	bookhandler "api/train/handlers/book"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	setupBookRoutes(r, db)
	setupAuthorRoutes(r, db)
	r.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})
}

func setupBookRoutes(r *gin.Engine, db *sql.DB) {
	bh := bookhandler.NewHandler(db)

	r.GET("/api/books", bh.GetBooks)
	r.GET("/api/book/:id", bh.FindBook)
	r.POST("/api/book", bh.CreateBook)
	r.PUT("/api/book/:id", bh.UpdateBook)
	r.DELETE("/api/book/:id", bh.DeleteBook)
	r.POST("/api/book/upload/:id", bh.UploadCover)
}

func setupAuthorRoutes(r *gin.Engine, db *sql.DB) {
	ah := authorhandler.NewHandler(db)

	r.GET("/api/authors", ah.GetAuthors)
	r.GET("/api/author/:id", ah.FindAuthor)
	r.POST("/api/author", ah.CreateAuthor)
	r.PUT("/api/author/:id", ah.UpdateAuthor)
	r.DELETE("/api/author/:id", ah.DeleteAuthor)
}
