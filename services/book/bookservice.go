package bookservice

import (
	"api/train/helpers"
	"api/train/models/dto"
	"api/train/models/forms/forms"
	"api/train/models/response"
	authorservice "api/train/services/author"
	"database/sql"
	"fmt"
)

func List(db *sql.DB) ([]response.BookResponse, error) {
	query := `
		SELECT b.id, b.title, b.release_year, b.summary, b.price,
			a.id, a.firstname, a.lastname, a.birthday
		FROM book b
		LEFT JOIN author a ON b.author_id = a.id
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []response.BookResponse

	for rows.Next() {
		var b response.BookResponse
		b.Author = &forms.AuthorForm{}

		err := rows.Scan(
			&b.ID,
			&b.Title,
			&b.ReleaseYear,
			&b.Summary,
			&b.Price,
			&b.Author.ID,
			&b.Author.Firstname,
			&b.Author.Lastname,
			&b.Author.Birthday,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}

func Find(id int, db *sql.DB) (response.BookResponse, error) {
	var b response.BookResponse

	row := db.QueryRow("SELECT id, title, release_year, summary, price, author_id from book WHERE id = $1", id)
	if err := row.Scan(
		&b.ID,
		&b.Title,
		&b.ReleaseYear,
		&b.Summary,
		&b.Price,
		&b.Author.ID,
		helpers.GetFullName(b.Author.Firstname, b.Author.Lastname),
	); err != nil {
		if err == sql.ErrNoRows {
			return b, fmt.Errorf("FindBookById %d: no such book", id)
		}
		return b, fmt.Errorf("FindBookById %d: %v", id, err)
	}
	return b, nil
}

func Create(br dto.BookDto, db *sql.DB) (int, error) {
	_, err := authorservice.Find(br.AuthorID, db)
	if err != nil {
		return 0, fmt.Errorf("AuthorFind error: %v", err)
	}

	var id int
	err = db.QueryRow(
		"INSERT INTO book (title, release_year, summary, price, author_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		br.Title, br.ReleaseYear, br.Summary, br.Price, br.AuthorID,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("CreateBook error: %v", err)
	}

	return id, nil
}

func Update(id int, br dto.BookDto, db *sql.DB) error {
	_, err := authorservice.Find(br.AuthorID, db)

	if err != nil {
		return fmt.Errorf("AuthorFind error: %v", err)
	}

	res, err := db.Exec(
		"UPDATE book SET title = $1, release_year = $2, summary = $3, price = $4, author_id = $5 WHERE id = $6",
		br.Title, br.ReleaseYear, br.Summary, br.Price, br.AuthorID, id,
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

func Delete(id int, db *sql.DB) error {
	res, err := db.Exec("DELETE FROM book WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("DeleteBook error: %v", err)
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
