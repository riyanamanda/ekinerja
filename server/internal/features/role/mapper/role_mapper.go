package mapper

import (
	"github.com/riyanamanda/ekinerja/internal/features/role/dto"
	"github.com/riyanamanda/ekinerja/internal/features/role/model"
)

func MapToListResponses(list []model.Role) []dto.RoleResponse {
	responses := make([]dto.RoleResponse, len(list))
	for i, role := range list {
		responses[i] = *MapToRoleResponse(&role)
	}
	return responses
}

func MapToRoleResponse(role *model.Role) *dto.RoleResponse {
	return &dto.RoleResponse{
		ID:   role.ID,
		Nama: role.Nama,
	}
}
