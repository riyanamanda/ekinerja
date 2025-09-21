package role

import (
	"context"
	"errors"
	"strings"

	"github.com/riyanamanda/ekinerja/internal/features/role/dto"
	"github.com/riyanamanda/ekinerja/internal/features/role/mapper"
	"github.com/riyanamanda/ekinerja/internal/features/role/model"
	"gorm.io/gorm"
)

type roleService struct {
	repo model.RoleRepository
}

func NewRoleService(repo model.RoleRepository) model.RoleService {
	return &roleService{
		repo: repo,
	}
}

func (r *roleService) GetAll(ctx context.Context) ([]dto.RoleResponse, error) {
	roles, err := r.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.MapToListResponses(roles), nil
}

func (r *roleService) GetByID(ctx context.Context, id int) (*dto.RoleResponse, error) {
	role, err := r.repo.GetByID(ctx, int64(id))
	if err != nil {
		return nil, err
	}
	return mapper.MapToRoleResponse(role), nil
}

func (r *roleService) Create(ctx context.Context, request dto.RoleRequest) error {
	isUnique, err := r.repo.IsRoleUnique(ctx, request.Nama)
	if err != nil {
		return err
	}
	if !isUnique {
		return gorm.ErrDuplicatedKey
	}
	role := &model.Role{Nama: request.Nama}
	return r.repo.Save(ctx, role)
}

func (r *roleService) Update(ctx context.Context, id int, request dto.RoleRequest) error {
	existing, err := r.repo.GetByID(ctx, int64(id))
	if err != nil {
		return err
	}
	if existing == nil {
		return gorm.ErrRecordNotFound
	}
	if !strings.EqualFold(existing.Nama, request.Nama) {
		isUnique, err := r.repo.IsRoleUnique(ctx, request.Nama)
		if err != nil {
			return err
		}
		if !isUnique {
			return gorm.ErrDuplicatedKey
		}
	}
	updates := map[string]any{
		"nama": request.Nama,
	}
	return r.repo.Update(ctx, id, updates)
}

func (r *roleService) Delete(ctx context.Context, id int) error {
	existing, err := r.repo.GetByID(ctx, int64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.repo.Delete(ctx, *existing)
}
