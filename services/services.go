package services

import (
	"api/train/models"
	"database/sql"
	"errors"
	"fmt"
)

type BookRequest struct {
	Title       string  `json:"title"`
	ReleaseDate int     `json:"release_date"`
	Summary     string  `json:"summary"`
	Price       float64 `json:"price"`
	AuthorID    int     `json:"author_id"`
}

func FetchBooks(db *sql.DB) ([]models.Book, error) {
	rows, err := db.Query("SELECT * FROM book")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		err := rows.Scan(&b.ID, &b.Title, &b.ReleaseDate, &b.Summary, &b.Price, &b.AuthorID)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func FindBookById(id int64, db *sql.DB) (models.Book, error) {
	var b models.Book

	row := db.QueryRow("SELECT id, title, release_date, summary, price, author_id from book WHERE id = ?", id)
	if err := row.Scan(&b.ID, &b.Title, &b.ReleaseDate, &b.Summary, &b.Price, &b.AuthorID); err != nil {
		if err == sql.ErrNoRows {
			return b, fmt.Errorf("FindBookById %d: no such book", id)
		}
		return b, fmt.Errorf("FindBookById %d: %v", id, err)
	}
	return b, nil
}

func CreateBook(br BookRequest, db *sql.DB) (int64, error) {
	if br.Title == "" {
		return 0, errors.New("book title is required")
	}

	result, err := db.Exec(
		"INSERT INTO book (title, release_date, summary, price, author_id) VALUES (?, ?, ?, ?, ?)",
		br.Title, br.ReleaseDate, br.Summary, br.Price, br.AuthorID,
	)
	if err != nil {
		return 0, fmt.Errorf("SynthaxError: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("CreateBook: %v", err)
	}

	return id, err
}
