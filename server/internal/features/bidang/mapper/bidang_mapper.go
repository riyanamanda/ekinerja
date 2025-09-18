package mapper

import (
	"github.com/riyanamanda/ekinerja/internal/features/bidang/dto"
	"github.com/riyanamanda/ekinerja/internal/features/bidang/model"
)

func MapToListResponse(list []model.Bidang) []dto.BidangResponse {
	responses := make([]dto.BidangResponse, len(list))
	for i, bidang := range list {
		responses[i] = MapToBidangResponse(bidang)
	}
	return responses
}

func MapToBidangResponse(bidang model.Bidang) dto.BidangResponse {
	return dto.BidangResponse{
		ID:        bidang.ID,
		Nama:      bidang.Nama,
		CreatedAt: bidang.CreatedAt,
		UpdatedAt: bidang.UpdatedAt,
	}
}
