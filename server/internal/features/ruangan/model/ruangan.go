package model

import (
	"context"
	"time"

	"github.com/riyanamanda/ekinerja/internal/features/ruangan/dto"
)

type Ruangan struct {
	ID        int        `json:"id"`
	Nama      string     `json:"nama"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime:false"`
}

func (Ruangan) TableName() string {
	return "ruangan"
}

type RuanganRepository interface {
	GetAll(ctx context.Context, page int, size int) ([]Ruangan, error)
	Save(ctx context.Context, ruangan *Ruangan) error
	GetById(ctx context.Context, id int) (*Ruangan, error)
	Update(ctx context.Context, id int, ruangan map[string]any) error
	Delete(ctx context.Context, id int) error
	Count(ctx context.Context) (int64, error)
	IsRuanganUnique(ctx context.Context, nama string) (bool, error)
}

type RuanganService interface {
	GetAll(ctx context.Context, page int, size int) ([]dto.RuanganResponse, int64, error)
	GetById(ctx context.Context, id int) (*dto.RuanganResponse, error)
	Create(ctx context.Context, request dto.RuanganRequest) error
	Update(ctx context.Context, id int, request dto.RuanganRequest) error
	Delete(ctx context.Context, id int) error
}
