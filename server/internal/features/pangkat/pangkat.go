package pangkat

import (
	"context"
	"time"

	"github.com/riyanamanda/ekinerja/internal/features/pangkat/dto"
)

type Pangkat struct {
	ID        int64      `json:"id"`
	Nama      string     `json:"nama"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime:false"`
}

func (Pangkat) TableName() string {
	return "pangkat"
}

type PangkatRepository interface {
	GetAll(ctx context.Context) ([]Pangkat, error)
	Save(ctx context.Context, pangkat Pangkat) error
	GetById(ctx context.Context, id int64) (*Pangkat, error)
	Update(ctx context.Context, id int64, pangkat map[string]any) error
	Delete(ctx context.Context, id int64) error
}

type PangkatService interface {
	GetAll(ctx context.Context) ([]dto.PangkatResponse, error)
	Save(ctx context.Context, request dto.PangkatRequest) error
	GetById(ctx context.Context, id int64) (*dto.PangkatResponse, error)
	Update(ctx context.Context, id int64, request dto.PangkatRequest) error
	Delete(ctx context.Context, id int64) error
}
