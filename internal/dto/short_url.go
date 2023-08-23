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

func TransformToGetListShortUrlResponse(data []entity.ShortUrls) []presentations.GetListShortUrlResponse {
	var resp []presentations.GetListShortUrlResponse

	for _, val := range data {
		currData := presentations.GetListShortUrlResponse{
			ID:        val.ID,
			UserID:    val.UserID,
			URL:       val.URL,
			ShortCode: val.ShortCode,
			CreatedAt: val.CreatedAt.Format(consts.LayoutDateTimeFormat),
			UpdatedAt: val.UpdatedAt.Format(consts.LayoutDateTimeFormat),
		}

		resp = append(resp, currData)
	}

	return resp
}

func TransformToDetailShortUrlResponse(data entity.ShortUrls) presentations.DetailShortUrlResponse {
	return presentations.DetailShortUrlResponse{
		ID:        data.ID,
		UserID:    data.UserID,
		URL:       data.URL,
		ShortCode: data.ShortCode,
		CreatedAt: data.CreatedAt.Format(consts.LayoutDateTimeFormat),
		UpdatedAt: data.UpdatedAt.Format(consts.LayoutDateTimeFormat),
	}
}

func TransformToUpdateEntity(body entity.ShortUrls, payload presentations.UpdateShortUrlPayload) entity.ShortUrls {
	// Skip updating visit count field (omitempty)
	body.VisitCount = 0

	// Update
	body.ShortCode = payload.ShortCode
	return body
}

func TransformToGetStatShortUrlResponse(data entity.ShortUrls) presentations.GetStatShortUrlResponse {
	return presentations.GetStatShortUrlResponse{
		ID:         data.ID,
		VisitCount: data.VisitCount,
	}
}
