package v1

import (
	"reflect"
	"strings"
)

// Selector 用于构造 SELECT 语句
type Selector[T any] struct {
	table string
}

// NewSelector 创建一个新的 Selector
func NewSelector[T any]() *Selector[T] {
	return &Selector[T]{}
}

// Form 指定表名，如果是空字符串，则使用 T 的类型名
func (s *Selector[T]) Form(table string) *Selector[T] {
	s.table = table
	return s
}

// Builder 构造 SELECT 语句
func (s *Selector[T]) Builder() (*Query, error) {
	var sb strings.Builder
	sb.WriteString("SELECT * FROM ")
	if s.table == "" {
		var t T
		sb.WriteByte('`')
		sb.WriteString(reflect.TypeOf(t).Name())
		sb.WriteByte('`')
	} else {
		sb.WriteByte('`')
		sb.WriteString(s.table)
		sb.WriteByte('`')
	}
	sb.WriteString(" ;")
	return &Query{
		SQL: sb.String(),
	}, nil
}
