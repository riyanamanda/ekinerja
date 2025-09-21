package dto

type RoleResponse struct {
	ID   int64  `json:"id"`
	Nama string `json:"nama"`
}

type RoleRequest struct {
	Nama string `json:"nama" validate:"required"`
}
