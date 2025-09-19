package dto

import "time"

type BidangResponse struct {
	ID        int        `json:"id"`
	Nama      string     `json:"nama"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type BidangRequest struct {
	Nama string `json:"nama" validate:"required"`
}
