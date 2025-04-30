package models

type Book struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	ReleaseDate int     `json:"release_date"`
	Summary     string  `json:"summary"`
	Price       float64 `json:"price"`
	AuthorID    int     `json:"author_id"`
}

type Author struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Birthday  string `json:"birthday"`
}
