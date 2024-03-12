package model

import "time"

type Location struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type LocationUpdateRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
