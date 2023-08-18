package presentations

// GetListPayload const
type GetListPayload struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}

// GetListResponse const
type GetListResponse struct {
	URL string `json:"url"`
}
