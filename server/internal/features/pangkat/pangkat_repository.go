package pangkat

import (
	"context"

	"github.com/riyanamanda/ekinerja/internal/features/pangkat/model"
	"gorm.io/gorm"
)

type pangkatRepository struct {
	DB *gorm.DB
}

func NewPangkatRepository(db *gorm.DB) model.PangkatRepository {
	return &pangkatRepository{DB: db}
}

func (r *pangkatRepository) GetAll(ctx context.Context, page, perPage int) ([]model.Pangkat, error) {
	var pangkatList []model.Pangkat
	offset := (page - 1) * perPage
	if err := r.DB.WithContext(ctx).Limit(perPage).Offset(offset).Find(&pangkatList).Error; err != nil {
		return nil, err
	}
	return pangkatList, nil
}

func (r *pangkatRepository) GetById(ctx context.Context, id int64) (model.Pangkat, error) {
	var pangkat model.Pangkat
	if err := r.DB.WithContext(ctx).First(&pangkat, id).Error; err != nil {
		return model.Pangkat{}, err
	}
	return pangkat, nil
}

func (r *pangkatRepository) GetByName(ctx context.Context, name string) (model.Pangkat, error) {
	var pangkat model.Pangkat
	if err := r.DB.WithContext(ctx).Where("LOWER(nama) = LOWER(?)", name).First(&pangkat).Error; err != nil {
		return model.Pangkat{}, err
	}
	return pangkat, nil
}

func (r *pangkatRepository) Save(ctx context.Context, pangkat model.Pangkat) error {
	return r.DB.WithContext(ctx).Create(&pangkat).Error
}

func (r *pangkatRepository) Update(ctx context.Context, id int64, pangkat map[string]any) error {
	return r.DB.WithContext(ctx).Model(&model.Pangkat{}).Where("id = ?", id).Updates(pangkat).Error
}

func (r *pangkatRepository) Delete(ctx context.Context, id int64) error {
	return r.DB.WithContext(ctx).Where("id = ?", id).Delete(&model.Pangkat{}).Error
}

func (r *pangkatRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := r.DB.WithContext(ctx).Model(&model.Pangkat{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *pangkatRepository) IsPangkatUnique(ctx context.Context, nama string) (bool, error) {
	var count int64
	if err := r.DB.WithContext(ctx).Model(&model.Pangkat{}).Where("LOWER(nama) = LOWER(?)", nama).Count(&count).Error; err != nil {
		return false, err
	}
	return count == 0, nil
}
