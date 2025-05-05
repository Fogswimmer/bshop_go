package authorservice_test

import (
	"api/train/models"
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

	rows := sqlmock.NewRows([]string{"id", "firstname", "lastname", "birthday"}).
		AddRow(1, "Jane", "Austen", "1775-12-16").
		AddRow(2, "George", "Orwell", "1903-06-25")

	mock.ExpectQuery("SELECT \\* FROM author").WillReturnRows(rows)

	authors, err := authorservice.List(db)
	assert.NoError(t, err)
	assert.Len(t, authors, 2)

	expected := models.Author{
		ID:        1,
		Firstname: "Jane",
		Lastname:  "Austen",
		Birthday:  "1775-12-16",
	}

	assert.Equal(t, expected, authors[0])
}

func TestCreateAuthorWithMockDB(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock db: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("INSERT INTO author").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	ar := models.AuthorRequest{
		Firstname: "John",
		Lastname:  "Doe",
	}

	createdAuthor, err := authorservice.Create(ar, db)
	assert.NoError(t, err)
	assert.Equal(t, int(1), createdAuthor)
}
