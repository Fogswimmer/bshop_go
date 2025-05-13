package bookservice

import (
	"api/train/mapper"
	"api/train/models/dto"
	"api/train/models/entities"
	"api/train/models/response"
	authorservice "api/train/services/author"
	fileservice "api/train/services/file"
	"database/sql"
	"fmt"
)

func List(db *sql.DB) ([]response.BookResponse, error) {
	query := `
		SELECT b.id, b.title, b.release_year, b.summary, b.price, b.cover,
			a.id, a.firstname, a.lastname, a.birthday
		FROM book b
		LEFT JOIN author a ON b.author_id = a.id
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []response.BookResponse

	for rows.Next() {
		var book entities.Book
		var author entities.Author
		book.Author = author

		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.ReleaseYear,
			&book.Summary,
			&book.Price,
			&book.Cover,
			&book.Author.ID,
			&book.Author.Firstname,
			&book.Author.Lastname,
			&book.Author.Birthday,
		)
		if err != nil {
			return nil, err
		}

		dto := mapper.MapToBookResponse(&book)
		result = append(result, *dto)
	}

	return result, nil
}

func Find(id int, db *sql.DB) (*response.BookResponse, error) {
	var b entities.Book
	b.Author = entities.Author{}

	query := `
		SELECT b.id, b.title, b.release_year, b.summary, b.price, b.cover,
			a.id, a.firstname, a.lastname, a.birthday
		FROM book b
		LEFT JOIN author a ON b.author_id = a.id
		WHERE b.id = $1
	`

	row := db.QueryRow(query, id)

	err := row.Scan(
		&b.ID,
		&b.Title,
		&b.ReleaseYear,
		&b.Summary,
		&b.Price,
		&b.Cover,
		&b.Author.ID,
		&b.Author.Firstname,
		&b.Author.Lastname,
		&b.Author.Birthday,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("FindBookById %d: no such book", id)
		}
		return nil, fmt.Errorf("FindBookById %d: %v", id, err)
	}

	br := mapper.MapToBookResponse(&b)
	return br, nil
}

func Create(br dto.BookDto, db *sql.DB) (int, error) {
	_, err := authorservice.FindEntity(br.AuthorID, db)
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
	_, err := authorservice.FindEntity(br.AuthorID, db)

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
func SaveCover(id int, path string, filename string, db *sql.DB) (*response.BookResponse, error) {
	url := fileservice.GetStaticFileURL(path)
	res, err := db.Exec(
		"UPDATE book SET cover = $1 WHERE id = $2",
		url, id,
	)
	if err != nil {
		return &response.BookResponse{}, fmt.Errorf("SaveCover error: %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return &response.BookResponse{}, fmt.Errorf("RowsAffected error: %v", err)
	}
	if rowsAffected == 0 {
		return &response.BookResponse{}, fmt.Errorf("no book with id %d found", id)
	}
	br, err := Find(id, db)
	if err != nil {
		return &response.BookResponse{}, fmt.Errorf("FindBookById error: %v", err)
	}
	return br, nil
}
