package atasan

import (
	"context"

	"github.com/riyanamanda/ekinerja/internal/features/atasan/model"
	"gorm.io/gorm"
)

type atasanRepository struct {
	DB *gorm.DB
}

func NewAtasanRepository(db *gorm.DB) model.AtasanRepository {
	return &atasanRepository{
		DB: db,
	}
}

func (a *atasanRepository) GetAll(ctx context.Context, page int, size int) ([]model.Atasan, error) {
	var list []model.Atasan
	offset := (page - 1) * size
	if err := a.DB.WithContext(ctx).Limit(size).Offset(offset).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (a *atasanRepository) Save(ctx context.Context, atasan *model.Atasan) error {
	return a.DB.WithContext(ctx).Create(atasan).Error
}

func (a *atasanRepository) GetByID(ctx context.Context, id int) (*model.Atasan, error) {
	var atasan model.Atasan
	if err := a.DB.WithContext(ctx).First(&atasan, id).Error; err != nil {
		return nil, err
	}
	return &atasan, nil
}

func (a *atasanRepository) Update(ctx context.Context, id int, atasan map[string]any) error {
	return a.DB.WithContext(ctx).Model(&model.Atasan{}).Where("id = ?", id).Updates(atasan).Error
}

func (a *atasanRepository) Delete(ctx context.Context, id int) error {
	return a.DB.WithContext(ctx).Where("id = ?", id).Delete(&model.Atasan{}).Error
}

func (a *atasanRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := a.DB.WithContext(ctx).Model(&model.Atasan{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (a *atasanRepository) IsAtasanUnique(ctx context.Context, nama string) (bool, error) {
	var count int64
	if err := a.DB.WithContext(ctx).Model(&model.Atasan{}).Where("LOWER(nama) = LOWER(?)", nama).Count(&count).Error; err != nil {
		return false, err
	}
	return count == 0, nil
}
