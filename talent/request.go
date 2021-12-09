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
	FullName           string        `json:"full_name"`
	Gender             string        `json:"gender"`
	YearOfBirth        int           `json:"year_of_birth"`
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
}
