// Package repositories
package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aryayunanta-ralali/shorty/pkg/mariadb"
	"github.com/aryayunanta-ralali/shorty/pkg/tracer"
)

type txDB struct {
	db mariadb.Adapter
}

// NewTxDB create new instance db tx
func NewTxDB(db mariadb.Adapter) DBTransaction {
	return &txDB{db: db}
}

// ExecTX database transaction
func (t *txDB) ExecTX(ctx context.Context, options *sql.TxOptions, fn func(context.Context, StoreTX) error) error {
	newCtx := tracer.SpanStartRepositories(ctx, "txDB.ExecTX")
	defer tracer.SpanFinish(newCtx)

	tx, err := t.db.BeginTx(newCtx, options)
	if err != nil {
		return err
	}

	q := NewStore(tx)

	err = fn(newCtx, q)
	if err != nil {
		tracer.SpanError(newCtx, err)
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}

		return err
	}

	return tx.Commit()
}
