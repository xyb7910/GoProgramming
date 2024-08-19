package reflect

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func TestZeroValue(t *testing.T) {
	type UserV1 struct {
		Name string
		age  int
	}
	user := &UserV1{Name: "ypb", age: 19}
	num := &[]int{1}
	testCases := []struct {
		name    string
		input   any
		wantRes map[string]any
		wantErr error
	}{
		{
			name:  "私有字段",
			input: user,
			wantRes: map[string]any{
				"Name": "ypb",
				"age":  0,
			},
			wantErr: nil,
		},
		{
			name:    "多重指针非法字段",
			input:   &num,
			wantRes: nil,
			wantErr: errors.New("传入的值必须是结构体类型"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := IterateFields(tc.input)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestIterateFields(t *testing.T) {
	user := &User{Name: "yxc", Age: 18}
	u := &user
	testCases := []struct {
		name string
		// 输入
		val any
		// 输出
		wantRes map[string]any
		wantErr error
	}{
		{
			name:    "测试nil",
			val:     nil,
			wantErr: errors.New("不能为nil"),
		},
		{
			name: "测试一级指针",
			val:  &User{Name: "yxc", Age: 18},
			wantRes: map[string]any{
				"Name": "yxc",
				"Age":  18,
			},
		},
		{
			name: "测试多级指针",
			val:  u,
			wantRes: map[string]any{
				"Name": "yxc",
				"Age":  18,
			},
		},
		{
			name: "测试结构体",
			val:  User{Name: "yxc", Age: 18},
			wantRes: map[string]any{
				"Name": "yxc",
				"Age":  18,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res, err := IterateFields(tt.val)
			if err != nil {
				return
			}
			assert.Equal(t, res, tt.wantRes)
		})
	}
}

type PrivateUser struct {
	name string
}

func TestSetField(t *testing.T) {
	u1 := &User{Name: "yxc", Age: 18}
	testcases := []struct {
		name    string
		field   string
		entity  any
		newVal  any
		wantErr error
	}{
		{
			name:    "结构体",
			entity:  User{},
			field:   "Name",
			wantErr: errors.New("非法类型"),
		},
		{
			name:    "private 字段",
			entity:  &PrivateUser{},
			field:   "name",
			wantErr: errors.New("字段不可设置"),
		},
		{
			name:    "多级指针",
			entity:  &u1,
			field:   "Name",
			wantErr: errors.New("非法类型"),
		},
		{
			name:    "非法字段",
			entity:  &User{},
			field:   "invalid_field",
			wantErr: errors.New("字段不存在"),
		},
		{
			name: "正常",
			entity: &User{
				Name: "yxc",
				Age:  18,
			},
			field:  "Name",
			newVal: "ypb",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			err := SetField(tt.entity, tt.field, tt.newVal)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
