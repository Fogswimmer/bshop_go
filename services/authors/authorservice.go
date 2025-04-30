package authorservice

import (
	"api/train/models"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

func List(db *sql.DB) ([]models.Author, error) {
	rows, err := db.Query("SELECT * FROM author")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []models.Author
	for rows.Next() {
		var a models.Author
		err := rows.Scan(&a.ID, &a.Firstname, &a.Lastname, &a.Birthday)
		if err != nil {
			return nil, err
		}
		authors = append(authors, a)
	}
	return authors, nil
}

func Find(id int64, db *sql.DB) (models.Author, error) {
	var a models.Author

	row := db.QueryRow("SELECT id, firstname, lastname, birthday from author WHERE id = $1", id)
	if err := row.Scan(&a.ID, &a.Firstname, &a.Lastname, &a.Birthday); err != nil {
		if err == sql.ErrNoRows {
			return a, fmt.Errorf("FindAuthorById %d: no such author", id)
		}
		return a, fmt.Errorf("FindAuthorById %d: %v", id, err)
	}
	return a, nil
}

func Create(ar models.AuthorRequest, db *sql.DB) (int64, error) {
	if ar.Firstname == "" {
		return 0, errors.New("author firstname is required")
	}
	var fmtBD string

	if ar.Birthday != "" {
		var err error
		fmtBD, err = formatBD(ar.Birthday)
		if err != nil {
			return 0, fmt.Errorf("error formatting date: %v", err)
		}
	} else {
		fmtBD = ""
	}

	var id int64
	err := db.QueryRow(
		"INSERT INTO author (firstname, lastname, birthday) VALUES ($1, $2, $3) RETURNING id",
		ar.Firstname, ar.Lastname, fmtBD,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("CreateAuthor error: %v", err)
	}

	return id, nil
}

func formatBD(bd string) (string, error) {
	fmtBD, err := time.Parse("2006-01-02", bd)
	if err != nil {
		return "", err
	}
	return fmtBD.Format("02 Jan 2006"), nil
}
