package v0

import (
	"errors"
	"reflect"
	"strings"
	"sync"
	"unicode"
)

type registry struct {
	models sync.Map

	//models map[reflect.Type]*Model // 使得map的key为唯一选择，使用reflect.Type
	//lock   sync.RWMutex            // 使用读写锁，以适应并发
}

func (r *registry) get(val any) (*Model, error) {
	typ := reflect.TypeOf(val)
	m, ok := r.models.Load(typ)
	if !ok {
		var err error
		if m, err = r.ParseModel(val); err != nil {
			return nil, err
		}
		r.models.Store(typ, m)
	}
	return m.(*Model), nil
}

// get 使用 double-check 的方式，以适应并发
//func (r *registry) get(val any) (*Model, error) {
//	r.lock.RLock()
//	typ := reflect.TypeOf(val)
//	m, ok := r.models[typ]
//	r.lock.RUnlock()
//	if ok {
//		return m, nil
//	}
//
//	r.lock.Lock()
//	defer r.lock.Unlock()
//
//	m, ok = r.models[typ]
//	r.lock.RUnlock()
//	if ok {
//		return m, nil
//	}
//
//	var err error
//	if m, err = r.parseModel(typ); err != nil {
//		return nil, err
//	}
//	r.models[typ] = m
//	return m, nil
//}

// get 使用的普通的map，不能适应并发
//func (r *registry) get(val any) (*Model, error) {
//	typ := reflect.TypeOf(val)
//	m, ok := r.models[typ]
//	if !ok {
//		var err error
//		if m, err = r.parseModel(typ); err != nil {
//			return nil, err
//		}
//		r.models[typ] = m
//	}
//	return m, nil
//}

func (r *registry) ParseTag(tag reflect.StructTag) (map[string]string, error) {
	TagField := tag.Get("orm")
	if TagField == "" {
		return map[string]string{}, nil
	}
	res := make(map[string]string, 1)

	// 对获取的tag进行解析
	pairs := strings.Split(TagField, ";")
	for _, pair := range pairs {
		kvp := strings.Split(pair, ":")
		if len(kvp) != 2 {
			return nil, errors.New("invalid tag")
		}
		res[kvp[0]] = kvp[1]
	}
	return res, nil
}

func (r *registry) ParseModel(val any) (*Model, error) {
	typ := reflect.TypeOf(val)
	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return nil, errors.New("model must be a pointer to a struct")
	}
	// 解指针 -- 仅支持一级指针
	if typ.Kind() == reflect.Ptr && typ.Elem().Kind() == reflect.Struct {
		typ = typ.Elem()
	}
	numField := typ.NumField()
	fds := make(map[string]*Field, numField)
	for i := 0; i < numField; i++ {
		fdType := typ.Field(i)
		tags, err := r.ParseTag(fdType.Tag)
		if err != nil {
			return nil, err
		}
		colName := tags[TAGKEYCOLUMN]
		if colName == "" {
			colName = GetName(fdType.Name)
		}
		fds[fdType.Name] = &Field{
			colName: colName,
		}
	}

	var tableName string
	// 如果实现了TableName接口，则使用接口返回的表名
	if tn, ok := val.(TableName); ok {
		tableName = tn.TableName()
	}
	if tableName == "" {
		tableName = GetName(typ.Name())
	}
	return &Model{
		tableName: tableName,
		fieldMaps: fds,
	}, nil
}

func GetName(args string) string {
	var buf []byte
	for i, v := range args {
		if unicode.IsUpper(v) {
			if i != 0 {
				buf = append(buf, '_')
			}
			buf = append(buf, byte(unicode.ToLower(v)))
		} else {
			buf = append(buf, byte(v))
		}
	}
	return string(buf)
}

// Registry 元数据注册中心的抽象
type Registry interface {
	// Get 查找元数据
	Get(val any) (*Model, error)
	//Register 注册模型
	Register(val any, Opts ...ModelOptions) (*Model, error)
}

func (r *registry) Get(val any) (*Model, error) {
	typ := reflect.TypeOf(val)
	m, ok := r.models.Load(typ)
	if ok {
		return m.(*Model), nil
	}
	return nil, errors.New("model not found")
}

func (r *registry) Register(val any, Opts ...ModelOptions) (*Model, error) {
	m, err := r.ParseModel(val)
	if err != nil {
		return nil, err
	}

	for _, opt := range Opts {
		err = opt(m)
		if err != nil {
			return nil, err
		}
	}

	typ := reflect.TypeOf(val)
	r.models.Store(typ, m)
	return m, nil
}
