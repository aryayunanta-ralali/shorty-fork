package dto

import (
	"github.com/aryayunanta-ralali/shorty/internal/consts"
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

func TransformToDetailShortUrlResponse(data entity.ShortUrls) presentations.DetailShortUrlResponse {
	return presentations.DetailShortUrlResponse{
		ID:         data.ID,
		UserID:     data.UserID,
		URL:        data.URL,
		ShortCode:  data.ShortCode,
		VisitCount: data.VisitCount,
		CreatedAt:  data.CreatedAt.Format(consts.LayoutDateTimeFormat),
		UpdatedAt:  data.UpdatedAt.Format(consts.LayoutDateTimeFormat),
	}
}

func TransformToUpdateEntity(body entity.ShortUrls, payload presentations.UpdateShortUrlPayload) entity.ShortUrls {
	// Skip updating visit count field (omitempty)
	body.VisitCount = 0

	// Update
	body.ShortCode = payload.ShortCode
	return body
}
