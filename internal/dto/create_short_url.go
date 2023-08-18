package dto

import (
	"github.com/aryayunanta-ralali/shorty/internal/entity"
	"github.com/aryayunanta-ralali/shorty/internal/presentations"
)

func TransformToInsertEntity(body presentations.InsertShortUrlPayload, id int64) entity.ShortUrls {
	return entity.ShortUrls{
		ID:        id,
		UserID:    body.UserID,
		URL:       body.URL,
		ShortCode: body.ShortCode,
	}
}
