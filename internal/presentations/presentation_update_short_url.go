package presentations

// UpdateShortUrlPayload const
type UpdateShortUrlPayload struct {
	UserID    string `json:"user_id"`
	ShortCode string `json:"short_code"`
}
