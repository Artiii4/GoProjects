package db_test

import (
	"errors"
	"example_mock/internal/db"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func provideNamesSqlData(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows = rows.AddRow(name)
	}
	return rows
}

func TestNew(t *testing.T) {
	mockDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error! '%s' in openning, TestNew", err)
	}
	dbService := db.Service{DB: mockDB}
	result := db.New(mockDB)
	if result != dbService {
		t.Errorf("Right result: %v, result: %v", dbService, result)
	}

}

var tableForTestGetName = []struct {
	names    []string
	rightErr error
}{
	{names: []string{"Fir", "Sec"}, rightErr: nil},
	{names: nil, rightErr: errors.New("ExpectedError")},
}

func TestGetName(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error! '%s' in openning, GetName", err)
	}

	dbService := db.Service{DB: mockDB}

	for num, row := range tableForTestGetName {
		rows := sqlmock.NewRows([]string{"name"})
		for _, name := range row.names {
			rows = rows.AddRow(name)
		}

		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows).WillReturnError(row.rightErr)

		names, err := dbService.GetNames()
		if row.rightErr != nil {
			require.ErrorIs(t, err, row.rightErr, "row: %d, right error:%w, error: %w", num, row.rightErr, err)
			require.Nil(t, names, "row: %d", num)
			continue
		}
		require.Equal(t, row.names, names, "row: %d, right names: %s, names: %s", num, row.rightErr, names)
	}
}

var tableForTestSelectUniqueValues = []struct {
	columnName  string
	tableName   string
	rightResult []string

	rightErr error
}{
	{
		columnName:  "names",
		tableName:   "users",
		rightResult: []string{"Perfecto", "ElectroNic", "Boombl4"},
		rightErr:    nil,
	},
	{
		columnName:  "names",
		tableName:   "users",
		rightResult: nil,
		rightErr:    errors.New("ExpectedError"),
	},
}

func TestSelectUniqueValues(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error! '%s'in opening, SelectUniqueValues", err)
	}
	dbService := db.Service{DB: mockDB}
	for num, test := range tableForTestSelectUniqueValues {
		mock.
			ExpectQuery("SELECT DISTINCT " + test.columnName + " FROM " + test.tableName).
			WillReturnRows(provideNamesSqlData(test.rightResult)).
			WillReturnError(test.rightErr)

		names, err := dbService.SelectUniqueValues(test.columnName, test.tableName)
		if test.rightErr != nil {
			require.ErrorIs(t, err, test.rightErr, "row: %d, right error:%w, error: %w", num, test.rightErr, err)
			require.Nil(t, names, "row:", num)
			continue
		}
		require.Equal(t, test.rightResult, names, "row: %d, right value: %s, value: %s", num, test.rightResult, names)
	}
}
