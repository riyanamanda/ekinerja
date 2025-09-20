package dto

type AtasanResponse struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	IsActive bool   `json:"is_active"`
}

type AtasanCreateRequest struct {
	Nama string `json:"nama" validate:"required"`
}

type AtasanUpdateRequest struct {
	Nama     string `json:"nama" validate:"required"`
	IsActive *bool  `json:"is_active" validate:"required"`
}
