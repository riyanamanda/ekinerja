package model

import (
	"context"
	"time"

	"github.com/riyanamanda/ekinerja/internal/features/bidang/dto"
)

type Bidang struct {
	ID        int        `json:"id"`
	Nama      string     `json:"nama"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime:false"`
}

func (Bidang) TableName() string {
	return "bidang"
}

type BidangRepository interface {
	GetAll(ctx context.Context, page int, size int) ([]Bidang, error)
	Save(ctx context.Context, bidang *Bidang) error
	GetById(ctx context.Context, id int64) (*Bidang, error)
	Update(ctx context.Context, id int64, bidang map[string]any) error
	Delete(ctx context.Context, id int64) error
	Count(ctx context.Context) (int64, error)
	IsBidangUnique(ctx context.Context, nama string) (bool, error)
}

type BidangService interface {
	GetAll(ctx context.Context, page int, size int) ([]dto.BidangResponse, int64, error)
	GetById(ctx context.Context, id int64) (*dto.BidangResponse, error)
	Save(ctx context.Context, request dto.BidangRequest) error
	Update(ctx context.Context, id int64, request dto.BidangRequest) error
	Delete(ctx context.Context, id int64) error
}
