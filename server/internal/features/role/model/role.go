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
	GetByID(ctx context.Context, id int64) (*Role, error)
}

type RoleService interface {
	GetAll(ctx context.Context) ([]dto.RoleResponse, error)
	GetByID(ctx context.Context, id int) (*dto.RoleResponse, error)
}
