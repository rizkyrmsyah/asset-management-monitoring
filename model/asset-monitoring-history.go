package model

import "time"

type AssetMonitoringHistory struct {
	ID        int       `json:"id"`
	AssetID   int       `json:"asset_id" binding:"required"`
	UserID    int       `json:"user_id"`
	Status    string    `json:"status" binding:"required"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
