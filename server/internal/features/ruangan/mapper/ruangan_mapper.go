package mapper

import (
	"github.com/riyanamanda/ekinerja/internal/features/ruangan/dto"
	"github.com/riyanamanda/ekinerja/internal/features/ruangan/model"
)

func MapToListResponses(list []model.Ruangan) []dto.RuanganResponse {
	responses := make([]dto.RuanganResponse, len(list))
	for i, ruangan := range list {
		responses[i] = *MapToRuanganResponse(&ruangan)
	}
	return responses
}

func MapToRuanganResponse(ruangan *model.Ruangan) *dto.RuanganResponse {
	return &dto.RuanganResponse{
		ID:        ruangan.ID,
		Nama:      ruangan.Nama,
		CreatedAt: ruangan.CreatedAt,
		UpdatedAt: ruangan.UpdatedAt,
	}
}
