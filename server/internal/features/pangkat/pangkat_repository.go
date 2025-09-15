package pangkat

import (
	"context"

	"github.com/sirupsen/logrus"
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
	err := r.DB.WithContext(ctx).Find(&pangkatList).Error
	if err != nil {
		logrus.Errorf("Error getting all Pangkat: %v", err)
	}
	return pangkatList, err
}

func (r *pangkatRepository) GetById(ctx context.Context, id int64) (Pangkat, error) {
	var pangkat Pangkat
	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&pangkat).Error
	if err != nil {
		logrus.Errorf("Error getting Pangkat by id %d: %v", id, err)
	}
	return pangkat, err
}

func (r *pangkatRepository) Save(ctx context.Context, pangkat Pangkat) error {
	err := r.DB.WithContext(ctx).Create(&pangkat).Error
	if err != nil {
		logrus.Errorf("Error saving Pangkat: %v", err)
	}
	return err
}

func (r *pangkatRepository) Update(ctx context.Context, id int64, pangkat map[string]any) error {
	err := r.DB.WithContext(ctx).Model(&Pangkat{}).Where("id = ?", id).Updates(pangkat).Error
	if err != nil {
		logrus.Errorf("Error updating Pangkat id %d: %v", id, err)
	}
	return err
}

func (r *pangkatRepository) Delete(ctx context.Context, id int64) error {
	err := r.DB.WithContext(ctx).Where("id = ?", id).Delete(&Pangkat{}).Error
	if err != nil {
		logrus.Errorf("Error deleting Pangkat id %d: %v", id, err)
	}
	return err
}
