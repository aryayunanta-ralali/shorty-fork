package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/aryayunanta-ralali/shorty/pkg/mariadb"
	"github.com/aryayunanta-ralali/shorty/pkg/util"
)

type storer struct {
	dbTx mariadb.Transaction
}

//  NewStore creates a new instance of database transaction.
func NewStore(dbTx mariadb.Transaction) *storer {
	return &storer{dbTx: dbTx}
}

type WhereCondition struct {
	EscapePreparedStatement bool
	Column                  string
	Operator                string
	Value                   string
}

func (r *storer) Store(ctx context.Context, tableName string, entity interface{}) (id int64, err error) {
	cols, vals, err := util.ToColumnsValues(entity, "db")
	if err != nil {
		return 0, err
	}

	q := `INSERT INTO %s (%s) VALUES(%s)`

	q = fmt.Sprintf(q, tableName, strings.Join(cols, ","), "?"+strings.Repeat(", ?", len(cols)-1))

	res, err := r.dbTx.ExecContext(ctx, q, vals...)
	if err != nil {
		return 0, err
	}

	id, err = res.LastInsertId()
	return id, err
}

func (r *storer) Update(ctx context.Context, tableName string, entity interface{}, conditions []WhereCondition) (affected int64, err error) {
	cols, vals, err := util.ToColumnsValues(entity, "db")
	if err != nil {
		return 0, err
	}

	q := `UPDATE %s SET %s`
	q = fmt.Sprintf(q, tableName, util.StringJoin(cols, "=?, ", "=?"))

	if len(conditions) > 0 {
		q = fmt.Sprintf(`%s WHERE 1=1`, q)
	}

	for _, v := range conditions {
		if v.EscapePreparedStatement {
			q = fmt.Sprintf(`%s AND %s %s %s`, q, v.Column, v.Operator, v.Value)
			continue
		}

		q = fmt.Sprintf(`%s AND %s %s %s`, q, v.Column, v.Operator, "?")
		vals = append(vals, v.Value)
	}

	result, err := r.dbTx.ExecContext(ctx, q, vals...)
	if err != nil {
		return affected, err
	}

	affected, _ = result.RowsAffected()
	return affected, err
}

// StoreBulk inserts data in bulk
func (s *storer) StoreBulk(ctx context.Context, tableName string, param interface{}) (int64, error) {
	var (
		affected int64
		err      error
	)

	cols, vals, patterns, err := util.SliceStructToBulkInsert(param, `db`)
	if err != nil {
		return affected, err
	}

	q := `INSERT INTO %s (%s) VALUES %s;`
	q = fmt.Sprintf(q, tableName, strings.Join(cols, ","), strings.Join(patterns, `,`))
	result, err := s.dbTx.ExecContext(ctx, q, vals...)
	if err != nil {
		return affected, err
	}

	affected, _ = result.RowsAffected()

	return affected, err
}

// Upsert inserts on new key and updates on duplicate key.
// The "operator" in "onUpdate" param will not be used in this function.
func (r *storer) Upsert(ctx context.Context, tableName string, entity interface{}, onUpdate []WhereCondition) (affected int64, err error) {
	cols, vals, err := util.ToColumnsValues(entity, "db")
	if err != nil {
		return 0, err
	}

	q := `INSERT INTO %s (%s) VALUES(%s)`
	q = fmt.Sprintf(q, tableName, strings.Join(cols, ","), "?"+strings.Repeat(", ?", len(cols)-1))

	if len(onUpdate) == 0 {
		return 0, fmt.Errorf("no on duplicate key update statement was given during upsert")
	}

	onUpdateColVal := []string{}

	for _, v := range onUpdate {
		if v.EscapePreparedStatement {
			statement := fmt.Sprintf(`%s = %s`, v.Column, v.Value)
			onUpdateColVal = append(onUpdateColVal, statement)
			continue
		}

		statement := fmt.Sprintf(`%s = %s`, v.Column, "?")
		vals = append(vals, v.Value)
		onUpdateColVal = append(onUpdateColVal, statement)
	}

	q = fmt.Sprintf(`%s ON DUPLICATE KEY UPDATE %s`, q, strings.Join(onUpdateColVal, ", "))

	res, err := r.dbTx.ExecContext(ctx, q, vals...)
	if err != nil {
		return affected, err
	}

	affected, _ = res.RowsAffected()
	return affected, err
}
