package domain

import (
	"context"
	"time"

	"github.com/huypher/kit/container"

	"github.com/gin-gonic/gin"
	tam "github.com/huypher/talent-acquisition-management"
)

type Talent struct {
	ID                 int           `json:"id"`
	FullName           string        `json:"full_name"`
	Gender             string        `json:"gender"`
	Birthdate          string        `json:"birthdate"`
	Phone              string        `json:"phone"`
	Email              string        `json:"email"`
	AppliedPosition    string        `json:"applied_position"`
	Level              tam.LevelType `json:"level"`
	Department         string        `json:"department"`
	Project            string        `json:"project"`
	CV                 string        `json:"cv"`
	Criteria           string        `json:"criteria"`
	ScheduledInterview time.Time     `json:"scheduled_interview"`
	Interviewer        string        `json:"interviewer"`
	InterviewResult    string        `json:"interview_result"`
	Note               string        `json:"note"`
	CreatedAt          time.Time     `json:"created_at"`
	UpdatedAt          time.Time     `json:"updated_at"`
}

type TalentUsecase interface {
	GetByID(ctx context.Context, id int) (*Talent, error)
	GetList(ctx context.Context, filter container.Map, pageID, perPage int) ([]Talent, error)
	AddTalent(ctx context.Context, talent Talent) error
	UpdateTalent(ctx context.Context, talentID int, params container.Map) error
}

type TalentRepository interface {
	GetByID(ctx context.Context, id int) (*Talent, error)
	GetList(ctx context.Context, filter container.Map, pageID, perPage int) ([]Talent, error)
	Create(ctx context.Context, talent Talent) error
	Update(ctx context.Context, talentID int, params container.Map) error
}

type TalentDelivery interface {
	Handler(c *gin.RouterGroup)
}
