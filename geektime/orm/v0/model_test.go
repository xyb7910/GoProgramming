package v0

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestParseModel(t *testing.T) {
	testCases := []struct {
		name    string
		input   any
		output  *Model
		wantErr error
	}{
		{
			name:  "test model",
			input: &TestModel{},
			output: &Model{
				tableName: "test_model",
				fieldMaps: map[string]*Field{
					"Id": {
						colName: "id",
					},
					"FirstName": {
						colName: "first_name",
					},
					"Age": {
						colName: "age",
					},
					"LastName": {
						colName: "last_name",
					},
				},
			},
		},
	}

	r, _ := NewDB()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := r.r.ParseModel(reflect.TypeOf(tc.input))
			assert.Equal(t, err, tc.wantErr)
			assert.Equal(t, res, tc.output)
		})
	}
}

func TestGetName(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "双单词",
			input:  "FirstName",
			output: "first_name",
		},
		{
			name:   "连续大写字母1",
			input:  "HELLO",
			output: "h_e_l_l_o",
		},
		{
			name:   "连续大写字母2",
			input:  "FirstNAme",
			output: "first_n_a_me",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := GetName(tc.input)
			t.Log(res)
			assert.Equal(t, res, tc.output)
		})
	}
}
