package pangkat

import (
	"context"

	"gorm.io/gorm"
)

type pangkatRepository struct {
	DB *gorm.DB
}

func NewPangkatRepository(db *gorm.DB) PangkatRepository {
	return &pangkatRepository{DB: db}
}

func (r *pangkatRepository) GetAll(ctx context.Context) ([]Pangkat, error) {
	var pangkatList []Pangkat
	if err := r.DB.WithContext(ctx).Find(&pangkatList).Error; err != nil {
		return nil, err
	}
	return pangkatList, nil
}

func (r *pangkatRepository) GetById(ctx context.Context, id int64) (Pangkat, error) {
	var pangkat Pangkat
	if err := r.DB.WithContext(ctx).First(&pangkat, id).Error; err != nil {
		return Pangkat{}, err
	}
	return pangkat, nil
}

func (r *pangkatRepository) Save(ctx context.Context, pangkat Pangkat) error {
	return r.DB.WithContext(ctx).Create(&pangkat).Error
}

func (r *pangkatRepository) Update(ctx context.Context, id int64, pangkat map[string]any) error {
	return r.DB.WithContext(ctx).Model(&Pangkat{}).Where("id = ?", id).Updates(pangkat).Error
}

func (r *pangkatRepository) Delete(ctx context.Context, id int64) error {
	return r.DB.WithContext(ctx).Where("id = ?", id).Delete(&Pangkat{}).Error
}
