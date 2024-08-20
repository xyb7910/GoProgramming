package v0

const (
	TAGKEYCOLUMN = "column"
)

type ModelOptions func(model *Model) error

type Model struct {
	tableName string
	fieldMaps map[string]*Field
}

type Field struct {
	colName string
}

// TableName 如果想要修改表名，可以实现这个接口
type TableName interface {
	TableName() string
}
