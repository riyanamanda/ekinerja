package role

import (
	"context"

	"github.com/riyanamanda/ekinerja/internal/features/role/model"
	"gorm.io/gorm"
)

type roleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) model.RoleRepository {
	return &roleRepository{
		DB: db,
	}
}

func (r *roleRepository) GetAll(ctx context.Context) ([]model.Role, error) {
	var roles []model.Role
	if err := r.DB.WithContext(ctx).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *roleRepository) GetByID(ctx context.Context, id int64) (*model.Role, error) {
	var role model.Role
	if err := r.DB.WithContext(ctx).First(&role, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
