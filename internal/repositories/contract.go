package repositories

import (
	"context"
	"database/sql"
	"github.com/aryayunanta-ralali/shorty/internal/entity"
)

type DBTransaction interface {
	ExecTX(ctx context.Context, options *sql.TxOptions, fn func(context.Context, StoreTX) error) error
}

type StoreTX interface {
	Store(ctx context.Context, tableName string, entity interface{}) (id int64, err error)
	Execute(ctx context.Context, query string, values ...interface{}) (affected int64, err error)
	Update(ctx context.Context, tableName string, entity interface{}, whereConditions []WhereCondition) (affected int64, err error)
	StoreBulk(ctx context.Context, tableName string, entity interface{}) (int64, error)
	Upsert(ctx context.Context, tableName string, entity interface{}, onUpdate []WhereCondition) (affected int64, err error)
}

type ShortUrls interface {
	FindBy(ctx context.Context, cri FindShortUrlsCriteria) ([]entity.ShortUrls, error)
	Insert(ctx context.Context, data entity.ShortUrls) error
	Update(ctx context.Context, data entity.ShortUrls) error
	IncrementViewCount(ctx context.Context, id int64) error
}
