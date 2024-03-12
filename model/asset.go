package model

import "time"

type Asset struct {
	ID        int        `json:"id"`
	Name      string     `json:"name" binding:"required"`
	Code      string     `json:"code" binding:"required"`
	InDate    string     `json:"in_date" binding:"required"`
	Source    string     `json:"source" binding:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type AssetDetail struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"name" binding:"required"`
	Code             string                   `json:"code" binding:"required"`
	InDate           string                   `json:"in_date" binding:"required"`
	Source           string                   `json:"source" binding:"required"`
	ControlHistories []AssetMonitoringHistory `json:"monitoring_histories"`
	CreatedAt        time.Time                `json:"created_at"`
	UpdatedAt        time.Time                `json:"updated_at"`
	DeletedAt        *time.Time               `json:"deleted_at"`
}
