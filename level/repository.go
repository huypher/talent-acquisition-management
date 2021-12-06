package level

import (
	"context"
	"database/sql"

	tam "github.com/pghuy/talent-acquistion-management"
	"github.com/pghuy/talent-acquistion-management/domain"
	"github.com/uptrace/bun"
)

type level struct {
	bun.BaseModel `bun:"users"`

	ID   int           `bun:"id"`
	Code tam.LevelType `bun:"code"`
	Name string        `bun:"name"`
}

func (l level) ToModel() domain.Level {
	return domain.Level{
		ID:   l.ID,
		Code: l.Code,
		Name: l.Name,
	}
}

type levelRepository struct {
	db *bun.DB
}

func NewLevelRepository(db *bun.DB) (*levelRepository, error) {
	return &levelRepository{
		db: db,
	}, nil
}

func (r *levelRepository) GetAll(ctx context.Context) ([]domain.Level, error) {
	m := []level{}

	err := r.db.NewSelect().Model(&m).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	talents := make([]domain.Level, len(m))
	for idx, t := range m {
		talents[idx] = t.ToModel()
	}

	return talents, nil
}
