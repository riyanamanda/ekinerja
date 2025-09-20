package mapper

import (
	"github.com/riyanamanda/ekinerja/internal/features/jabatan/dto"
	"github.com/riyanamanda/ekinerja/internal/features/jabatan/model"
)

func MapToListResponse(list []model.Jabatan) []dto.JabatanResponse {
	responses := make([]dto.JabatanResponse, len(list))
	for i, jabatan := range list {
		responses[i] = *MapToJabatanResponse(&jabatan)
	}
	return responses
}

func MapToJabatanResponse(jabatan *model.Jabatan) *dto.JabatanResponse {
	return &dto.JabatanResponse{
		ID:        jabatan.ID,
		Nama:      jabatan.Nama,
		CreatedAt: jabatan.CreatedAt,
		UpdatedAt: jabatan.UpdatedAt,
	}
}
