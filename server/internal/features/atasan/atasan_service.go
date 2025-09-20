package atasan

import (
	"context"
	"errors"

	"github.com/riyanamanda/ekinerja/internal/features/atasan/dto"
	"github.com/riyanamanda/ekinerja/internal/features/atasan/mapper"
	"github.com/riyanamanda/ekinerja/internal/features/atasan/model"
	"gorm.io/gorm"
)

type atasanService struct {
	repo model.AtasanRepository
}

func NewAtasanService(repository model.AtasanRepository) model.AtasanService {
	return &atasanService{
		repo: repository,
	}
}

func (a *atasanService) GetAll(ctx context.Context, page int, size int) ([]dto.AtasanResponse, int64, error) {
	list, err := a.repo.GetAll(ctx, page, size)
	if err != nil {
		return nil, 0, err
	}
	total, err := a.repo.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	return mapper.MapToListReponses(list), total, nil
}

func (a *atasanService) Create(ctx context.Context, request *dto.AtasanCreateRequest) error {
	isUnique, err := a.repo.IsAtasanUnique(ctx, request.Nama)
	if err != nil {
		return err
	}
	if !isUnique {
		return gorm.ErrDuplicatedKey
	}
	atasan := &model.Atasan{Nama: request.Nama}
	return a.repo.Save(ctx, atasan)
}

func (a *atasanService) GetByID(ctx context.Context, id int) (*dto.AtasanResponse, error) {
	atasan, err := a.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mapper.MapToAtasanResponse(atasan), nil
}

func (a *atasanService) Update(ctx context.Context, id int, request dto.AtasanUpdateRequest) error {
	existing, err := a.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return gorm.ErrRecordNotFound
	}
	updates := map[string]any{
		"nama":      request.Nama,
		"is_active": request.IsActive,
	}
	return a.repo.Update(ctx, id, updates)
}

func (a *atasanService) Delete(ctx context.Context, id int) error {
	atasan, err := a.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if atasan == nil {
		return gorm.ErrRecordNotFound
	}
	return a.repo.Delete(ctx, id)
}
