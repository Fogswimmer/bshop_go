package entities

type Book struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	ReleaseYear int     `json:"release_year"`
	Summary     string  `json:"summary"`
	Price       float64 `json:"price"`
	Author      Author  `json:"author"`
	Cover       string  `json:"cover"`
}

type Author struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birthday  string `json:"birthday"`
	Books     []Book `json:"books"`
	Avatar    string `json:"avatar"`
}

func GetAuthorFullName(author Author) string {
	return author.Firstname + " " + author.Lastname
}
