package sql

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestMock(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = mockDB.Close()
	}()

	mock.ExpectBegin()

	mockRows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "name")

	mock.ExpectQuery("SELECT .*").WillReturnRows(mockRows)

	mockResult := sqlmock.NewResult(12, 1)
	mock.ExpectExec("UPDATE .*").WillReturnResult(mockResult)
	mock.ExpectCommit()
}
