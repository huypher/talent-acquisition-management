package talent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/huypher/kit/container"
	tam "github.com/huypher/talent-acquisition-management"
	"github.com/huypher/talent-acquisition-management/domain"
	"gorm.io/gorm"
)

type talent struct {
	ID                 int
	FullName           string
	Gender             string
	YearOfBirth        int
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

func ToEntity(model domain.Talent) talent {
	return talent{
		ID:                 model.ID,
		FullName:           model.FullName,
		Gender:             model.Gender,
		YearOfBirth:        model.YearOfBirth,
		Phone:              model.Phone,
		Email:              model.Email,
		AppliedPosition:    model.AppliedPosition,
		Level:              model.Level,
		Department:         model.Department,
		Project:            model.Project,
		CV:                 model.CV,
		Criteria:           model.Criteria,
		ScheduledInterview: model.ScheduledInterview,
		InterviewResult:    model.InterviewResult,
		CreatedAt:          model.CreatedAt,
		UpdatedAt:          model.UpdatedAt,
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
	entity := talent{}

	if err := r.db.WithContext(ctx).First(&entity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	talent := entity.ToModel()

	return &talent, nil
}

func (r *talentRepository) GetList(ctx context.Context, filter container.Map, pageID, perPage int) ([]domain.Talent, error) {
	var entity []talent

	q := r.db.WithContext(ctx)
	if !filter.IsEmpty() {
		q = r.db.Where(map[string]interface{}(filter))
	}

	if pageID > 0 && perPage > 0 {
		offset := perPage * (pageID - 1)
		q = q.Offset(offset).Limit(perPage)
	}

	q.Order("id DESC")

	if err := q.Find(&entity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	fmt.Printf("entity=%v\n", entity)
	talents := make([]domain.Talent, len(entity))
	for idx, t := range entity {
		talents[idx] = t.ToModel()
	}

	return talents, nil
}

func (r *talentRepository) Create(ctx context.Context, model domain.Talent) error {
	entity := ToEntity(model)

	if err := r.db.WithContext(ctx).Create(&entity).Error; err != nil {
		return err
	}

	return nil
}
