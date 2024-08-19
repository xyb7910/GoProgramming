package reflect

import (
	"LearingGo/orm/reflect/types"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCallMethod(t *testing.T) {
	testcases := []struct {
		name    string
		input   any
		wantRes map[string]*FuncInfo
		wantErr error
	}{
		{
			name:  "normal struct",
			input: types.User{},
			wantRes: map[string]*FuncInfo{
				"GetAge": {
					Name:   "GetAge",
					In:     []reflect.Type{reflect.TypeOf(types.User{})},
					Out:    []reflect.Type{reflect.TypeOf(0)},
					Result: []any{0},
				},
			},
		},
		{
			// 指针
			name:  "pointer",
			input: &types.User{},
			wantRes: map[string]*FuncInfo{
				"GetAge": {
					Name:   "GetAge",
					In:     []reflect.Type{reflect.TypeOf(&types.User{})},
					Out:    []reflect.Type{reflect.TypeOf(0)},
					Result: []any{0},
				},
				"ChangeName": {
					Name:   "ChangeName",
					In:     []reflect.Type{reflect.TypeOf(&types.User{}), reflect.TypeOf("")},
					Out:    []reflect.Type{},
					Result: []any{},
				},
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			res, err := CallMethod(tt.input)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
