package model

import (
	"context"
	"time"

	"github.com/riyanamanda/ekinerja/internal/features/jabatan/dto"
)

type Jabatan struct {
	ID        int        `json:"id"`
	Nama      string     `json:"nama"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime:false"`
}

func (Jabatan) TableName() string {
	return "jabatan"
}

type JabatanRepository interface {
	GetAll(ctx context.Context, page, size int) ([]Jabatan, error)
	Save(ctx context.Context, jabatan *Jabatan) error
	GetById(ctx context.Context, id int) (*Jabatan, error)
	Update(ctx context.Context, id int, jabatan map[string]any) error
	Delete(ctx context.Context, id int) error
	Count(ctx context.Context) (int64, error)
	IsJabatanUnique(ctx context.Context, nama string) (bool, error)
}

type JabatanService interface {
	GetAll(ctx context.Context, page, size int) ([]dto.JabatanResponse, int64, error)
	Save(ctx context.Context, request dto.JabatanRequest) error
	GetById(ctx context.Context, id int) (*dto.JabatanResponse, error)
	Update(ctx context.Context, id int, request dto.JabatanRequest) error
	Delete(ctx context.Context, id int) error
}
