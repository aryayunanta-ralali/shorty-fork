package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/spf13/cast"

	"github.com/aryayunanta-ralali/shorty/internal/entity"
	"github.com/aryayunanta-ralali/shorty/pkg/mariadb"
	"github.com/aryayunanta-ralali/shorty/pkg/tracer"
)

// MOVE THIS INITIALIZATION TO router.go FILE IF NEEDED
// shortUrlsRepo := repositories.NewShortUrlsRepo(db)

type shortUrlsRepo struct {
	db mariadb.Adapter
	tx DBTransaction
}

func NewShortUrlsRepo(db mariadb.Adapter) ShortUrls {
	return &shortUrlsRepo{
		db: db,
		tx: NewTxDB(db),
	}
}

type FindShortUrlsCriteria struct {
	ID        int
	ShortCode string
	IDs       []int
	Limit     int64
	Offset    int64
}

func (s *shortUrlsRepo) FindBy(ctx context.Context, cri FindShortUrlsCriteria) ([]entity.ShortUrls, error) {
	ctx = tracer.SpanStartRepositories(ctx, "short_urls.FindBy")
	defer tracer.SpanFinish(ctx)

	// check again for the table name, make sure it was correct
	res := []entity.ShortUrls{}
	q :=
		`
		SELECT
  			id,
  			IFNULL(user_id, '') user_id,
  			url,
  			short_code,
  			visit_count,
  			created_at,
  			updated_at
		FROM
			short_urls
		WHERE
			deleted_at IS NULL
		`

	q, vals := s.applyFindCriteria(q, cri)

	err := s.db.Fetch(ctx, &res, q, vals...)

	if err != nil {
		tracer.SpanError(ctx, err)
		return nil, err
	}

	return res, nil
}

func (s *shortUrlsRepo) Insert(ctx context.Context, data entity.ShortUrls) error {
	ctx = tracer.SpanStartRepositories(ctx, "short_urls.Insert")
	defer tracer.SpanFinish(ctx)

	err := s.tx.ExecTX(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead}, func(ctx context.Context, tx StoreTX) error {
		_, err := tx.Store(ctx, entity.TableNameShortUrls, data)

		if err != nil {
			return fmt.Errorf("insert ShortUrls err: %v", err)
		}

		return nil
	})

	if err != nil {
		tracer.SpanError(ctx, err)
		return err
	}

	return nil
}

func (s *shortUrlsRepo) Update(ctx context.Context, data entity.ShortUrls) error {
	newCtx := tracer.SpanStartRepositories(ctx, "short_urls.Update")
	defer tracer.SpanFinish(newCtx)

	err := s.tx.ExecTX(newCtx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead}, func(ctx context.Context, tx StoreTX) error {
		_, err := tx.Update(ctx, entity.TableNameShortUrls, data, []WhereCondition{
			{
				Column:   "id",
				Operator: "=",
				Value:    cast.ToString(data.ID),
			},
		})

		if err != nil {
			return fmt.Errorf("update ShortUrls err: %v", err)
		}

		return nil
	})

	if err != nil {
		tracer.SpanError(ctx, err)
		return err
	}

	return nil
}

func (s *shortUrlsRepo) IncrementViewCount(ctx context.Context, id int64) error {
	newCtx := tracer.SpanStartRepositories(ctx, "short_urls.IncrementViewCount")
	defer tracer.SpanFinish(newCtx)

	q := "UPDATE short_urls SET visit_count = visit_count + 1 WHERE id = ?"

	err := s.tx.ExecTX(newCtx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead}, func(ctx context.Context, tx StoreTX) error {
		_, err := tx.Execute(ctx, q, id)

		if err != nil {
			return fmt.Errorf("increment view count ShortUrls err: %v", err)
		}

		return nil
	})

	if err != nil {
		tracer.SpanError(ctx, err)
		return err
	}

	return nil
}

func (s *shortUrlsRepo) applyFindCriteria(q string, cri FindShortUrlsCriteria) (qry string, vals []interface{}) {
	additionalQry := ""
	additionalQry2 := ""

	vals = make([]interface{}, 0)

	if cri.ID != 0 {
		additionalQry += ` AND id = ?`
		vals = append(vals, cri.ID)
	}

	if cri.ShortCode != "" {
		additionalQry += ` AND short_code = ?`
		vals = append(vals, cri.ShortCode)
	}

	if len(cri.IDs) > 0 {
		additionalQry += ` AND id IN (?` + strings.Repeat(",?", len(cri.IDs)-1) + `)`

		for _, v := range cri.IDs {
			vals = append(vals, v)
		}
	}

	q = fmt.Sprintf("%s %s", q, additionalQry)

	if cri.Limit != 0 {
		additionalQry2 = " LIMIT ?"
		vals = append(vals, cri.Limit)

		if cri.Offset != 0 {
			additionalQry2 = fmt.Sprintf("%s %s", additionalQry2, " OFFSET ?")
			vals = append(vals, cri.Offset)
		}
	}

	q = fmt.Sprintf("%s %s", q, additionalQry2)

	return q, vals
}
