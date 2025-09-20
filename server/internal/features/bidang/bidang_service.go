package bidang

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/riyanamanda/ekinerja/internal/features/bidang/dto"
	"github.com/riyanamanda/ekinerja/internal/features/bidang/mapper"
	"github.com/riyanamanda/ekinerja/internal/features/bidang/model"
	"gorm.io/gorm"
)

type bidangService struct {
	repo model.BidangRepository
}

func NewBidangService(repo model.BidangRepository) model.BidangService {
	return &bidangService{repo: repo}
}

func (b *bidangService) GetAll(ctx context.Context, page int, size int) ([]dto.BidangResponse, int64, error) {
	list, err := b.repo.GetAll(ctx, page, size)
	if err != nil {
		return nil, 0, err
	}
	count, err := b.repo.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	return mapper.MapToListResponse(list), count, nil
}

func (b *bidangService) Create(ctx context.Context, request dto.BidangRequest) error {
	isUnique, err := b.repo.IsBidangUnique(ctx, request.Nama)
	if err != nil {
		return err
	}
	if !isUnique {
		return gorm.ErrDuplicatedKey
	}
	bidang := &model.Bidang{Nama: request.Nama}
	return b.repo.Save(ctx, bidang)
}

func (b *bidangService) GetById(ctx context.Context, id int64) (*dto.BidangResponse, error) {
	bidang, err := b.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mapper.MapToBidangResponse(bidang), nil
}

func (b *bidangService) Update(ctx context.Context, id int64, request dto.BidangRequest) error {
	existing, err := b.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return gorm.ErrRecordNotFound
	}
	if !strings.EqualFold(existing.Nama, request.Nama) {
		isUnique, err := b.repo.IsBidangUnique(ctx, request.Nama)
		if err != nil {
			return err
		}
		if !isUnique {
			return gorm.ErrDuplicatedKey
		}
	}
	updates := map[string]any{
		"nama":       request.Nama,
		"updated_at": time.Now(),
	}
	return b.repo.Update(ctx, id, updates)
}

func (b *bidangService) Delete(ctx context.Context, id int64) error {
	bidang, err := b.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	if bidang == nil {
		return gorm.ErrRecordNotFound
	}
	return b.repo.Delete(ctx, id)
}
