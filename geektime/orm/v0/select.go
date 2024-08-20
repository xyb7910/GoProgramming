package v0

import (
	"errors"
	"reflect"
	"strings"
	"unicode"
)

var ErrUnsupportedExpression = errors.New("unsupported expression")

type Selector[T any] struct {
	sb        strings.Builder
	tableName string
	where     []Predicate
	args      []any
	model     *Model

	db *DB
}

func NewSelector[T any](db *DB) *Selector[T] {
	return &Selector[T]{db: db}
}

func (s *Selector[T]) Where(ps ...Predicate) *Selector[T] {
	s.where = ps
	return s
}

func (s *Selector[T]) Form(tableName string) *Selector[T] {
	s.tableName = tableName
	return s
}

func (s *Selector[T]) Build() (*Query, error) {
	var (
		t   T
		err error
	)
	s.model, err = s.db.r.get(&t)
	if err != nil {
		return nil, err
	}
	//var sb strings.Builder
	s.sb.WriteString("SELECT * FROM ")
	if s.tableName != "" {
		s.sb.WriteString(s.tableName)
	} else {
		s.sb.WriteString(s.GetTableName())
	}

	// 加入where
	if len(s.where) > 0 {
		s.sb.WriteString(" WHERE ")
		p := s.where[0]
		for i := 1; i < len(s.where); i++ {
			p = p.And(s.where[i])
		}
		err = s.buildExpression(p)
		if err != nil {
			return nil, err
		}
	}
	s.sb.WriteString(";")
	return &Query{
		SQL:  s.sb.String(),
		Args: s.args,
	}, nil
}

func (s *Selector[T]) buildExpression(p Expression) error {
	if p == nil {
		return nil
	}
	switch expr := p.(type) {
	case Column:
		//s.sb.WriteString(expr.name)
		fd, ok := s.model.fieldMaps[expr.name]
		if !ok {
			return errors.New("column not found")
		}
		s.sb.WriteString(fd.colName)
	case Value:
		s.sb.WriteString("?")
		if s.args == nil {
			s.args = make([]any, 0, 8)
		}
		s.args = append(s.args, expr.val)
	case Predicate:
		_, ok := expr.left.(Predicate)
		if ok {
			s.sb.WriteString("(")
		}
		if err := s.buildExpression(expr.left); err != nil {
			return err
		}
		if ok {
			s.sb.WriteString(")")
		}

		s.sb.WriteString(" ")
		s.sb.WriteString(expr.op.String())
		s.sb.WriteString(" ")

		_, ok = expr.right.(Predicate)
		if ok {
			s.sb.WriteString("(")
		}
		if err := s.buildExpression(expr.right); err != nil {
			return err
		}
		if ok {
			s.sb.WriteString(")")
		}
	default:
		return ErrUnsupportedExpression
	}
	return nil
}

func (s *Selector[T]) GetTableName() string {
	// 输出结果为: table_name
	var t T
	typ := reflect.TypeOf(t)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	name := typ.Name()
	var sb strings.Builder
	for i, r := range name {
		if i > 0 && unicode.IsUpper(r) {
			sb.WriteRune('_')
		}
		sb.WriteRune(unicode.ToLower(r))
	}
	return sb.String()
}
