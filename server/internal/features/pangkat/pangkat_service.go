package pangkat

import (
	"context"
	"database/sql"
	"time"

	"github.com/riyanamanda/ekinerja/internal/features/pangkat/dto"
)

type pangkatService struct {
	repo PangkatRepository
}

func NewPangkatService(repo PangkatRepository) PangkatService {
	return &pangkatService{
		repo: repo,
	}
}

func (p *pangkatService) GetAll(ctx context.Context) ([]dto.PangkatResponse, error) {
	pangkatList, err := p.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	responses := make([]dto.PangkatResponse, len(pangkatList))
	for i, pangkat := range pangkatList {
		responses[i] = dto.PangkatResponse{
			ID:        pangkat.ID,
			Nama:      pangkat.Nama,
			CreatedAt: pangkat.CreatedAt,
			UpdatedAt: pangkat.UpdatedAt,
		}
	}
	return responses, nil
}

func (p *pangkatService) Save(ctx context.Context, request dto.PangkatRequest) error {
	pangkat := Pangkat{
		Nama: request.Nama,
	}

	return p.repo.Save(ctx, &pangkat)
}

func (p *pangkatService) GetById(ctx context.Context, id int64) (*dto.PangkatResponse, error) {
	pangkat, err := p.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	if pangkat == nil {
		return nil, nil
	}
	response := &dto.PangkatResponse{
		ID:        pangkat.ID,
		Nama:      pangkat.Nama,
		CreatedAt: pangkat.CreatedAt,
		UpdatedAt: pangkat.UpdatedAt,
	}
	return response, nil
}

func (p *pangkatService) Update(ctx context.Context, id int64, request dto.PangkatRequest) error {
	_, err := p.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	return p.repo.Update(ctx, id, map[string]any{
		"Nama":      request.Nama,
		"UpdatedAt": sql.NullTime{Time: time.Now(), Valid: true},
	})
}

func (p *pangkatService) Delete(ctx context.Context, id int64) error {
	return p.repo.Delete(ctx, id)
}
