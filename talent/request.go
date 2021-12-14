package talent

import (
	"time"

	tam "github.com/huypher/talent-acquisition-management"
)

type getTalentListRequest struct {
	FullName string `json:"full_name" keyname:"full_name,omitempty"`
	PageID   int    `json:"page_id" keyname:"page_id,omitempty"`
	PerPage  int    `json:"per_page" keyname:"per_page,omitempty"`
}

type addTalentRequest struct {
	FullName           string        `json:"full_name" binding:"required"`
	Gender             string        `json:"gender"`
	Birthdate          string        `json:"birthdate"`
	Phone              string        `json:"phone"`
	Email              string        `json:"email"`
	AppliedPosition    string        `json:"applied_position" binding:"required"`
	Level              tam.LevelType `json:"level"`
	Department         string        `json:"department"`
	Project            string        `json:"project"`
	CV                 string        `json:"cv"`
	Criteria           string        `json:"criteria"`
	ScheduledInterview time.Time     `json:"scheduled_interview"`
	InterviewResult    string        `json:"interview_result"`
}

type updateTalentRequest struct {
	FullName           string        `json:"full_name" keyname:"full_name"`
	Gender             string        `json:"gender" keyname:"gender"`
	Birthdate          string        `json:"birthdate" keyname:"birthdate"`
	Phone              string        `json:"phone" keyname:"phone"`
	Email              string        `json:"email" keyname:"email"`
	AppliedPosition    string        `json:"applied_position" keyname:"applied_position" binding:"required"`
	Level              tam.LevelType `json:"level" keyname:"level"`
	Department         string        `json:"department" keyname:"department"`
	Project            string        `json:"project" keyname:"project"`
	CV                 string        `json:"cv" keyname:"cv"`
	Criteria           string        `json:"criteria" keyname:"criteria"`
	ScheduledInterview time.Time     `json:"scheduled_interview" keyname:"scheduled_interview"`
	InterviewResult    string        `json:"interview_result" keyname:"interview_result"`
}
