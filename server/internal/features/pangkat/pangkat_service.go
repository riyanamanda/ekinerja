package pangkat

import (
	"context"
	"time"

	"github.com/riyanamanda/ekinerja/internal/features/pangkat/dto"
)

type pangkatService struct {
	repo PangkatRepository
}

func NewPangkatService(repo PangkatRepository) PangkatService {
	return &pangkatService{repo: repo}
}

func (p *pangkatService) GetAll(ctx context.Context, page, perPage int) ([]dto.PangkatResponse, int64, error) {
	pangkatList, err := p.repo.GetAll(ctx, page, perPage)
	if err != nil {
		return nil, 0, err
	}
	total, err := p.repo.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	return mapToResponses(pangkatList), total, nil
}

func (p *pangkatService) Save(ctx context.Context, request dto.PangkatRequest) error {
	pangkat := Pangkat{Nama: request.Nama}
	return p.repo.Save(ctx, pangkat)
}

func (p *pangkatService) GetById(ctx context.Context, id int64) (dto.PangkatResponse, error) {
	pangkat, err := p.repo.GetById(ctx, id)
	if err != nil {
		return dto.PangkatResponse{}, err
	}
	return mapToResponse(pangkat), nil
}

func (p *pangkatService) Update(ctx context.Context, id int64, request dto.PangkatRequest) error {
	updates := map[string]any{
		"Nama":      request.Nama,
		"UpdatedAt": time.Now(),
	}
	return p.repo.Update(ctx, id, updates)
}

func (p *pangkatService) Delete(ctx context.Context, id int64) error {
	return p.repo.Delete(ctx, id)
}

func mapToResponses(list []Pangkat) []dto.PangkatResponse {
	responses := make([]dto.PangkatResponse, len(list))
	for i, pangkat := range list {
		responses[i] = mapToResponse(pangkat)
	}
	return responses
}

func mapToResponse(pangkat Pangkat) dto.PangkatResponse {
	return dto.PangkatResponse{
		ID:        pangkat.ID,
		Nama:      pangkat.Nama,
		CreatedAt: pangkat.CreatedAt,
		UpdatedAt: pangkat.UpdatedAt,
	}
}
