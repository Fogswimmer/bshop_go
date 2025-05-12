package dto

type BookDto struct {
	ID          *int    `json:"id"`
	Title       string  `json:"title"`
	ReleaseYear int     `json:"release_year"`
	Summary     string  `json:"summary"`
	Price       float64 `json:"price"`
	AuthorID    int     `json:"author_id"`
}

type AuthorDto struct {
	ID        *int   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birthday  string `json:"birthday"`
}