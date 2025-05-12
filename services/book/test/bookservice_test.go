package bookservice_test

import (
	"api/train/models"
	bookservice "api/train/services/book"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListBooks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock db: %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"id", "title", "release_date", "summary", "price",
		"author_id", "firstname", "lastname", "birthday",
	}).
		AddRow(1, "Tom Sawyer", 1976, "An adventure book", 12.2, 1, "Mark", "Twain", "1835-11-30").
		AddRow(2, "The Red Hat", 1986, "A fairy tale book", 14.2, 2, "Charles", "Perrault", "1628-01-12")

	query := `
		SELECT b.id, b.title, b.release_date, b.summary, b.price,
			a.id, a.firstname, a.lastname, a.birthday
		FROM book b
		LEFT JOIN author a ON b.author_id = a.id
	`

	mock.ExpectQuery(query).WillReturnRows(rows)

	books, err := bookservice.List(db)
	assert.NoError(t, err)
	assert.Len(t, books, 2)

	expected := models.Book{
		ID:          1,
		Title:       "Tom Sawyer",
		ReleaseYear: 1976,
		Summary:     "An adventure book",
		Price:       12.2,
		Author: models.Author{
			ID:        1,
			Firstname: "Mark",
			Lastname:  "Twain",
			Birthday:  "1835-11-30",
		},
	}

	assert.Equal(t, expected, books[0])
}

func TestCreateBookWithMockDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mockAuthorFind(mock, 1)

	br := models.BookRequest{
		Title:       "Tom Sawyer",
		ReleaseYear: 1923,
		Summary:     "An adventure book",
		Price:       12.2,
		AuthorID:    1,
	}

	mock.ExpectQuery("INSERT INTO book \\(title, release_year, summary, price, author_id\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5\\) RETURNING id").
		WithArgs(br.Title, br.ReleaseYear, br.Summary, br.Price, br.AuthorID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	createdBook, err := bookservice.Create(br, db)
	assert.NoError(t, err)
	assert.Equal(t, 1, createdBook)
}

func TestUpdateBookWithcMockDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock db: %v", err)
	}
	defer db.Close()
	bookId := 1

	mockAuthorFind(mock, 1)

	br := models.BookRequest{
		Title:       "Tom Sawyer",
		ReleaseYear: 1923,
		Summary:     "An adventure book",
		Price:       12.2,
		AuthorID:    1,
	}

	mock.ExpectExec("UPDATE book SET title = \\$1, release_year = \\$2, summary = \\$3, price = \\$4, author_id = \\$5 WHERE id = \\$6").
		WithArgs(br.Title, br.ReleaseYear, br.Summary, br.Price, br.AuthorID, bookId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = bookservice.Update(bookId, br, db)
	assert.NoError(t, err)
}

func mockAuthorFind(mock sqlmock.Sqlmock, id int) {
	mock.ExpectQuery("SELECT id, firstname, lastname, birthday FROM author WHERE id = \\$1").
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "firstname", "lastname", "birthday"}).
			AddRow(id, "Mark", "Twain", time.Now()))
}
