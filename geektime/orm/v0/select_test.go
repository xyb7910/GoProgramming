package v0

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestModel struct {
	Id        int64
	FirstName string
	Age       int64
	LastName  *sql.NullString
}

func TestSelector_Build(t *testing.T) {
	db, err := NewDB()
	if err != nil {
		t.Fatal(err)
	}
	testCases := []struct {
		name      string
		q         QueryBuilder
		wantQuery *Query
		wantErr   error
	}{
		{
			name: "no form",
			q:    NewSelector[TestModel](db),
			wantQuery: &Query{
				SQL:  "SELECT * FROM test_model;",
				Args: nil,
			},
			wantErr: nil,
		},
		{
			name: "empty form",
			q:    NewSelector[TestModel](db).Form(""),
			wantQuery: &Query{
				SQL:  "SELECT * FROM test_model;",
				Args: nil,
			},
			wantErr: nil,
		},
		{
			name: "with form",
			q:    NewSelector[TestModel](db).Form("test1_model"),
			wantQuery: &Query{
				SQL:  "SELECT * FROM test1_model;",
				Args: nil,
			},
			wantErr: nil,
		},
		{
			name: "with db",
			q:    NewSelector[TestModel](db).Form("test_db.test1_model"),
			wantQuery: &Query{
				SQL:  "SELECT * FROM test_db.test1_model;",
				Args: nil,
			},
			wantErr: nil,
		},
		{
			name: "or",
			q: NewSelector[TestModel](db).Form("test_model").
				Where(C("Id").Eq(12).Or(C("Age").Gt(18))),
			wantQuery: &Query{
				SQL:  "SELECT * FROM test_model WHERE (id = ?) OR (age > ?);",
				Args: []interface{}{12, 18},
			},
			wantErr: nil,
		},
		{
			name: "and",
			q: NewSelector[TestModel](db).Form("test_model").
				Where(C("Id").Eq(12).And(C("Age").Gt(18))),
			wantQuery: &Query{
				SQL:  "SELECT * FROM test_model WHERE (id = ?) AND (age > ?);",
				Args: []interface{}{12, 18},
			},
			wantErr: nil,
		},
		//{
		//	name: "not",
		//	q: NewSelector[TestModel](db).Form("test_model").
		//		Where(Not(C("Age").Gt(18))),
		//	wantQuery: &Query{
		//		SQL:  "SELECT * FROM test_model WHERE NOT (age > ?);",
		//		Args: []interface{}{18},
		//	},
		//	wantErr: errors.New("column not found"),
		//},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			query, err := tc.q.Build()
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantQuery, query)
		})
	}
}
