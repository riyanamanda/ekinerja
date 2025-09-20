package ruangan

import (
	"context"

	"github.com/riyanamanda/ekinerja/internal/features/ruangan/model"
	"gorm.io/gorm"
)

type ruanganRepository struct {
	DB *gorm.DB
}

func NewRuanganRepository(db *gorm.DB) model.RuanganRepository {
	return &ruanganRepository{DB: db}
}

func (r *ruanganRepository) GetAll(ctx context.Context, page int, size int) ([]model.Ruangan, error) {
	var list []model.Ruangan
	offset := (page - 1) * size
	if err := r.DB.WithContext(ctx).Limit(size).Offset(offset).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *ruanganRepository) Save(ctx context.Context, ruangan *model.Ruangan) error {
	return r.DB.WithContext(ctx).Create(ruangan).Error
}

func (r *ruanganRepository) GetById(ctx context.Context, id int) (*model.Ruangan, error) {
	var ruangan model.Ruangan
	if err := r.DB.WithContext(ctx).First(&ruangan, id).Error; err != nil {
		return nil, err
	}
	return &ruangan, nil
}

func (r *ruanganRepository) Update(ctx context.Context, id int, ruangan map[string]any) error {
	return r.DB.WithContext(ctx).Model(&model.Ruangan{}).Where("id = ?", id).Updates(ruangan).Error
}

func (r *ruanganRepository) Delete(ctx context.Context, id int) error {
	return r.DB.WithContext(ctx).Delete(&model.Ruangan{}, id).Error
}

func (r *ruanganRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := r.DB.WithContext(ctx).Model(&model.Ruangan{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *ruanganRepository) IsRuanganUnique(ctx context.Context, nama string) (bool, error) {
	var count int64
	if err := r.DB.WithContext(ctx).Model(&model.Ruangan{}).Where("LOWER(nama) = LOWER(?)", nama).Count(&count).Error; err != nil {
		return false, err
	}
	return count == 0, nil
}
