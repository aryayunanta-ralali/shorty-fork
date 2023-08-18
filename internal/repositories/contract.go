package repositories

import (
	"context"
	"github.com/aryayunanta-ralali/shorty/internal/entity"
)

type ShortUrls interface {
	FindBy(ctx context.Context, cri FindShortUrlsCriteria) ([]entity.ShortUrls, error)
	Insert(ctx context.Context, data entity.ShortUrls) error
	Update(ctx context.Context, data entity.ShortUrls) error
}
