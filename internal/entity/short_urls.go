// Package entity
package entity

import "time"

const (
	TableNameShortUrls  = `short_urls` // adjust here if wrong value
	EntityNameShortUrls = `ShortUrls`
)

// ShortUrls represent ShortUrls table
type ShortUrls struct {
	ID         int64     `json:"id" db:"id" qb:"id,omitempty"`
	UserID     string    `json:"user_id" db:"user_id,omitempty" qb:"user_id,omitempty"`
	URL        string    `json:"url" db:"url" qb:"url,omitempty"`
	ShortCode  string    `json:"short_code" db:"short_code" qb:"short_code,omitempty"`
	VisitCount int64     `json:"visit_count" db:"visit_count,omitempty" qb:"visit_count,omitempty"`
	CreatedAt  time.Time `json:"created_at" db:"created_at,omitempty" qb:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at,omitempty" qb:"updated_at,omitempty"`
}
