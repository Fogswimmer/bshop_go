package authorservice

import (
	"api/train/models/dto"
	"api/train/models/entities"
	"api/train/models/forms/forms"
	"api/train/models/response"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

func List(db *sql.DB) ([]response.AuthorResponse, error) {
	query := `
		SELECT a.id, a.firstname, a.lastname, a.birthday,
			b.id, b.title, b.release_year, b.summary, b.price
		FROM author a
		LEFT JOIN book b ON a.id = b.author_id
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	authorMap := make(map[int]*response.AuthorResponse)

	for rows.Next() {
		var (
			authorID        int
			firstname       string
			lastname        string
			birthday        string
			bookID          sql.NullInt64
			bookTitle       sql.NullString
			bookReleaseYear sql.NullInt64
			bookSummary     sql.NullString
			bookPrice       sql.NullFloat64
		)

		err := rows.Scan(&authorID, &firstname, &lastname, &birthday,
			&bookID, &bookTitle, &bookReleaseYear, &bookSummary, &bookPrice)
		if err != nil {
			return nil, err
		}

		author, exists := authorMap[authorID]
		if !exists {
			author = &response.AuthorResponse{
				ID:        authorID,
				Firstname: firstname,
				Lastname:  lastname,
				Birthday:  birthday,
				Books:     []forms.BookForm{},
			}
			authorMap[authorID] = author
		}

		if bookID.Valid {
			book := forms.BookForm{
				ID:          int(bookID.Int64),
				Title:       bookTitle.String,
				ReleaseYear: int(bookReleaseYear.Int64),
				Summary:     bookSummary.String,
				Price:       bookPrice.Float64,
			}
			author.Books = append(author.Books, book)
		}
	}

	var authors []response.AuthorResponse
	for _, author := range authorMap {
		authors = append(authors, *author)
	}
	return authors, nil
}

func Find(id int, db *sql.DB) (entities.Author, error) {
	var a entities.Author

	row := db.QueryRow("SELECT id, firstname, lastname, birthday FROM author WHERE id = $1", id)
	if err := row.Scan(&a.ID, &a.Firstname, &a.Lastname, &a.Birthday); err != nil {
		if err == sql.ErrNoRows {
			return a, fmt.Errorf("FindAuthorById %d: no such author", id)
		}
		return a, fmt.Errorf("FindAuthorById %d: %v", id, err)
	}
	return a, nil
}

func Create(dto dto.AuthorDto, db *sql.DB) (int, error) {
	if dto.Firstname == "" {
		return 0, errors.New("author firstname is required")
	}
	var fmtBD string

	if dto.Birthday != "" {
		var err error
		fmtBD, err = FormatBD(dto.Birthday)
		if err != nil {
			return 0, fmt.Errorf("error formatting date: %v", err)
		}
	} else {
		fmtBD = ""
	}

	var id int
	err := db.QueryRow(
		"INSERT INTO author (firstname, lastname, birthday) VALUES ($1, $2, $3) RETURNING id",
		dto.Firstname, dto.Lastname, fmtBD,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("CreateAuthor error: %v", err)
	}

	return id, nil
}

func Update(id int, dto dto.AuthorDto, db *sql.DB) error {
	if dto.Firstname == "" {
		return errors.New("author firstname is required")
	}

	var fmtBD string

	if dto.Birthday != "" {
		var err error
		fmtBD, err = FormatBD(dto.Birthday)
		if err != nil {
			return fmt.Errorf("error formatting date: %v", err)
		}
	} else {
		fmtBD = ""
	}

	res, err := db.Exec(
		"UPDATE author SET firstname = $1, lastname = $2, birthday = $3 WHERE id = $4",
		dto.Firstname, dto.Lastname, fmtBD, id,
	)

	if err != nil {
		return fmt.Errorf("UpdateAuthor error: %v", err)
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

func DeleteCascade(id int, db *sql.DB) error {
	_, err := db.Exec("DELETE FROM book WHERE author_id = $1", id)
	if err != nil {
		return fmt.Errorf("DeleteAuthor error: %v", err)
	}

	_, err = db.Exec("DELETE FROM author WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("DeleteAuthor error: %v", err)
	}

	return nil
}

func FormatBD(bd string) (string, error) {
	fmtBD, err := time.Parse("2006-01-02", bd)
	if err != nil {
		return "", err
	}
	return fmtBD.Format("02 Jan 2006"), nil
}
