package pangkat

import (
	"context"
	"strings"
	"time"

	"gorm.io/gorm"
)

type pangkatService struct {
	repo PangkatRepository
}

func NewPangkatService(repo PangkatRepository) PangkatService {
	return &pangkatService{repo: repo}
}

func (p *pangkatService) GetAll(ctx context.Context, page, perPage int) ([]PangkatResponse, int64, error) {
	list, err := p.repo.GetAll(ctx, page, perPage)
	if err != nil {
		return nil, 0, err
	}
	total, err := p.repo.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	return MapToListResponse(list), total, nil
}

func (p *pangkatService) Save(ctx context.Context, request PangkatRequest) error {
	isUnique, err := p.repo.IsPangkatUnique(ctx, request.Nama)
	if err != nil {
		return err
	}
	if !isUnique {
		return gorm.ErrDuplicatedKey
	}
	pangkat := Pangkat{Nama: request.Nama}
	return p.repo.Save(ctx, pangkat)
}

func (p *pangkatService) GetById(ctx context.Context, id int64) (PangkatResponse, error) {
	pangkat, err := p.repo.GetById(ctx, id)
	if err != nil {
		return PangkatResponse{}, err
	}
	return MapToPangkatResponse(pangkat), nil
}

func (p *pangkatService) Update(ctx context.Context, id int64, request PangkatRequest) error {
	existing, err := p.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	if !strings.EqualFold(existing.Nama, request.Nama) {
		isUnique, err := p.repo.IsPangkatUnique(ctx, request.Nama)
		if err != nil {
			return err
		}
		if !isUnique {
			return gorm.ErrDuplicatedKey
		}
	}
	updates := map[string]any{
		"Nama":      request.Nama,
		"UpdatedAt": time.Now(),
	}
	return p.repo.Update(ctx, id, updates)
}

func (p *pangkatService) Delete(ctx context.Context, id int64) error {
	if _, err := p.repo.GetById(ctx, id); err != nil {
		return err
	}
	return p.repo.Delete(ctx, id)
}
