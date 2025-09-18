package dto

import "time"

type JabatanResponse struct {
	ID        int        `json:"id"`
	Nama      string     `json:"nama"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type JabatanRequest struct {
	Nama string `json:"nama" validate:"required,min=3,max=100"`
}
