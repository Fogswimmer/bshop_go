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
	router.POST("/api/book", bh.CreateBook)
}

func SetupAuthorRoutes(router *gin.Engine, db *sql.DB) {
	ah := authorhandler.NewHandler(db)

	router.GET("/api/authors", ah.GetAuthors)
	router.POST("/api/author", ah.CreateAuthor)
}
