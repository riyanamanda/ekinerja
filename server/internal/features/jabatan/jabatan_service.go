package jabatan

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/riyanamanda/ekinerja/internal/features/jabatan/dto"
	"github.com/riyanamanda/ekinerja/internal/features/jabatan/mapper"
	"github.com/riyanamanda/ekinerja/internal/features/jabatan/model"
	"gorm.io/gorm"
)

type jabatanService struct {
	repo model.JabatanRepository
}

func NewJabatanService(repo model.JabatanRepository) model.JabatanService {
	return &jabatanService{repo: repo}
}

func (j *jabatanService) GetAll(ctx context.Context, page int, size int) ([]dto.JabatanResponse, int64, error) {
	list, err := j.repo.GetAll(ctx, page, size)
	if err != nil {
		return nil, 0, err
	}
	count, err := j.repo.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	return mapper.MapToListResponse(list), count, nil
}

func (j *jabatanService) GetById(ctx context.Context, id int) (*dto.JabatanResponse, error) {
	jabatan, err := j.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mapper.MapToJabatanResponse(jabatan), nil
}

func (j *jabatanService) Create(ctx context.Context, request dto.JabatanRequest) error {
	isUnique, err := j.repo.IsJabatanUnique(ctx, request.Nama)
	if err != nil {
		return err
	}
	if !isUnique {
		return gorm.ErrDuplicatedKey
	}
	jabatan := &model.Jabatan{Nama: request.Nama}
	return j.repo.Save(ctx, jabatan)
}

func (j *jabatanService) Update(ctx context.Context, id int, request dto.JabatanRequest) error {
	existing, err := j.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return gorm.ErrRecordNotFound
	}
	if !strings.EqualFold(existing.Nama, request.Nama) {
		isUnique, err := j.repo.IsJabatanUnique(ctx, request.Nama)
		if err != nil {
			return err
		}
		if !isUnique {
			return gorm.ErrDuplicatedKey
		}
	}
	updates := map[string]any{
		"nama":      request.Nama,
		"UpdatedAt": time.Now(),
	}
	return j.repo.Update(ctx, id, updates)
}

func (j *jabatanService) Delete(ctx context.Context, id int) error {
	jabatan, err := j.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	if jabatan == nil {
		return gorm.ErrRecordNotFound
	}
	return j.repo.Delete(ctx, id)
}
