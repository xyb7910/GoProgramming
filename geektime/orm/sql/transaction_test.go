package sql

import (
	"context"
	"time"
)

func (s *sqlTestSuite) TestTx() {
	t := s.T()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := tx.ExecContext(ctx, "INSERT INTO user (first_name, last_name, age) VALUES (?, ?, ?)", "李", "斯曼", 18)
	if err != nil {
		t.Fatal(err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		t.Fatal(err)
	}
	if affected != 1 {
		t.Fatal(err)
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
	}
}
