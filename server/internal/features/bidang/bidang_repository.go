package bidang

import (
	"context"

	"github.com/riyanamanda/ekinerja/internal/features/bidang/model"
	"gorm.io/gorm"
)

type bidangRepository struct {
	DB *gorm.DB
}

func NewBidangRepository(db *gorm.DB) model.BidangRepository {
	return &bidangRepository{DB: db}
}

func (b *bidangRepository) GetAll(ctx context.Context, page int, size int) ([]model.Bidang, error) {
	var list []model.Bidang
	offset := (page - 1) * size
	if err := b.DB.WithContext(ctx).Limit(size).Offset(offset).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (b *bidangRepository) Save(ctx context.Context, bidang *model.Bidang) error {
	return b.DB.WithContext(ctx).Create(bidang).Error
}

func (b *bidangRepository) GetById(ctx context.Context, id int64) (*model.Bidang, error) {
	var bidang model.Bidang
	if err := b.DB.WithContext(ctx).First(&bidang, id).Error; err != nil {
		return nil, err
	}
	return &bidang, nil
}

func (b *bidangRepository) GetByName(ctx context.Context, name string) (*model.Bidang, error) {
	var bidang model.Bidang
	if err := b.DB.WithContext(ctx).Where("LOWER(nama) = LOWER(?)", name).First(&bidang).Error; err != nil {
		return nil, err
	}
	return &bidang, nil
}

func (b *bidangRepository) Update(ctx context.Context, id int64, bidang map[string]any) error {
	return b.DB.WithContext(ctx).Model(&model.Bidang{}).Where("id = ?", id).Updates(bidang).Error
}

func (b *bidangRepository) Delete(ctx context.Context, id int64) error {
	return b.DB.WithContext(ctx).Delete(&model.Bidang{}, id).Error
}

func (b *bidangRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := b.DB.WithContext(ctx).Model(&model.Bidang{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (b *bidangRepository) IsBidangUnique(ctx context.Context, nama string) (bool, error) {
	var count int64
	if err := b.DB.WithContext(ctx).Model(&model.Bidang{}).Where("LOWER(nama) = LOWER(?)", nama).Count(&count).Error; err != nil {
		return false, err
	}
	return count == 0, nil
}
