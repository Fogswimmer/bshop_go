package response

import (
	simplemodels "api/train/models/simple"
)

type AuthorResponse struct {
	ID        int                       `json:"id"`
	Firstname string                    `json:"firstname"`
	Lastname  string                    `json:"lastname"`
	Birthday  string                    `json:"birthday"`
	Books     []simplemodels.BookSimple `json:"books"`
}

type BookResponse struct {
	ID          int                        `json:"id"`
	Title       string                     `json:"title"`
	ReleaseYear int                        `json:"release_year"`
	Summary     string                     `json:"summary"`
	Price       float64                    `json:"price"`
	Author      *simplemodels.AuthorSimple `json:"author"`
}
