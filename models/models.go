package models

type Book struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	ReleaseYear int     `json:"release_year"`
	Summary     string  `json:"summary"`
	Price       float64 `json:"price"`
	Author      Author  `json:"author"`
}

type Author struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birthday  string `json:"birthday"`
	Books     []Book `json:"books"`
}

type BookRequest struct {
	ID          *int    `json:"id"`
	Title       string  `json:"title"`
	ReleaseYear int     `json:"release_year"`
	Summary     string  `json:"summary"`
	Price       float64 `json:"price"`
	AuthorID    int     `json:"author_id"`
}

type AuthorRequest struct {
	ID        *int   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birthday  string `json:"birthday"`
}

func GetAuthorFullName(author Author) string {
	return author.Firstname + " " + author.Lastname
}
