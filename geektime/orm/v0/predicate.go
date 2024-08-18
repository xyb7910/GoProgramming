package v0

type op string

const (
	OptEq = "="
	OptLt = "<"
	OptGt = ">"

	OptNot = "NOT"
	OptAND = "AND"
	OptOR  = "OR"
)

func (o op) String() string {
	return string(o)
}

type Column struct {
	name string
}

func (c Column) expr() {}

type Value struct {
	val any
}

func (v Value) expr() {}

func (Predicate) expr() {}

func C(name string) Column {
	return Column{name: name}
}

// Eq C("id").Eq(1)
func (c Column) Eq(arg any) Predicate {
	return Predicate{
		left:  c,
		op:    OptEq,
		right: Value{val: arg},
	}
}

func (c Column) Lt(arg any) Predicate {
	return Predicate{
		left:  c,
		op:    OptLt,
		right: Value{val: arg},
	}
}

func (c Column) Gt(arg any) Predicate {
	return Predicate{
		left:  c,
		op:    OptGt,
		right: Value{val: arg},
	}
}

func Not(back Predicate) Predicate {
	return Predicate{
		op:    OptNot,
		right: back,
	}
}

func (front Predicate) And(back Predicate) Predicate {
	return Predicate{
		left:  front,
		op:    OptAND,
		right: back,
	}
}

func (front Predicate) Or(back Predicate) Predicate {
	return Predicate{
		left:  front,
		op:    OptOR,
		right: back,
	}
}

type Predicate struct {
	left  Expression
	op    op
	right Expression
}
