package v0

type DB struct {
	r *registry
}

type DBOption func(*DB)

func NewDB(Opt ...DBOption) (*DB, error) {
	db := &DB{r: &registry{}}
	for _, o := range Opt {
		o(db)
	}
	return db, nil
}

func MustNewDB(Opt ...DBOption) *DB {
	db, err := NewDB(Opt...)
	if err != nil {
		panic(err)
	}
	return db
}
