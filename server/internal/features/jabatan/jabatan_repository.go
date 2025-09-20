package jabatan

import (
	"context"

	"github.com/riyanamanda/ekinerja/internal/features/jabatan/model"
	"gorm.io/gorm"
)

type jabatanRepository struct {
	DB *gorm.DB
}

func NewJabatanRepository(db *gorm.DB) model.JabatanRepository {
	return &jabatanRepository{DB: db}
}

func (j *jabatanRepository) GetAll(ctx context.Context, page int, size int) ([]model.Jabatan, error) {
	var list []model.Jabatan
	offset := (page - 1) * size
	if err := j.DB.WithContext(ctx).Limit(size).Offset(offset).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (j *jabatanRepository) GetById(ctx context.Context, id int) (*model.Jabatan, error) {
	var jabatan model.Jabatan
	if err := j.DB.WithContext(ctx).First(&jabatan, id).Error; err != nil {
		return nil, err
	}
	return &jabatan, nil
}

func (j *jabatanRepository) Save(ctx context.Context, jabatan *model.Jabatan) error {
	return j.DB.WithContext(ctx).Create(jabatan).Error
}

func (j *jabatanRepository) Update(ctx context.Context, id int, jabatan map[string]any) error {
	return j.DB.WithContext(ctx).Model(&model.Jabatan{}).Where("id = ?", id).Updates(jabatan).Error
}

func (j *jabatanRepository) Delete(ctx context.Context, id int) error {
	return j.DB.WithContext(ctx).Delete(&model.Jabatan{}, id).Error
}

func (j *jabatanRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := j.DB.WithContext(ctx).Model(&model.Jabatan{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (j *jabatanRepository) IsJabatanUnique(ctx context.Context, nama string) (bool, error) {
	var count int64
	if err := j.DB.WithContext(ctx).Model(&model.Jabatan{}).Where("LOWER(nama) = LOWER(?)", nama).Count(&count).Error; err != nil {
		return false, err
	}
	return count == 0, nil
}
