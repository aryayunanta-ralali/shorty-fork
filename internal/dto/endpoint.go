package dto

import "github.com/aryayunanta-ralali/shorty/internal/presentations"

func TransformToGetListResponse(data []string) *presentations.GetListResponses {
	var resp []presentations.GetListResponse

	for _, val := range data {
		currData := presentations.GetListResponse{
			URL: val,
		}

		resp = append(resp, currData)
	}

	return &presentations.GetListResponses{Endpoints: resp}
}
