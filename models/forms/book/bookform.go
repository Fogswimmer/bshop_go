package bookform

type BookForm struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	ReleaseYear int     `json:"release_year"`
	Summary     string  `json:"summary"`
	Price       float64 `json:"price"`
}
