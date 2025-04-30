package authorservice_test

import (
	"api/train/models"
	authorservice "api/train/services/author"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

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
	assert.Equal(t, int64(1), createdAuthor)
}
