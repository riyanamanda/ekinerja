package pangkat

import "time"

type PangkatResponse struct {
	ID        int64      `json:"id"`
	Nama      string     `json:"nama"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type PangkatRequest struct {
	Nama string `json:"nama" validate:"required,min=3,max=100"`
}
