package pangkat

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/riyanamanda/ekinerja/internal/features/pangkat/dto"
	"github.com/riyanamanda/ekinerja/internal/features/pangkat/mapper"
	"github.com/riyanamanda/ekinerja/internal/features/pangkat/model"
	"gorm.io/gorm"
)

type pangkatService struct {
	repo model.PangkatRepository
}

func NewPangkatService(repo model.PangkatRepository) model.PangkatService {
	return &pangkatService{repo: repo}
}

func (p *pangkatService) GetAll(ctx context.Context, page, size int) ([]dto.PangkatResponse, int64, error) {
	list, err := p.repo.GetAll(ctx, page, size)
	if err != nil {
		return nil, 0, err
	}
	total, err := p.repo.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	return mapper.MapToListResponse(list), total, nil
}

func (p *pangkatService) Save(ctx context.Context, request dto.PangkatRequest) error {
	isUnique, err := p.repo.IsPangkatUnique(ctx, request.Nama)
	if err != nil {
		return err
	}
	if !isUnique {
		return gorm.ErrDuplicatedKey
	}
	pangkat := &model.Pangkat{Nama: request.Nama}
	return p.repo.Save(ctx, pangkat)
}

func (p *pangkatService) GetById(ctx context.Context, id int64) (*dto.PangkatResponse, error) {
	pangkat, err := p.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mapper.MapToPangkatResponse(pangkat), nil
}

func (p *pangkatService) Update(ctx context.Context, id int64, request dto.PangkatRequest) error {
	existing, err := p.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return gorm.ErrRecordNotFound
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
	pangkat, err := p.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	if pangkat == nil {
		return gorm.ErrRecordNotFound
	}
	return p.repo.Delete(ctx, id)
}
