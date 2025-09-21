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

func (r *roleRepository) Save(ctx context.Context, role *model.Role) error {
	return r.DB.WithContext(ctx).Create(role).Error
}

func (r *roleRepository) GetByID(ctx context.Context, id int64) (*model.Role, error) {
	var role model.Role
	if err := r.DB.WithContext(ctx).First(&role, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) Update(ctx context.Context, id int, role map[string]any) error {
	return r.DB.WithContext(ctx).Model(&model.Role{}).Where("id = ?", id).Updates(role).Error
}

func (r *roleRepository) Delete(ctx context.Context, role model.Role) error {
	return r.DB.WithContext(ctx).Delete(&role).Error
}

func (r *roleRepository) IsRoleUnique(ctx context.Context, name string) (bool, error) {
	var count int64
	if err := r.DB.WithContext(ctx).Model(&model.Role{}).Where("LOWER(nama) = LOWER(?)", name).Count(&count).Error; err != nil {
		return false, err
	}
	return count == 0, nil
}
