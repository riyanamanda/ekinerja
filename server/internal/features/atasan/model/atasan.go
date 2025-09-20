package model

import (
	"context"

	"github.com/riyanamanda/ekinerja/internal/features/atasan/dto"
)

type Atasan struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	IsActive bool   `json:"is_active" gorm:"default:true"`
}

func (Atasan) TableName() string {
	return "atasan"
}

type AtasanRepository interface {
	GetAll(ctx context.Context, page int, size int) ([]Atasan, error)
	Save(ctx context.Context, atasan *Atasan) error
	GetByID(ctx context.Context, id int) (*Atasan, error)
	Update(ctx context.Context, id int, atasan map[string]any) error
	Delete(ctx context.Context, id int) error
	Count(ctx context.Context) (int64, error)
	IsAtasanUnique(ctx context.Context, nama string) (bool, error)
}

type AtasanService interface {
	GetAll(ctx context.Context, page int, size int) ([]dto.AtasanResponse, int64, error)
	Create(ctx context.Context, request *dto.AtasanCreateRequest) error
	GetByID(ctx context.Context, id int) (*dto.AtasanResponse, error)
	Update(ctx context.Context, id int, request dto.AtasanUpdateRequest) error
	Delete(ctx context.Context, id int) error
}
