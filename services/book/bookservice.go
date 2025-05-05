package bookservice

import (
	"api/train/models"
	authorservice "api/train/services/author"
	"database/sql"
	"errors"
	"fmt"
)

func List(db *sql.DB) ([]models.Book, error) {
	rows, err := db.Query("SELECT * FROM book")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		err := rows.Scan(&b.ID, &b.Title, &b.ReleaseYear, &b.Summary, &b.Price, &b.AuthorID)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func Find(id int, db *sql.DB) (models.Book, error) {
	var b models.Book

	row := db.QueryRow("SELECT id, title, release_year, summary, price, author_id from book WHERE id = $1", id)
	if err := row.Scan(
		&b.ID,
		&b.Title,
		&b.ReleaseYear,
		&b.Summary,
		&b.Price,
		&b.AuthorID); err != nil {
		if err == sql.ErrNoRows {
			return b, fmt.Errorf("FindBookById %d: no such book", id)
		}
		return b, fmt.Errorf("FindBookById %d: %v", id, err)
	}
	return b, nil
}

func Create(br models.BookRequest, db *sql.DB) (int, error) {
	if br.Title == "" {
		return 0, errors.New("book title is required")
	}

	if br.AuthorID == nil {
		return 0, errors.New("author id is required")
	}

	authorID := *br.AuthorID
	_, err := authorservice.Find(authorID, db)
	if err != nil {
		return 0, fmt.Errorf("AuthorFind error: %v", err)
	}

	var id int
	err = db.QueryRow(
		"INSERT INTO book (title, release_year, summary, price, author_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		br.Title, br.ReleaseYear, br.Summary, br.Price, authorID,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("CreateBook error: %v", err)
	}

	return id, nil
}

func Update(id int, br models.BookRequest, db *sql.DB) error {
	authorID := *br.AuthorID
	_, err := authorservice.Find(authorID, db)

	if err != nil {
		return fmt.Errorf("AuthorFind error: %v", err)
	}

	res, err := db.Exec(
		"UPDATE book SET title = $1, release_year = $2, summary = $3, price = $4, author_id = $5 WHERE id = $6",
		br.Title, br.ReleaseYear, br.Summary, br.Price, authorID, id,
	)

	if err != nil {
		return fmt.Errorf("UpdateBook error: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("RowsAffected error: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no book with id %d found", id)
	}

	return nil
}
