package v0

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_registry_get(t *testing.T) {
	testCases := []struct {
		name        string
		input       any
		wantedModel *Model
		wantedErr   error
	}{
		{
			name:        "test_model",
			input:       TestModel{},
			wantedModel: nil,
			wantedErr:   errors.New("model must be a pointer to a struct"),
		},
		{
			name:  "pointer",
			input: &TestModel{},
			wantedModel: &Model{
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
		{
			name: "多级指针",
			input: func() **TestModel {
				val := &TestModel{}
				return &val
			}(),
			wantedModel: nil,
			wantedErr:   errors.New("model must be a pointer to a struct"),
		},

		// 测试标签相关
		{
			name: "正常标签",
			input: func() any {
				type TestModel struct {
					Id int64 `orm:"column:id"`
				}
				return &TestModel{}
			}(),
			wantedModel: &Model{
				tableName: "test_model",
				fieldMaps: map[string]*Field{
					"Id": {
						colName: "id",
					},
				},
			},
		},
		{
			name: "无值标签",
			input: func() any {
				type TestModel struct {
					Id int64 `orm:"column:"`
				}
				return &TestModel{}
			}(),
			wantedModel: &Model{
				tableName: "test_model",
				fieldMaps: map[string]*Field{
					"Id": {
						colName: "id",
					},
				},
			},
		},
		{
			name: "非法标签",
			input: func() any {
				type TestModel struct {
					Id int64 `orm:"column"`
				}
				return &TestModel{}
			}(),
			wantedModel: nil,
			wantedErr:   errors.New("invalid tag"),
		},
		{
			name: "忽略标签",
			input: func() any {
				type TestModel struct {
					Id int64 `orm:"abc:abc"`
				}
				return &TestModel{}
			}(),
			wantedModel: &Model{
				tableName: "test_model",
				fieldMaps: map[string]*Field{
					"Id": {
						colName: "id",
					},
				},
			},
		},

		{
			name:  "自定义表名",
			input: &CustomTableName{},
			wantedModel: &Model{
				tableName: "custom_table_name_t",
				fieldMaps: map[string]*Field{
					"Name": {
						colName: "name",
					},
				},
			},
		},
		{
			name:  "table name ptr",
			input: &CustomTableNamePtr{},
			wantedModel: &Model{
				tableName: "custom_table_name_ptr_t",
				fieldMaps: map[string]*Field{
					"Name": {
						colName: "name",
					},
				},
			},
		},
	}

	r := &registry{}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m, err := r.get(tc.input)
			assert.Equal(t, tc.wantedErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantedModel, m)
		})
	}
}

type CustomTableName struct {
	Name string
}

func (c CustomTableName) TableName() string {
	return "custom_table_name_t"
}

type CustomTableNamePtr struct {
	Name string
}

func (c *CustomTableNamePtr) TableName() string {
	return "custom_table_name_ptr_t"
}

type EmptyTableName struct {
	Name string
}

func (c *EmptyTableName) TableName() string {
	return ""
}
