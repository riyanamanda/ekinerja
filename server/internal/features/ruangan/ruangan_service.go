package ruangan

import (
	"context"
	"errors"
	"strings"

	"github.com/riyanamanda/ekinerja/internal/features/ruangan/dto"
	"github.com/riyanamanda/ekinerja/internal/features/ruangan/mapper"
	"github.com/riyanamanda/ekinerja/internal/features/ruangan/model"
	"gorm.io/gorm"
)

type ruanganService struct {
	repo model.RuanganRepository
}

func NewRuanganService(repo model.RuanganRepository) model.RuanganService {
	return &ruanganService{repo: repo}
}

func (r *ruanganService) GetAll(ctx context.Context, page int, size int) ([]dto.RuanganResponse, int64, error) {
	list, err := r.repo.GetAll(ctx, page, size)
	if err != nil {
		return nil, 0, err
	}
	count, err := r.repo.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	return mapper.MapToListResponses(list), count, nil
}

func (r *ruanganService) Create(ctx context.Context, request dto.RuanganRequest) error {
	isUnique, err := r.repo.IsRuanganUnique(ctx, request.Nama)
	if err != nil {
		return err
	}
	if !isUnique {
		return gorm.ErrDuplicatedKey
	}
	ruangan := &model.Ruangan{Nama: request.Nama}
	return r.repo.Save(ctx, ruangan)
}

func (r *ruanganService) GetById(ctx context.Context, id int) (*dto.RuanganResponse, error) {
	ruangan, err := r.repo.GetById(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return mapper.MapToRuanganResponse(ruangan), nil
}

func (r *ruanganService) Update(ctx context.Context, id int, request dto.RuanganRequest) error {
	existing, err := r.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return gorm.ErrRecordNotFound
	}
	if !strings.EqualFold(existing.Nama, request.Nama) {
		isUnique, err := r.repo.IsRuanganUnique(ctx, request.Nama)
		if err != nil {
			return err
		}
		if !isUnique {
			return gorm.ErrDuplicatedKey
		}
	}
	updates := map[string]any{
		"nama": request.Nama,
	}
	return r.repo.Update(ctx, id, updates)
}

func (r *ruanganService) Delete(ctx context.Context, id int) error {
	if _, err := r.repo.GetById(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.repo.Delete(ctx, id)
}
