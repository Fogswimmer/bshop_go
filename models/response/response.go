package response

import "api/train/models/forms/forms"

type AuthorResponse struct {
	ID        int              `json:"id"`
	Firstname string           `json:"firstname"`
	Lastname  string           `json:"lastname"`
	Birthday  string           `json:"birthday"`
	Books     []forms.BookForm `json:"books"`
}

type BookResponse struct {
	ID          int               `json:"id"`
	Title       string            `json:"title"`
	ReleaseYear int               `json:"release_year"`
	Summary     string            `json:"summary"`
	Price       float64           `json:"price"`
	Author      *forms.AuthorForm `json:"author"`
}
