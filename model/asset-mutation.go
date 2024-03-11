package model

import "time"

type AssetMutation struct {
	ID        int       `json:"id"`
	AssetID   int       `json:"asset_id"`
	Type      string    `json:"type"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
