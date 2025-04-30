package models

type Book struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	ReleaseYear int     `json:"release_year"`
	Summary     string  `json:"summary"`
	Price       float64 `json:"price"`
	AuthorID    int     `json:"author_id"`
	AuthorName  string  `json:"author_name"`
}

type Author struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birthday  string `json:"birthday"`
}

type BookRequest struct {
	Title       string  `json:"title"`
	ReleaseYear int     `json:"release_year"`
	Summary     string  `json:"summary"`
	Price       float64 `json:"price"`
	AuthorID    *int64  `json:"author_id"`
}

type AuthorRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birthday  string `json:"birthday"`
}
