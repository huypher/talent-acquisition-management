package level

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	tam "github.com/huypher/talent-acquisition-management"
	"github.com/huypher/talent-acquisition-management/domain"
)

type level struct {
	gorm.Model `bun:"levels"`

	ID   int
	Code tam.LevelType
	Name string
}

func (l level) ToModel() domain.Level {
	return domain.Level{
		ID:   l.ID,
		Code: l.Code,
		Name: l.Name,
	}
}

type levelRepository struct {
	db *gorm.DB
}

func NewLevelRepository(db *gorm.DB) (*levelRepository, error) {
	return &levelRepository{
		db: db,
	}, nil
}

func (r *levelRepository) GetAll(ctx context.Context) ([]domain.Level, error) {
	m := []level{}

	if err := r.db.WithContext(ctx).Find(&m).Error; err != nil {
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
