package routes

import (
	authorhandler "api/train/handlers/author"
	bookhandler "api/train/handlers/book"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupBookRoutes(router *gin.Engine, db *sql.DB) {
	bh := bookhandler.NewHandler(db)

	router.GET("/api/books", bh.GetBooks)
	router.GET("/api/book/:id", bh.FindBook)
	router.POST("/api/book", bh.CreateBook)
	router.PUT("/api/book/:id", bh.UpdateBook)
	router.DELETE("/api/book/:id", bh.DeleteBook)
}

func SetupAuthorRoutes(router *gin.Engine, db *sql.DB) {
	ah := authorhandler.NewHandler(db)

	router.GET("/api/authors", ah.GetAuthors)
	router.GET("/api/author/:id", ah.FindAuthor)
	router.POST("/api/author", ah.CreateAuthor)
	router.PUT("/api/author/:id", ah.UpdateAuthor)
	router.DELETE("/api/author/:id", ah.DeleteAuthor)

}
