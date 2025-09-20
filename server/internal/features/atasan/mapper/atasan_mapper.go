package mapper

import (
	"github.com/riyanamanda/ekinerja/internal/features/atasan/dto"
	"github.com/riyanamanda/ekinerja/internal/features/atasan/model"
)

func MapToListReponses(list []model.Atasan) []dto.AtasanResponse {
	responses := make([]dto.AtasanResponse, len(list))
	for i, v := range list {
		responses[i] = *MapToAtasanResponse(&v)
	}
	return responses
}

func MapToAtasanResponse(atasan *model.Atasan) *dto.AtasanResponse {
	return &dto.AtasanResponse{
		ID:       atasan.ID,
		Nama:     atasan.Nama,
		IsActive: atasan.IsActive,
	}
}
