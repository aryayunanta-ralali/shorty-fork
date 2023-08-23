package presentations

// GetListShortUrlPayload const
type GetListShortUrlPayload struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}

// GetListShortUrlResponse const
type GetListShortUrlResponse struct {
	ID        int64  `json:"id"`
	UserID    string `json:"user_id,omitempty"`
	URL       string `json:"url"`
	ShortCode string `json:"short_code"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
