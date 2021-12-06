package domain

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	tam "github.com/pghuy/talent-acquistion-management"
)

type Talent struct {
	ID                 int           `json:"id"`
	FullName           string        `json:"full_name"`
	Gender             string        `json:"gender"`
	YearOfBirth        string        `json:"year_of_birth"`
	Phone              string        `json:"phone"`
	Email              string        `json:"email"`
	AppliedPosition    string        `json:"applied_position"`
	Level              tam.LevelType `json:"level"`
	Department         string        `json:"department"`
	Project            string        `json:"project"`
	CV                 string        `json:"cv"`
	Criteria           string        `json:"criteria"`
	ScheduledInterview time.Time     `json:"scheduled_interview"`
	InterviewResult    string        `json:"interview_result"`
	CreatedAt          time.Time     `bun:"created_at"`
	UpdatedAt          time.Time     `json:"updated_at"`
}

type TalentUsecase interface {
	GetByUserName(ctx context.Context, username string) (*Talent, error)
	GetByID(ctx context.Context, id int) (*Talent, error)
}

type TalentRepository interface {
	GetByUserName(ctx context.Context, username string) (*Talent, error)
	GetByID(ctx context.Context, id int) (*Talent, error)
}

type TalentDelivery interface {
	Handler(c *gin.RouterGroup)
}
