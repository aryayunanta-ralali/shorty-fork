package presentations

// InsertShortUrlPayload const
type InsertShortUrlPayload struct {
	UserID    string `json:"user_id"`
	URL       string `json:"url"`
	ShortCode string `json:"short_code"`
}
