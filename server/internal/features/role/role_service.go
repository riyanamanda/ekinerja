package role

import (
	"context"

	"github.com/riyanamanda/ekinerja/internal/features/role/dto"
	"github.com/riyanamanda/ekinerja/internal/features/role/mapper"
	"github.com/riyanamanda/ekinerja/internal/features/role/model"
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
