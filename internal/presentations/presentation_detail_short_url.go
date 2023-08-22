package presentations

// DetailShortUrlResponse const
type DetailShortUrlResponse struct {
	ID        int64  `json:"id"`
	UserID    string `json:"user_id,omitempty"`
	URL       string `json:"url"`
	ShortCode string `json:"short_code"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
