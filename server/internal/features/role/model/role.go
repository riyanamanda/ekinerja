package model

import (
	"context"

	"github.com/riyanamanda/ekinerja/internal/features/role/dto"
)

type Role struct {
	ID   int64  `db:"id"`
	Nama string `db:"nama"`
}

func (Role) TableName() string {
	return "role"
}

type RoleRepository interface {
	GetAll(ctx context.Context) ([]Role, error)
	Save(ctx context.Context, role *Role) error
	GetByID(ctx context.Context, id int64) (*Role, error)
	Update(ctx context.Context, id int, role map[string]any) error
	Delete(ctx context.Context, role Role) error
	IsRoleUnique(ctx context.Context, name string) (bool, error)
}

type RoleService interface {
	GetAll(ctx context.Context) ([]dto.RoleResponse, error)
	Create(ctx context.Context, request dto.RoleRequest) error
	GetByID(ctx context.Context, id int) (*dto.RoleResponse, error)
	Update(ctx context.Context, id int, request dto.RoleRequest) error
	Delete(ctx context.Context, id int) error
}
