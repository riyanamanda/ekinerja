package dto

import "time"

type RuanganResponse struct {
	ID        int        `json:"id"`
	Nama      string     `json:"nama"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type RuanganRequest struct {
	Nama string `json:"nama" validate:"required,min=3,max=50"`
}
