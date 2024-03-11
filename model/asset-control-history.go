package model

import "time"

type AssetControlHistory struct {
	ID        int       `json:"id"`
	AssetID   int       `json:"asset_id"`
	UserID    int       `json:"user_id"`
	Status    string    `json:"status"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
