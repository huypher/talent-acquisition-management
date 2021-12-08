package talent

import (
	"context"
	"errors"
	"time"

	"github.com/huypher/kit/container"

	"gorm.io/gorm"

	tam "github.com/huypher/talent-acquisition-management"

	"github.com/huypher/talent-acquisition-management/domain"
)

type talent struct {
	gorm.Model `bun:"talents"`

	ID                 int
	FullName           string
	Gender             string
	YearOfBirth        string
	Phone              string
	Email              string
	AppliedPosition    string
	Level              tam.LevelType
	Department         string
	Project            string
	CV                 string
	Criteria           string
	ScheduledInterview time.Time
	InterviewResult    string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (t talent) ToModel() domain.Talent {
	return domain.Talent{
		ID:                 t.ID,
		FullName:           t.FullName,
		Gender:             t.Gender,
		YearOfBirth:        t.YearOfBirth,
		Phone:              t.Phone,
		Email:              t.Email,
		AppliedPosition:    t.AppliedPosition,
		Level:              t.Level,
		Department:         t.Department,
		Project:            t.Project,
		CV:                 t.CV,
		Criteria:           t.Criteria,
		ScheduledInterview: t.ScheduledInterview,
		InterviewResult:    t.InterviewResult,
		CreatedAt:          t.CreatedAt,
		UpdatedAt:          t.UpdatedAt,
	}
}

type talentRepository struct {
	db *gorm.DB
}

func NewTalentRepository(db *gorm.DB) (*talentRepository, error) {
	return &talentRepository{
		db: db,
	}, nil
}

func (r *talentRepository) GetByID(ctx context.Context, id int) (*domain.Talent, error) {
	m := talent{}

	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	talent := m.ToModel()

	return &talent, nil
}

func (r *talentRepository) GetList(ctx context.Context, filter container.Map, pageID, perPage int) ([]domain.Talent, error) {
	var m []talent

	q := r.db.Where(map[string]interface{}(filter))

	if pageID > 0 && perPage > 0 {
		offset := perPage * (pageID - 1)
		q = q.Offset(offset).Limit(perPage)
	}

	if err := q.WithContext(ctx).Find(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	talents := make([]domain.Talent, len(m))
	for idx, t := range m {
		talents[idx] = t.ToModel()
	}

	return talents, nil
}
