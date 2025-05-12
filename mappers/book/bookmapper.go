package bookmapper

import (
	"api/train/models/entities"
	"api/train/models/response"
)

func MapToList(book *entities.Book) *response.BookResponse
