package mapper

import (
	"api/train/models/entities"
	"api/train/models/response"
	simplemodels "api/train/models/simple"
)

func MapToBookResponse(book *entities.Book) *response.BookResponse {
	if book == nil {
		return nil
	}
	var cover string
	if book.Cover != nil {
		cover = *book.Cover
	}
	return &response.BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		ReleaseYear: book.ReleaseYear,
		Summary:     book.Summary,
		Price:       book.Price,
		Cover:       cover,
		Author:      MapToSimpleAuthor(book.Author),
	}
}

func MapToAuthorResponse(author *entities.Author) *response.AuthorResponse {
	if author == nil {
		return nil
	}
	return &response.AuthorResponse{
		ID:        author.ID,
		Firstname: author.Firstname,
		Lastname:  author.Lastname,
		Birthday:  author.Birthday,
		Books:     MapToSimpleBooksSlice(author.Books),
	}

}

func MapToSimpleAuthor(author entities.Author) *simplemodels.AuthorSimple {
	return &simplemodels.AuthorSimple{
		ID:        author.ID,
		Firstname: author.Firstname,
		Lastname:  author.Lastname,
		Birthday:  author.Birthday,
	}
}

func MapToSimpleBook(book entities.Book) *simplemodels.BookSimple {
	return &simplemodels.BookSimple{
		ID:          book.ID,
		Title:       book.Title,
		ReleaseYear: book.ReleaseYear,
		Summary:     book.Summary,
		Price:       book.Price,
	}
}

func MapToSimpleBooksSlice(books []entities.Book) []simplemodels.BookSimple {
	var simpleBooks []simplemodels.BookSimple
	for _, book := range books {
		simpleBooks = append(simpleBooks, *MapToSimpleBook(book))
	}
	return simpleBooks
}
