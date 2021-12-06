package talent

import (
	"context"
	"database/sql"
	"time"

	tam "github.com/pghuy/talent-acquistion-management"

	"github.com/pghuy/talent-acquistion-management/domain"
	"github.com/uptrace/bun"
)

type talent struct {
	bun.BaseModel `bun:"users"`

	ID                 int           `bun:"id"`
	FullName           string        `bun:"full_name"`
	Gender             string        `bun:"gender"`
	YearOfBirth        string        `bun:"year_of_birth"`
	Phone              string        `bun:"phone"`
	Email              string        `bun:"email"`
	AppliedPosition    string        `bun:"applied_position"`
	Level              tam.LevelType `bun:"level"`
	Department         string        `bun:"department"`
	Project            string        `bun:"project"`
	CV                 string        `bun:"cv"`
	Criteria           string        `bun:"criteria"`
	ScheduledInterview time.Time     `bun:"scheduled_interview"`
	InterviewResult    string        `bun:"interview_result"`
	CreatedAt          time.Time     `bun:"created_at"`
	UpdatedAt          time.Time     `bun:"updated_at"`
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
	db *bun.DB
}

func NewTalentRepository(db *bun.DB) (*talentRepository, error) {
	return &talentRepository{
		db: db,
	}, nil
}

func (r *talentRepository) GetByUserName(ctx context.Context, userName string) (*domain.Talent, error) {
	m := talent{}

	err := r.db.NewSelect().Model(&m).Where("username = ?", userName).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	talent := m.ToModel()

	return &talent, nil
}

func (r *talentRepository) GetByID(ctx context.Context, id int) (*domain.Talent, error) {
	m := talent{}

	err := r.db.NewSelect().Model(&m).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	talent := m.ToModel()

	return &talent, nil
}
