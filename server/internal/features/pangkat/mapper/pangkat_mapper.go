package mapper

import (
	"github.com/riyanamanda/ekinerja/internal/features/pangkat/dto"
	"github.com/riyanamanda/ekinerja/internal/features/pangkat/model"
)

func MapToListResponse(list []model.Pangkat) []dto.PangkatResponse {
	responses := make([]dto.PangkatResponse, len(list))
	for i, pangkat := range list {
		responses[i] = MapToPangkatResponse(pangkat)
	}
	return responses
}

func MapToPangkatResponse(pangkat model.Pangkat) dto.PangkatResponse {
	return dto.PangkatResponse{
		ID:        pangkat.ID,
		Nama:      pangkat.Nama,
		CreatedAt: pangkat.CreatedAt,
		UpdatedAt: pangkat.UpdatedAt,
	}
}
