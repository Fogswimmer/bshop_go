package authorservice_test

import (
	"api/train/helpers"
	"api/train/models/dto"
	authorservice "api/train/services/author"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestListAuthors(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock db: %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"id", "firstname", "lastname", "birthday",
		"id", "title", "release_year", "summary", "price"}).
		AddRow(1, "Jane", "Austen", "1775-12-16",
			1, "Pride and Prejudice", 1813, "Pride and Prejudice is the second novel by English author Jane Austen, published in 1813", 23.0).
		AddRow(2, "George", "Orwell", "1903-06-25",
			nil, nil, nil, nil, nil)

	query := `
		SELECT a.id, a.firstname, a.lastname, a.birthday,
			b.id, b.title, b.release_year, b.summary, b.price
		FROM author a
		LEFT JOIN book b ON a.id = b.author_id
	`

	mock.ExpectQuery(query).WillReturnRows(rows)

	authors, _ := authorservice.List(db)

	assert.Equal(t, 2, len(authors))
	assert.Equal(t, "Jane", authors[0].Firstname)
	assert.Equal(t, 1, len(authors[0].Books))
	assert.Equal(t, "Pride and Prejudice", authors[0].Books[0].Title)
}

func TestFinAuthorWithMockDB(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock db: %v", err)
	}
	defer db.Close()

	authorId := 1

	rows := sqlmock.NewRows([]string{
		"id", "firstname", "lastname", "birthday",
		"id", "title", "release_year", "summary", "price"}).
		AddRow(1, "Jane", "Austen", "1775-12-16",
			1, "Pride and Prejudice", 1813, "Pride and Prejudice is the second novel by English author Jane Austen, published in 1813", 23.0).
		AddRow(1, "Jane", "Austen", "1775-12-16",
			2, "Sense and Sensibility", 1813, "Sense and Sensibility is a novel by English author Jane Austen, published in 1813", 23.0)

	query := `
		SELECT a.id, a.firstname, a.lastname, a.birthday,
			b.id, b.title, b.release_year, b.summary, b.price
		FROM author a
		LEFT JOIN book b ON a.id = b.author_id
		WHERE a.id = \$1
	`

	mock.ExpectQuery(query).WithArgs(authorId).WillReturnRows(rows)

	author, _ := authorservice.Find(authorId, db)

	assert.Equal(t, "Jane", author.Firstname)
	assert.Equal(t, 2, len(author.Books))
	assert.Equal(t, "Pride and Prejudice", author.Books[0].Title)
}

func TestCreateAuthorWithMockDB(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock db: %v", err)
	}
	defer db.Close()
	mock.ExpectQuery("INSERT INTO author").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	dto := dto.AuthorDto{
		Firstname: "John",
		Lastname:  "Doe",
	}

	createdAuthor, err := authorservice.Create(dto, db)
	assert.NoError(t, err)
	assert.Equal(t, int(1), createdAuthor)
}

func TestUpdateAuthorWithMockDB(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock db: %v", err)
	}
	defer db.Close()
	authorId := 1
	dto := dto.AuthorDto{
		Firstname: "John",
		Lastname:  "Doe",
		Birthday:  "2000-01-01",
	}

	fmtBD, _ := helpers.FormatBD(dto.Birthday)

	mock.ExpectExec("UPDATE author SET firstname = \\$1, lastname = \\$2, birthday = \\$3 WHERE id = \\$4").
		WithArgs(dto.Firstname, dto.Lastname, fmtBD, authorId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = authorservice.Update(authorId, dto, db)
	assert.NoError(t, err)
}

func TestDeleteCascadeAuthorWithMockDB(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock db: %v", err)
	}
	defer db.Close()
	authorId := 1

	mock.ExpectExec("DELETE FROM book WHERE author_id = \\$1").
		WithArgs(authorId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectExec("DELETE FROM author WHERE id = \\$1").
		WithArgs(authorId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = authorservice.DeleteCascade(authorId, db)
	assert.NoError(t, err)
}
